package logic1

import (
	"git.xiaojukeji.com/falcon/pope-fs/config"
	"git.xiaojukeji.com/falcon/pope-fs/model"
	"git.xiaojukeji.com/falcon/pope-fs/popefsthrift"
	"git.xiaojukeji.com/falcon/pope-fs/util/errutil"
	"git.xiaojukeji.com/gobiz/logger"
	"golang.org/x/net/context"
)

type FeatureValType int64

// Attributes:
//  - Val
//  - Type
type FeatureVal struct {
	Val  string         `thrift:"val,1,required" json:"val"`
	Type FeatureValType `thrift:"type,2,required" json:"type"`
}

// ThriftFeatureMap : 指代返回值中的 feature map，为了在开发中避免修改 map 结构导致多处修改，特抽出到 model
type ThriftFeatureMap map[string]*popefsthrift.FeatureVal

// DispatcherLogicResponse 接口的返回数据结构
type DispatcherLogicResponse struct {
	FeatureMap  model.ThriftFeatureMap // 请求成功后得到的特征列表
	FailedList  []string               // 请求失败
	NoDataList  []string               // 请求没有查到数据
	InvalidList []string               // 没有在 dmp 系统中配置该特征值
}

type featureRequest struct {
	featureList []string
	domain      string
	apiName     string
	protocol    string
	canBeMerged bool
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
			FeatureMap:  map[string]*popefsthrift.FeatureVal{},
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
