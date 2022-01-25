package logic1

import (
	variable1 "GolangTrick/PopeFS/logic/variable"
	model1 "GolangTrick/PopeFS/model"
	"golang.org/x/net/context"
)

// domain => api_name => protocol => adapter
var adapterMap map[string]map[string]map[string]adapter

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

// get 接口对下游的函数签名约束
type adapter func(context.Context, []string, map[string]string) (model1.ThriftFeatureMap, []string, error)

// InitAdapters 初始化各种适配器
func InitAdapters() {
	adapterMap = map[string]map[string]map[string]adapter{
		"driver_variable": {
			"get": {
				"http": variable1.GetVarQuery,
			},
		},
	}
}
