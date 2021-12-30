package comm

import (
	_struct "GolangTrick/Engine/struct"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type Controller struct {
	//根据header初始化ctx中的trace信息
	InitCtxTraceFlag bool

	TimeOut time.Duration

	//getRequestFn        func(*http.Request) (string, error)
	GetRequestFn        func(*http.Request) (*_struct.ReqData, error)
	GetRequestErrRespFn func() string
	//processFn           func(ctx context.Context, reqStr string, es *EngineServer) (*structs.Response, error)
	ProcessFn      func(ctx context.Context, reqData *_struct.ReqData, peh PushEventHandler) (*_struct.RespData, error)
	GererateRespFn func(resp *_struct.RespData, err error) string
	EngineServer   *EngineServer
}

func NewDefaultController() *Controller {
	ctl := &Controller{}
	ctl.InitCtxTraceFlag = true
	ctl.TimeOut = _struct.GConfig.Proxy.TimeOut.Duration
	ctl.GetRequestFn = GetReqFromBody
	ctl.GetRequestErrRespFn = DefaultgetRequestErrRespFn
	ctl.GererateRespFn = DefaultGererateRespFn
	ctl.ProcessFn = StandardDispatcher
	return ctl
}
