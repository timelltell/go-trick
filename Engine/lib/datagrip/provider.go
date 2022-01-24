package datagrip

import (
	_struct "GolangTrick/Engine/struct"
	"github.com/codegangsta/inject"
	"golang.org/x/net/context"
)

type providerConstructor func(ConstructParam) DataProvider

// ConstructParam 用于Provider的实例化
type ConstructParam struct {
	EventData  *_struct.EventData
	ObjectList []_struct.Object
	StepList   []_struct.Step
}

const (
	drvSurplusPointsKey = "pope.drv_surplus_points"
)

var (
	stepRelatedConstructor = map[string]providerConstructor{
		drvSurplusPointsKey: NewDrvSurplusPointsProvider,
	}
)

type drvSurplusPointsProvider struct {
	eventData *_struct.EventData
	stepList  []_struct.Step
}

func (p *drvSurplusPointsProvider) Match(key string) bool {
	return key == drvSurplusPointsKey
}

func (p *drvSurplusPointsProvider) Prepare(ctx context.Context) ([]string, map[string]string, error) {

	return []string{drvSurplusPointsKey}, map[string]string{}, nil
}

func (p *drvSurplusPointsProvider) Bind(ctx context.Context, result map[string]string, container inject.TypeMapper) {

	container.Map(0)
}

func (p *drvSurplusPointsProvider) merge(param, dmpParam map[string]string) map[string]string {
	return MergeCouponList(param, dmpParam)
}

// NewCouponMaxAmountProvider 获取司机剩余积分的provider
func NewDrvSurplusPointsProvider(queryParam ConstructParam) DataProvider {
	return &drvSurplusPointsProvider{eventData: queryParam.EventData, stepList: queryParam.StepList}
}

func mergeMap(src, dst map[string]string) map[string]string {
	if dst == nil {
		dst = make(map[string]string)
	}
	if src == nil {
		return dst
	}
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
func MergeCouponList(param, dmpParam map[string]string) map[string]string {
	_, ok := dmpParam["coupon_list"]
	if !ok {
		return mergeMap(param, dmpParam)
	}
	if !ok {
		return mergeMap(param, dmpParam)
	}
	dmpParam = mergeMap(param, dmpParam)
	return dmpParam
}
