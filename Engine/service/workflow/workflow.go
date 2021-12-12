package workflow

import (
	_struct "GolangTrick/Engine/struct"
	"context"
	"errors"
)

type WorkStepData struct {
	Objects   []_struct.Objects
	OpStreams interface{}
	EventData _struct.EventData
}

type GoNext string

const (
	GOTO_NEXT = "next"
)

type WorkSteper interface {
	Key() string
	Handle(context.Context, WorkStepData) (WorkStepData, error, GoNext)
}

type WorkFlow struct {
	sceneWorksteps     map[string][]WorkSteper
	sceneWorkstepIndex map[string]map[string]int
	headSteps          map[string]WorkSteper
}

func NewWorkFlow(scenes []string) *WorkFlow {
	this := &WorkFlow{
		make(map[string][]WorkSteper),
		make(map[string]map[string]int),
		make(map[string]WorkSteper),
	}
	for _, scene := range scenes {
		this.sceneWorksteps[scene] = make([]WorkSteper, 0, 5)
		this.sceneWorkstepIndex[scene] = make(map[string]int)
		this.headSteps[scene] = nil
	}
	return this
}

func (this *WorkFlow) sceneExists(scene string) bool {
	_, ok1 := this.sceneWorksteps[scene]
	_, ok2 := this.sceneWorkstepIndex[scene]
	_, ok3 := this.headSteps[scene]
	if ok1 && ok2 && ok3 {
		return true
	}
	return false
}

func (this *WorkFlow) HeadStep(scene string) WorkSteper {
	if this.sceneExists(scene) {
		headstep, _ := this.headSteps[scene]
		return headstep
	}
	return nil
}

func (this *WorkFlow) AddStep(scene string, step WorkSteper) error {
	if this.sceneExists(scene) {
		_, ok := this.headSteps[scene]
		if !ok || len(this.sceneWorksteps[scene]) == 0 {
			this.headSteps[scene] = step
		}
		this.sceneWorkstepIndex[scene][step.Key()] = len(this.sceneWorksteps[scene])
		this.sceneWorksteps[scene] = append(this.sceneWorksteps[scene], step)
		return nil
	}
	return errors.New("not scene")
}

func (this *WorkFlow) Next(scene string, stepKey string) WorkSteper {
	if this.sceneExists(scene) {
		index, ok := this.sceneWorkstepIndex[scene][stepKey]
		if ok {
			nextSeq := index + 1
			if nextSeq >= len(this.sceneWorksteps[scene]) {
				return nil
			}
			return this.sceneWorksteps[scene][nextSeq]
		}
	}
	return nil
}

func (this *WorkFlow) Goto(scene string, stepKey GoNext) WorkSteper {
	if this.sceneExists(scene) {
		seq, ok := this.sceneWorkstepIndex[scene][string(stepKey)]
		if ok && seq < len(this.sceneWorksteps[scene]) {
			return this.sceneWorksteps[scene][seq]
		}
	}
	return nil
}
