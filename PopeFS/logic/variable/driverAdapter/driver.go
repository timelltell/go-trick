package driverAdapter1

import (
	"strings"
)

// 司机的占位符逻辑
type Driver struct {
	Mapping      map[string]string
	ParamRequire map[string]map[string]string
}

// Mapping: 变量key和feature一对多关系
var KeyFeature = map[string]string{
	DRV_COLLECTED_POINTS: FEATURE_DR_POINTS,
	DRV_SURPLUS_POINTS:   FEATURE_DR_POINTS,
}

// feature所需参数
var KeyParams = map[string]map[string]string{
	FEATURE_DR_POINTS: map[string]string{"uid": "uid"},
}

func New() (adapter *Driver) {
	adapter = &Driver{}
	adapter.init(KeyFeature, KeyParams)
	return adapter
}

// Mapping: 变量feature和key多对一关系
// Param:   feature所需参数
func (Adapter *Driver) init(mapping map[string]string, paramRequire map[string]map[string]string) {
	Adapter.Mapping = mapping
	Adapter.ParamRequire = paramRequire
}

func (Adapter *Driver) VarKey(keys []string) (res []string) {
	res = make([]string, 0)
	for _, key := range keys {
		varkey := strings.Split(key, ".")
		if len(varkey) > 2 && Adapter.Mapping[varkey[2]] != "" {
			res = append(res, varkey[2])
		}
	}
	res = UniqueStringSlice(res)
	return res
}

func (Adapter *Driver) GetRealFeature(varkeys []string) (feature []string) {
	feature = make([]string, 0)
	for _, key := range varkeys {
		if _, ok := Adapter.Mapping[key]; ok {
			feature = append(feature, Adapter.Mapping[key])
		}
	}
	return
}

func UniqueStringSlice(ss []string) []string {
	var newStrSlice []string
	strMap := make(map[string]int)
	for _, s := range ss {
		strMap[s] = 1
	}
	for k, _ := range strMap {
		newStrSlice = append(newStrSlice, k)
	}
	return newStrSlice
}
