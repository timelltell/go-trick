package comm

import (
	_struct "GolangTrick/Engine/struct"
	"github.com/urfave/negroni"
	"net/http"
)

func ServeHTTP() error {
	es := GetEngineServer()
	serverInstance := &Server{
		Mux:          http.NewServeMux(),
		N:            negroni.New(),
		EngineServer: es,
	}

	//pope 对外统一接口
	standardApiCtl := NewDefaultController()
	standardApiCtl.EngineServer = es
	serverInstance.AddRoute("/gulfstream/popeproxy/standard", standardApiCtl)

	recovery := negroni.NewRecovery()
	serverInstance.N.Use(recovery)
	//recovery.PanicHandlerFunc = ReportToLogger
	recovery.PrintStack = false

	serverInstance.N.UseHandler(serverInstance.Mux)

	serverInstance.S = &http.Server{
		Addr:         _struct.GConfig.Proxy.HTTPPort,
		Handler:      serverInstance.N,
		ReadTimeout:  _struct.GConfig.Proxy.HTTPServerReadTimeout.Duration,
		WriteTimeout: _struct.GConfig.Proxy.HTTPServerWriteTimeout.Duration,
	}

	return serverInstance.S.ListenAndServe()
}
