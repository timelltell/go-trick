package psg

import (
	"GolangTrick/Engine/dao"
	"GolangTrick/Engine/service/workflow"
	"GolangTrick/Engine/service/workstep/common"
	"GolangTrick/Engine/service/workstep/common/psg"
	_struct "GolangTrick/Engine/struct"
	"context"
	"errors"
)

type BsFilterWorkStep struct {
}

func (this *BsFilterWorkStep) Key() string {
	return "BSFilter"
}

func (this *BsFilterWorkStep) Handle(ctx context.Context, in workflow.WorkStepData) (workflow.WorkStepData, error, string) {
	out := in
	canvases, err := FilterByBS(ctx, in.EventData)
	if err != nil {
		return out, errors.New("next"), workflow.GOTO_NEXT
	}
	out.Objects = canvases

	return out, nil, workflow.GOTO_NEXT
}

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
