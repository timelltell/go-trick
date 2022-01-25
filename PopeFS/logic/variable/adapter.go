package variable1

import (
	drv_points1 "GolangTrick/PopeFS/logic/drv_points"
	driverAdapter1 "GolangTrick/PopeFS/logic/variable/driverAdapter"
	model1 "GolangTrick/PopeFS/model"
	"git.xiaojukeji.com/falcon/pope-fs/config"
	"git.xiaojukeji.com/falcon/pope-fs/logic/adapter/variable/driverAdapter"
	"git.xiaojukeji.com/falcon/pope-fs/util/datautil"
	"git.xiaojukeji.com/falcon/pope-fs/util/errutil"
	"golang.org/x/net/context"
	"sync"
)

const CacheKey = "variable_feature_%s_%s_%v_%v"

// 变量处理层,从feature中提取变量做占位符
type Adapter struct {
	Virtual
	adapter *driverAdapter1.Driver
	ctx     context.Context
}

func GetVarQuery(ctx context.Context, keys []string, params map[string]string) (model1.ThriftFeatureMap, []string, error) {
	instance := &Adapter{}
	instance.Init(ctx, keys, params)

	// get from redis
	cacheResult := make(model1.ThriftFeatureMap)
	for i := 0; i < len(instance.Keys); i++ {
		//varKey := instance.Keys[i]
		//rediskey := fmt.Sprintf(CacheKey, params["driver_id"], varKey, params["canvas_id"], params["step_id"])
		//cacheRes, _ := common.Get(ctx, rediskey)
		//if cacheRes != nil {
		//	cacheResult[varKey] = &model1.FeatureVal{Val: string(cacheRes.([]byte))}
		//	instance.Keys = append(instance.Keys[:i], instance.Keys[i+1:]...)
		//	i--
		//}
	}

	// get from feautre
	resFeature, err := instance.Run(instance)
	noDatalist := []string{}
	if len(resFeature) == 0 {
		noDatalist = instance.Keys
	}

	//for varKey, v := range resFeature {
	//rediskey := fmt.Sprintf(CacheKey, params["driver_id"], varKey, params["canvas_id"], params["step_id"])
	//common.Set(ctx, rediskey, v.Val, 40)
	//}

	// merge redis and data
	for k, v := range cacheResult {
		resFeature[k] = v
	}

	return resFeature, noDatalist, err
}

// Init init
func (o *Adapter) Init(ctx context.Context, keys []string, params map[string]string) {
	o.adapter = driverAdapter1.New()
	o.Keys = o.adapter.VarKey(keys)
	o.Response = &Response{}
	o.Request = &Request{}
	o.Params = params
	o.Symbol = params["symbol"]
	o.ctx = ctx
}

func (o *Virtual) Run(p VirtualI) (model1.ThriftFeatureMap, error) {

	if err := p.BeforeRead(); err != nil {
		return model1.ThriftFeatureMap{}, err
	}

	res, err := p.Read()
	if err != nil {
		return model1.ThriftFeatureMap{}, err
	}
	return p.AfterRead(res)
}

func (o *Adapter) BeforeRead() (err error) {
	err = o.Request.PreBuildParams(o.Params)
	if err != nil {
		return err
	}

	// 整合TCA的上下文
	o.Params, err = o.Request.BuildRequestParams(o.adapter, o.Keys)

	//err = o.adapter.CheckParam(o.Keys, o.Params)
	if err != nil {
		return err
	}

	return
}

func (o *Adapter) Read() (feature model1.ThriftFeatureMap, err error) {

	// 获取变量的值
	featureList := o.adapter.GetRealFeature(o.Keys)
	ret, err := Dispatch(o.ctx, featureList, o.Params)

	// todo drive_id + cavans_id + step_id 变量作用域

	// o.Response
	return ret.FeatureMap, nil
}

func (o *Adapter) AfterRead(featureMap model1.ThriftFeatureMap) (model1.ThriftFeatureMap, error) {

	o.Response.ParseResponse(featureMap, o.Keys)
	o.Response.Logic(featureMap)

	// sepcial: 从action获取i18n文案 todo elvish-cgo比较臃肿，暂时没有接入
	if _, ok := featureMap[driverAdapter.REWARD_MONEY_ORDER]; ok {
		featureMap[driverAdapter.REWARD_MONEY_ORDER].Val = o.Symbol + featureMap[driverAdapter.REWARD_MONEY_ORDER].Val
	}
	if _, ok := featureMap[driverAdapter.REWARD_MONEY_ORDER_STREAM]; ok {
		featureMap[driverAdapter.REWARD_MONEY_ORDER_STREAM].Val = o.Symbol + featureMap[driverAdapter.REWARD_MONEY_ORDER_STREAM].Val
	}

	return featureMap, nil
}

// Dispatch 分配务给 adapter
func Dispatch(ctx context.Context, featureList []string, params map[string]string) (*DispatcherLogicResponse, error) {
	requestList := []featureRequest{}
	notInConfigList := []string{}
	requestGroupMap := make(map[string]*featureRequest)
	if len(featureList) == 0 {
		return &DispatcherLogicResponse{
			FeatureMap:  map[string]*model1.FeatureVal{},
			FailedList:  []string{},
			NoDataList:  []string{},
			InvalidList: notInConfigList,
		}, nil
	}

	var groupDataErr *errutil.ErrorInfo
	for _, featureName := range featureList {
		metaData, getMetaErr := config.GetFeatureMeta(featureName)
		if getMetaErr != nil {
			notInConfigList = append(notInConfigList, featureName)
			groupDataErr = errutil.New(errutil.ErrnoPartlySuccess, getMetaErr.Error())
			continue
		}

		if !metaData.CanBeMerged {
			singleFeatureGroup := featureRequest{
				featureList: []string{featureName},
				domain:      metaData.Domain,
				apiName:     metaData.APIName,
				protocol:    metaData.Protocol,
				canBeMerged: false,
			}
			requestList = append(requestList, singleFeatureGroup)
			continue
		}

		requestKey := metaData.APIName + metaData.Domain + metaData.Protocol
		if group, ok := requestGroupMap[requestKey]; ok {
			group.featureList = append(group.featureList, featureName)
		} else {
			requestGroupMap[requestKey] = &featureRequest{
				featureList: []string{featureName},
				domain:      metaData.Domain,
				apiName:     metaData.APIName,
				protocol:    metaData.Protocol,
				canBeMerged: false,
			}
		}
	}

	for _, req := range requestGroupMap {
		req.featureList = datautil.UniqueStringSlice(req.featureList)
		requestList = append(requestList, *req)
	}

	featureMap, failedList, noDataList, err := dispatchCall(ctx, requestList, params)

	// 如果在分组数据时有错，那么优先返回分组错误)
	resp := &DispatcherLogicResponse{
		FeatureMap:  featureMap,
		FailedList:  failedList,
		NoDataList:  noDataList,
		InvalidList: notInConfigList,
	}

	if groupDataErr != nil && len(featureMap) > 0 {
		return resp, groupDataErr
	}

	return resp, err

}

type featureRequest struct {
	featureList []string
	domain      string
	apiName     string
	protocol    string
	canBeMerged bool
}

type featureAdapterResponse struct {
	FeatureMap model1.ThriftFeatureMap // 合法的 feature map
	FailedList []string                // 请求失败
	NoDataList []string                // 请求查不到相应的特征数据
}

// DispatcherLogicResponse 接口的返回数据结构
type DispatcherLogicResponse struct {
	FeatureMap  model1.ThriftFeatureMap // 请求成功后得到的特征列表
	FailedList  []string                // 请求失败
	NoDataList  []string                // 请求没有查到数据
	InvalidList []string                // 没有在 dmp 系统中配置该特征值
}

// get 接口对下游的函数签名约束
type adapter func(context.Context, []string, map[string]string) (model1.ThriftFeatureMap, []string, error)

// domain => api_name => protocol => adapter
var adapterMap = map[string]map[string]map[string]adapter{
	"driver_points": {
		"get": {
			"http": drv_points1.GetDriverPoints,
		},
	},
}

func getAdapter(domain string, apiName string, protocol string) adapter {
	if _, ok := adapterMap[domain]; !ok {
		return nil
	}
	innerMap := adapterMap[domain]

	if _, ok := innerMap[apiName]; !ok {
		return nil
	}
	nestedInnerMap := innerMap[apiName]

	if _, ok := nestedInnerMap[protocol]; !ok {
		return nil
	}

	return nestedInnerMap[protocol]
}

// dispatchCall 并发请求 feature 结果
func dispatchCall(ctx context.Context, requestList []featureRequest, params map[string]string) (model1.ThriftFeatureMap, []string, []string, error) {
	var finalFeatureMap = model1.ThriftFeatureMap{}
	var finalNoDataList, finalFailedList = []string{}, []string{}
	var err error

	if len(requestList) == 0 {
		err = errutil.New(errutil.ErrnoEmptyRequestList, "empty feature request list")
		return finalFeatureMap, finalFailedList, finalNoDataList, err
	}

	var wg = &sync.WaitGroup{}
	wg.Add(len(requestList))
	var resultChan = make(chan featureAdapterResponse, 16)

	for _, req := range requestList {
		req := req
		go func() {
			defer func() {
				if err := recover(); err != nil {
					resultChan <- featureAdapterResponse{FailedList: req.featureList}
				}
				wg.Done()
			}()

			var noDataList []string
			h := getAdapter(req.domain, req.apiName, req.protocol)
			if h == nil {
				resultChan <- featureAdapterResponse{FailedList: req.featureList}
				return
			}

			featureMap, noDataList, fetchDataErr := h(ctx, req.featureList, DeepCopyMap(params))
			if fetchDataErr != nil {
				resultChan <- featureAdapterResponse{FailedList: req.featureList}
			} else {

				resultChan <- featureAdapterResponse{FailedList: []string{}, NoDataList: noDataList, FeatureMap: featureMap}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for res := range resultChan {
		for k, v := range res.FeatureMap {
			finalFeatureMap[k] = v
		}
		finalFailedList = append(finalFailedList, res.FailedList...)
		finalNoDataList = append(finalNoDataList, res.NoDataList...)
	}

	if len(finalFailedList)+len(finalNoDataList) > 0 && len(finalFeatureMap) > 0 {
		err = errutil.New(errutil.ErrnoPartlySuccess, "feature request partly success")
	}

	if len(finalFeatureMap) == 0 {
		err = errutil.New(errutil.ErrnoFeatureAllFailed, "feature request all failed")
	}

	return finalFeatureMap, finalFailedList, finalNoDataList, err
}

// DeepCopyMap map copy
func DeepCopyMap(data map[string]string) map[string]string {
	dict := make(map[string]string)
	for key, value := range data {
		dict[key] = value
	}
	return dict
}
