package logic1

import (
	model1 "GolangTrick/PopeFS/model"
	"git.xiaojukeji.com/falcon/pope-fs/config"
	"git.xiaojukeji.com/falcon/pope-fs/util/errutil"
	"git.xiaojukeji.com/gobiz/logger"
	"golang.org/x/net/context"
	"sync"
)

type featureAdapterResponse struct {
	FeatureMap model1.ThriftFeatureMap // 合法的 feature map
	FailedList []string                // 请求失败
	NoDataList []string                // 请求查不到相应的特征数据
}

type featureRequest struct {
	featureList []string
	domain      string
	apiName     string
	protocol    string
	canBeMerged bool
}

// DispatcherLogicResponse 接口的返回数据结构
type DispatcherLogicResponse struct {
	FeatureMap  model1.ThriftFeatureMap // 请求成功后得到的特征列表
	FailedList  []string                // 请求失败
	NoDataList  []string                // 请求没有查到数据
	InvalidList []string                // 没有在 dmp 系统中配置该特征值
}

// Dispatch 分配务给 adapter
func Dispatch(ctx context.Context, featureList []string, params map[string]string) (*DispatcherLogicResponse, error) {
	//domain := config.GetFeatureMeta()
	// group featureList by domain，can_be_merged，api_name，protocol
	requestList := []featureRequest{}
	notInConfigList := []string{}
	requestGroupMap := make(map[string]*featureRequest)
	// Step-1 若请求 featureList 为空，则直接返回
	if len(featureList) == 0 {
		return &DispatcherLogicResponse{
			FeatureMap:  map[string]*model1.FeatureVal{},
			FailedList:  []string{},
			NoDataList:  []string{},
			InvalidList: notInConfigList,
		}, nil
	}

	// Step-2 获取每个 feature 的配置信息
	// 1. 找不到的放到 notInConfigList
	// 1. 若 CanBeMerged = true，合并到 requestGroupMap[APIName+Domain+Protocol]，否则直接放到 requestList
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

		// get+tag+http
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

	// Step-3 遍历 requestGroupMap 追加到 requestList
	for _, req := range requestGroupMap {
		requestList = append(requestList, *req)
	}

	// Step-4 发起请求
	featureMap, failedList, noDataList, err := dispatchCall(ctx, requestList, params)

	// Step-5 返回请求
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

	// channel 的 buffer 大小应该动态调整
	var resultChan = make(chan featureAdapterResponse, 30)

	// 启动 len(requestList) 个 goroutine 发起请求
	for _, req := range requestList {
		req := req
		go func() {
			// 确保 wg.Done() 能执行到，当请求出现错误时，确保向 resultChan 发送数据
			defer func() {
				if err := recover(); err != nil {
					resultChan <- featureAdapterResponse{FailedList: req.featureList}
					//logger.Errorf(ctx, logger.DLTagUndefined, "fetch feature: %v panic||errno:%d||errmsg:%v||stack:%s",
					//	req.featureList, errutil.ErrPanic, err, string(debug.Stack()))
				}
				wg.Done()
			}()

			var noDataList []string
			// 获取相应的 adapter
			h := getAdapter(req.domain, req.apiName, req.protocol)
			if h == nil {
				logger.Errorf(ctx, logger.DLTagUndefined, "feature adapter not found from getAdapter %#v, check adapter config file? domain:%v, apiname:%v, protocol:%v", req, req.domain, req.apiName, req.protocol)

				resultChan <- featureAdapterResponse{FailedList: req.featureList}

				return
			}
			logger.Debugf(ctx, logger.DLTagUndefined, "#### get adapter handler: %#v ####", h)
			p := DeepCopyMap(params)
			// 发起请求
			featureMap, noDataList, fetchDataErr := h(ctx, req.featureList, p)
			if fetchDataErr != nil {
				logger.Warnf(ctx, logger.DLTagUndefined, "%v_%v", req.domain, fetchDataErr)
				resultChan <- featureAdapterResponse{FailedList: req.featureList}
				//monitor.GetStatistic().Counter("err_Count", map[string]string{"req.domain": req.domain})
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
