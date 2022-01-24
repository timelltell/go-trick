package datagrip

import (
	_struct "GolangTrick/Engine/struct"
	"github.com/codegangsta/inject"
	"golang.org/x/net/context"
)

// QueryParam 调用Fetch需要的参数，调用方只填关心的项
type QueryParam struct {
	EventData  *_struct.EventData
	ObjectList []_struct.Object
	UosList    []_struct.UserOPStream
}

// GetStepProvider 根据uosList获取需要的provider
// 根据condition的key对step进行分组，每个key对应一个provider
// 每个provider处理一组step
func GetStepProvider(queryParam QueryParam) []DataProvider {
	stepGroup := make(map[string][]_struct.Step)
	for _, uos := range queryParam.UosList {
		for _, step := range uos.Steps {
			for _, condition := range step.Conditions {
				key := condition.Key
				if _, ok := stepGroup[key]; ok {
					stepGroup[key] = append(stepGroup[key], step)
					continue
				}
				stepGroup[key] = []_struct.Step{step}
			}
		}
	}
	providers := make([]DataProvider, 0, len(stepGroup))
	for key := range stepGroup {
		provider := getProvider(key, stepGroup[key], queryParam.EventData)
		if provider != nil {
			providers = append(providers, provider)
		}
	}
	return providers
}

// DataProvider 每个conditionKey都需要实现此接口
type DataProvider interface {
	Match(condtitionKey string) bool
	Prepare(context.Context) (keyList []string, param map[string]string, err error)
	Bind(context.Context, map[string]string, inject.TypeMapper)
}

// 由于最终curl DMP的参数放在一个大的map[string]string对象中
// 如果两个provider的param使用到了同一个key，那么该provider
// 需要实现这个接口来处理这种冲突，否则该key的值会被覆盖
type collisionResolver interface {
	merge(map[string]string, map[string]string) map[string]string
}

func getProvider(conditionKey string, stepList []_struct.Step, eventData *_struct.EventData) DataProvider {
	for key, provider := range stepRelatedConstructor {
		if key == conditionKey {
			return provider(ConstructParam{
				StepList:  stepList,
				EventData: eventData,
			})
		}
	}
	return nil
}

// FetchData 根据provider获取数据
// 业务可以自由组合
// 非canvas维度的数据可以通过GetStepProvider获取对应DataProvider
// canvas级别的过滤，根据需要自行New对应的provider
func FetchData(ctx context.Context, eventData *_struct.EventData, providerList ...DataProvider) (inject.Invoker, error) {
	return fetchData(ctx, eventData, providerList...)
}

func fetchData(ctx context.Context, eventData *_struct.EventData, providerList ...DataProvider) (inject.Invoker, error) {
	if 0 == len(providerList) {
		return inject.New(), nil
	}
	dmpParam := make(map[string]string)

	// 存储真正需要curl dmp的key以及这个key对应的请求参数
	paramMap := make(map[string]map[string]string)
	// 从缓存中取的对应于某个key的返回结果
	var cachedResult map[string]string
	/*
		批量请求需要做cache，非批量就不cache
		因为批量预估一次会把快车、优享、拼车等不同车型的预估发过来
		但是对于一个活动N天命中了多少次，起点终点的围栏，用户标签以及apollo分的人群
		这些结果其实是一样的
		因此如果是批量预估的话，同一个traceID带了N种不同车型的请求
		根据"traceID+key+对应key的请求参数"
		把部分结果缓存起来，减少对DMP的调用以及网络开销。
		但是cache只对一个请求有效，所以cache的expireTime只设置了两秒（可讨论）
	*/
	shouldUseCache := eventData.BatchRequestFlag
	if shouldUseCache {
		cachedResult = make(map[string]string)
	}
	// datagrip
	container := inject.New()
	commomConditionFlag := -1
	for index, provider := range providerList {
		keys, param, err := provider.Prepare(ctx)
		if nil != err {
			return nil, err
		}
		if 0 == len(keys) {
			continue
		}
		// 枚举每个key，如果当前traceID下以同样的参数请求过这个key，就直接从缓存里取
		for _, k := range keys {
			if k == "pope.common" {
				//当key为openConditionPrefix, 直接将结果放进ioc
				provider.Bind(ctx, param, container)
				commomConditionFlag = index
				continue
			}
			// 当前key是否命中缓存
			matchedCache := false
			// 如果需要走缓存
			if shouldUseCache {
				cachedVal, ok := GetResultFromCache(ctx, k, param)
				if ok {
					matchedCache = true
					cachedResult[k] = cachedVal
				}
			}
			// 如果缓存没取到
			if !matchedCache {
				// copyMap需要处理复用key的情况
				if resolver, ok := provider.(collisionResolver); ok {
					dmpParam = resolver.merge(param, dmpParam)
				} else {
					dmpParam = mergeMap(param, dmpParam)
				}
				paramMap[k] = param
			}
		}
	}

	var result map[string]string
	var err error
	// 需要curl dmp
	if len(paramMap) > 0 {
		// 真正需要curl dmp的key
		keyList := make([]string, 0, len(paramMap))
		for k := range paramMap {
			keyList = append(keyList, k)
		}
		params := AddCommonParams(dmpParam, eventData) // 涉及用户信息
		result, err = GetOriginFeatureValues(ctx, keyList, params)
		if nil != err {
			return nil, err
		}
		if shouldUseCache {
			// 把dmp的返回结果缓存起来
			err = StoreResult(ctx, result, paramMap)
			if nil != err {
				//logutil.AddErrorLog(ctx, logutil.MDU_DATAGRIP, logutil.IDX_DATAGRIP_DATA_STORE_FAILED, err.Error(), fmt.Sprintf("result:%+v,params:%+v", result, paramMap))
			}
		}
	}

	if shouldUseCache {
		// 把dmp返回的结果和从缓存中取的结果合并成最终的结果
		result = mergeMap(cachedResult, result)
	}
	for index, provider := range providerList {
		if commomConditionFlag == index {
			continue
		}
		provider.Bind(ctx, result, container)
	}
	return container, nil
}

// GetResultFromCache 查看某个feature以及对应的查询条件是否已经存在
// 若存在直接返回cache的结果+true，否则返回空+false
func GetResultFromCache(ctx context.Context, feature string, params map[string]string) (string, bool) {

	return string(""), true
}

// StoreResult 把dmp的返回缓存起来
// key是被用来请求dmp的key加请求参数的值
// val就是dmp的返回
func StoreResult(ctx context.Context, result map[string]string, paramMap map[string]map[string]string) (err error) {
	//traceID, err := ctxutil.GetTraceID(ctx)
	//if nil != err {
	//	return err
	//}
	//for k, v := range result {
	//	params, ok := paramMap[k]
	//	if !ok {
	//		continue
	//	}
	//	key := assembleKey(traceID, k, params)
	//	err = cache.Set(key, []byte(v), cacheExpireSeconds)
	//	if nil != err {
	//		return err
	//	}
	//}
	return err
}

func AddCommonParams(params map[string]string, eventData *_struct.EventData) map[string]string {
	if nil == params {
		params = make(map[string]string)
	}
	if nil == eventData {
		return params
	}
	//_, ok := params["pid"]
	//if !ok && eventData.PassengerInfo != nil {
	//	params["pid"] = strconv.FormatInt(eventData.PassengerInfo.PassengerId, 10)
	//}
	//_, ok = params["uid"]
	//if !ok && eventData.PassengerInfo != nil {
	//	params["uid"] = strconv.FormatInt(eventData.PassengerInfo.UserId, 10)
	//}
	//_, ok = params["product_id"]
	//if !ok && eventData.ProductInfo != nil {
	//	params["product_id"] = strconv.FormatInt(eventData.ProductInfo.OriProductId, 10)
	//}
	//_, ok = params["phone"]
	//if !ok && eventData.PassengerInfo != nil {
	//	params["phone"] = eventData.PassengerInfo.Telphone
	//}
	//_, ok = params["utc_offset"]
	//if !ok && eventData.InternationalInfo != nil {
	//	params["utc_offset"] = strconv.FormatInt(eventData.InternationalInfo.UtcOffset, 10)
	//}
	//_, ok = params["city_id"]
	//if !ok && eventData.EventInfo != nil {
	//	params["city_id"] = strconv.FormatInt(eventData.EventInfo.CityId, 10)
	//}
	//_, ok = params["is_offline"]
	//if !ok && eventData.TriggerInfo != nil {
	//	if actionmanage.IsOfflineAction(eventData) {
	//		params["is_offline"] = "1"
	//	} else {
	//		params["is_offline"] = "0"
	//	}
	//}
	return params
}

// GetOriginFeatureValues query dmp for feature values
var GetOriginFeatureValues = func(ctx context.Context, features []string, params map[string]string) (map[string]string, error) {

	//traceInfo := getTraceInfo(ctx)
	//cfg := global.GlobalConfig.PopeFs
	//if cfg.MockReturn {
	//	dt, err := getMockData(cfg.MockDataFile)
	//	logutil.AddInfoLog(ctx, logutil.MDU_DMP, logutil.IDX_DMP_MOCK, "mock", fmt.Sprintf("return:%+v", dt))
	//	//logutil.FileLog(&logutil.DebugLog{Ctx: ctx, Msg: fmt.Sprintf("mockReturn=%+v", dt)})
	//	return dt, err
	//}
	//
	//trace := &popefs.Trace{
	//	TraceID:     traceInfo.TraceID,
	//	Caller:      traceInfo.Caller,
	//	SrcMethod:   traceInfo.SrcMethod,
	//	HintCode:    traceInfo.HintCode,
	//	HintContent: traceInfo.HintContent,
	//}
	//result, err := retryGet(ctx, params, features, trace, 1)
	//if nil != err {
	//	return nil, errors.New(fmt.Sprintf("callMGetFeaturesError=%s", err))
	//}

	//resultMap := make(map[string]string)
	//for k, v := range result {
	//	resultMap[k] = v.GetVal()
	//}
	//return resultMap, nil
	return nil, nil
}

func retryGet(ctx context.Context, params map[string]string, featureKeys []string, traceInfo *_struct.Trace, level int) (map[string]*_struct.FeatureVal, error) {
	// 重试maxRetryNumber次还出错就直接报错不重试了
	//startTime := time.Now()
	//if level > 3 {
	//	return nil, errors.New(fmt.Sprintf("call dmp retry times over %d limitation", 3))
	//}
	//// 如果整个请求超时就不重试直接报错了
	//select {
	//case <-ctx.Done():
	//	return nil, errors.New(fmt.Sprintf("context timeout %s", ctx.Err()))
	//default:
	//}
	//if len(featureKeys) == 0 {
	//	return nil, nil
	//}
	//mGetParams := &popefs.MGetParams{
	//	Params:   params,
	//	Features: featureKeys,
	//}
	//
	//cSpanId := http.GenChildSpanId()
	//traceInfo.SpanID = cSpanId
	//ctx = getLegoTrace(ctx, traceInfo)
	//
	//cli, err := popefs.NewClient(global.GetPopeFsDisfName())
	//if err != nil {
	//	logutil.AddErrorLog(ctx, logutil.MDU_DMP, logutil.IDX_DMP_FEATURE_ERROR, err.Error(), params)
	//	//logutil.FileLog(&DmpFeatureFailedLog{Ctx: ctx, Type: "NewClientError", Params: params, Err: err})
	//	return nil, err
	//}
	//result, err := cli.MGetFeatures(ctx, mGetParams, traceInfo)
	//// err != nil表示根本没curl通，不重试
	//if nil != err {
	//	logutil.AddThriftFailedLog(
	//		ctx,
	//		logutil.MDU_DMP,
	//		logutil.IDX_DMP_FEATURE_ERROR,
	//		err.Error(),
	//		result,
	//		fmt.Sprint("cspanid=", cSpanId),
	//	)
	//	//logutil.FileLog(&DmpFeatureFailedLog{Ctx: ctx, Type: "Network Connection Fail", Params: params, FailedList: result, Err: err})
	//	return nil, err
	//}
	//realResultMap := result.Data.GetFeatureMap()
	//// 请求成功直接返回
	//if result.ErrNo == 0 {
	//	logutil.AddThriftSuccessLog(
	//		ctx,
	//		logutil.MDU_DMP,
	//		logutil.IDX_DMP_FEATURE,
	//		"",
	//		mGetParams,
	//		result,
	//		fmt.Sprintf("cspanid=%+v||proc_time=%+v", cSpanId, time.Since(startTime).Nanoseconds()/time.Millisecond.Nanoseconds()),
	//	)
	//	//logutil.FileLog(&DmpFeatureBizLog{Ctx: ctx, Type: "cli.MGetFeatures Response", Resp: result, Params: mGetParams, TraceInfo: *traceInfo})
	//	return realResultMap, nil
	//}
	//logutil.AddThriftFailedLog(
	//	ctx,
	//	logutil.MDU_DMP,
	//	logutil.IDX_DMP_FEATURE_ERROR,
	//	"GetFeatureMapFailedList",
	//	params,
	//	result,
	//	fmt.Sprint("cspanid=", cSpanId),
	//)
	////logutil.FileLog(&DmpFeatureFailedLog{Ctx: ctx, Type: "GetFeatureMapFailedList", Params: params, FailedList: result.Data.FailedList, Err: errors.New("GetFeatureMapFailedList")})
	//// 重试只重试失败的key，减少DMP压力
	//newResult, err := retryGet(ctx, params, result.Data.FailedList, traceInfo, level+1)
	//if nil != err {
	//	//logutil.AddErrorLog(ctx, logutil.MDU_DMP, logutil.IDX_DMP_FEATURE_ERROR, err.Error())
	//	return nil, errors.New(fmt.Sprintf("retryGetErr:%s", err))
	//}
	//// 把重试的结果和第一次的进行merge
	//for key := range newResult {
	//	realResultMap[key] = newResult[key]
	//}
	//return realResultMap, nil
	return nil, nil
}
