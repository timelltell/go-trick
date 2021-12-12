package workstep

import (
	"GolangTrick/Engine/service/workflow"
	"context"
)

type BsFilterWorkStep struct {
}

func (this *BsFilterWorkStep) Key() string {
	return "BSFilter"
}

func (this *BsFilterWorkStep) Handle(context.Context, workflow.WorkStepData) (workflow.WorkStepData, error, workflow.GoNext) {

}
