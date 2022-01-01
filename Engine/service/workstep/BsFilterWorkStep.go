package workstep

import (
	"GolangTrick/Engine/service/workflow"
	"GolangTrick/Engine/service/workstep/scene"
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
	canvases, err := scene.FilterByBS(ctx, in.EventData)
	if err != nil {
		return out, errors.New("next"), "next"
	}
	out.Objects = canvases

	return out, nil, "next"
}
