package comm

import (
	"GolangTrick/Engine/constant"
	"GolangTrick/Engine/service/workFlow"
	_struct "GolangTrick/Engine/struct"
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

var engineServer *EngineServer

func GetEngineServer() *EngineServer {
	return engineServer
}

type EngineServerFunc func() error
type EngineServer struct {
	initFuncs []EngineServerFunc
	Conf      *_struct.ServerConf
	signal    chan int
	workFlow  *workflow.WorkFlow
}

func NewEngineServer(conf *_struct.ServerConf) (*EngineServer, error) {
	engineServer = &EngineServer{
		initFuncs: make([]EngineServerFunc, 0, 10),
		Conf:      conf,
		workFlow:  workflow.NewWorkFlow([]string{constant.SCENE_1, constant.SCENE_2}),
		signal:    make(chan int),
	}
	err := engineServer.init()
	if err != nil {
		return nil, err
	}
	return engineServer, nil
}

func (this *EngineServer) init() error {
	err := _struct.InitConfig("/conf.toml")
	if err != nil {
		return err
	}
	this.Conf.Global = &_struct.GConfig
	return nil
}

func (this *EngineServer) watchSignal() int {
	return <-this.signal
}

func (this *EngineServer) loadWorkflow() error {
	this.loadStreaming()
	return nil
}

func (this *EngineServer) loadStreaming() {
	this.loadPsgProcess()
}

func (this *EngineServer) loadPsgProcess() error {
	//flowDiversonWorkStep := &wss.FlowDiversionWorkstep{}
	//this.workFlow.AddStep(constant.SCENE_PASSENGER, flowDiversonWorkStep)
	//
	//canvasFilterWorkStep := &wss.CanvasFilterWorkStep{}
	//this.workFlow.AddStep(constant.SCENE_PASSENGER, canvasFilterWorkStep)
	//
	//stepFilterWorkstep := &wss.StepFilterWorkstep{}
	//this.workFlow.AddStep(constant.SCENE_PASSENGER, stepFilterWorkstep)
	//
	//// 风控拦截过滤
	//this.workFlow.AddStep(constant.SCENE_PASSENGER, &wss.RiskControlWorkstep{})
	//
	//this.workFlow.AddStep(constant.SCENE_PASSENGER, &wss.PriorityCheckWorkstep{})
	//
	////	ActionRunner失败时进行重试
	//actionRunnerWorkstep := &wss.ActionRunnerWorkstep{}
	//this.workFlow.AddStep(constant.SCENE_PASSENGER, actionRunnerWorkstep)
	//
	////	由于ActionRunner目前不可能返回err，统一在response阶段处理失败
	//responseWorkstep := &wss.ResponseWorkstep{}
	//this.workFlow.AddStep(constant.SCENE_PASSENGER, responseWorkstep)

	return nil
}

func (this *EngineServer) loadInitFuncs() error {
	var err error
	for _, f := range this.initFuncs {
		err = f()
		if err != nil {
			funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
			return errors.New(fmt.Sprintf("err occurr %s, err %s", funcName, err))
		}
	}
	return nil
}

func (this *EngineServer) AddInitFuncs(f EngineServerFunc) {
	this.initFuncs = append(this.initFuncs, f)
}

func (this *EngineServer) Serv() error {
	err := this.loadWorkflow()
	if err != nil {
		return err
	}

	err = this.loadInitFuncs()
	if err != nil {
		return err
	}

	signal := this.watchSignal()
	if signal > 0 {
		return errors.New(fmt.Sprintf("signal : %d", signal))
	}
	return nil
}
