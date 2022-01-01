package scene

import (
	"GolangTrick/Engine/dao"
	"GolangTrick/Engine/service/workstep/common"
	"GolangTrick/Engine/service/workstep/common/psg"
	_struct "GolangTrick/Engine/struct"
	"errors"
	"golang.org/x/net/context"
)

func GenCommonBSInstance(ctx context.Context, bsInstance *common.Bs, eventData _struct.EventData) {
	bsInstance.SetSourceObjects(ctx, nil)
	bsInstance.AddFilter(psg.NewAgeFilter((eventData.Age), dao.GetInstance()))
}

func FilterByBS(ctx context.Context, eventData _struct.EventData) ([]_struct.Object, error) {
	bsInstance := common.NewBS()
	GenCommonBSInstance(ctx, bsInstance, eventData)
	canvases, err := bsInstance.MatchedObjects(ctx)
	if err != nil {
		return nil, errors.New("1")
	}
	return canvases, nil

}
