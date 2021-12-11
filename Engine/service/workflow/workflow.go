package workflow

import (
	_struct "GolangTrick/Engine/struct"
	"context"
)
type WorkStepData struct {
	Objects []_struct.Objects
	OpStreams interface{}
}

type GoNext string

type WorkStep interface {
	Key() string
	Handle(context.Context,WorkStepData)(WorkStepData,error,GoNext)
}

type WorkFlow struct {
	scenceWorksteps map[string][]WorkStep
	headSteps map[string]WorkStep
}

func NewWorkflow(scenes []string) *WorkFlow {
	wf := &WorkFlow{
		scenceWorksteps:   make(map[string][]WorkStep),
		headSteps:        make(map[string]WorkStep),
	}
	for _, scene := range scenes {
		wf.scenceWorksteps[scene] = make([]WorkStep, 0, 5)
		wf.headSteps[scene] = nil
	}
	return wf
}