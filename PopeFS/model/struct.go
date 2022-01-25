package model1

type FeatureValType int64

// Attributes:
//  - Val
//  - Type
type FeatureVal struct {
	Val  string         `thrift:"val,1,required" json:"val"`
	Type FeatureValType `thrift:"type,2,required" json:"type"`
}

// ThriftFeatureMap : 指代返回值中的 feature map，为了在开发中避免修改 map 结构导致多处修改，特抽出到 model
type ThriftFeatureMap map[string]*FeatureVal

// DispatcherLogicResponse 接口的返回数据结构
type DispatcherLogicResponse struct {
	FeatureMap  ThriftFeatureMap // 请求成功后得到的特征列表
	FailedList  []string         // 请求失败
	NoDataList  []string         // 请求没有查到数据
	InvalidList []string         // 没有在 dmp 系统中配置该特征值
}

type FeatureRequest struct {
	featureList []string
	domain      string
	apiName     string
	protocol    string
	canBeMerged bool
}

type featureAdapterResponse struct {
	FeatureMap ThriftFeatureMap // 合法的 feature map
	FailedList []string         // 请求失败
	NoDataList []string         // 请求查不到相应的特征数据
}
