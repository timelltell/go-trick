package comm

import (
	_struct "GolangTrick/Engine/struct"
	"golang.org/x/net/context"
	"net/http"
)
import "github.com/urfave/negroni"

type Server struct {
	Mux          *http.ServeMux
	N            *negroni.Negroni
	S            *http.Server
	EngineServer *EngineServer
}

func (s *Server) AddRoute(path string, ctl *Controller) {
	s.AddRouteWithScene(path, ctl, "")
}

func (s *Server) AddRouteWithScene(path string, ctl *Controller, logScene string) {
	if ctl == nil || ctl.GererateRespFn == nil || ctl.GetRequestErrRespFn == nil || ctl.ProcessFn == nil {
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), ctl.TimeOut)
		defer cancel()

		reqData, err := ctl.GetRequestFn(r)
		if err != nil {
			respStr := ctl.GetRequestErrRespFn()
			//logutil.AddRequestOutLog(
			//	ctx,
			//	logutil.MDU_BIZ_FRAME,
			//	logutil.IDX_HTTP_RESPONSE,
			//	fmt.Sprintf("get raw query error: %s", err.Error()),
			//	0,
			//	"",
			//	"http",
			//	respStr,
			//	time.Since(startTime).Nanoseconds()/time.Millisecond.Nanoseconds(),
			//	true,
			//	r.URL,
			//)
			w.Write([]byte(respStr))
			return
		}

		defer func() {
			if err := recover(); err != nil {
				//metricsutil.Add("panic", 1)
				//logutil.AddErrorLog(ctx, logutil.MDU_BIZ_FRAME, logutil.IDX_MQ_CONSUMER_FAILED, "standard dispatcher panic", string(debug.Stack()), err)
			}
		}()
		// 设置httprequest参数
		//ctx = ctxutil.SetHTTPRequest(ctx, r)

		resp, err := ctl.ProcessFn(ctx, reqData, PushEvent)
		respStr := ctl.GererateRespFn(resp, err)

		//非预估的请求为了方便在把脉查询记录下traceid，pid，phone
		if resp == nil {
			resp = &_struct.RespData{
				Ctx: ctx,
			}
		}

		w.Write([]byte(respStr))
	}

	s.Mux.HandleFunc(path, handler)
}
