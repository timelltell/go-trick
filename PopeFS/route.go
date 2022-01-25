package PopeFS

import (
	controller1 "GolangTrick/PopeFS/controller"
	"git.xiaojukeji.com/falcon/pope-fs/middleware"
	"github.com/go-chi/chi"
)

func initHandler(mux *chi.Mux) {

	// 添加 middleware
	// Header middleware 必须放在第一个
	mux.Use(middleware.HeaderSetter)
	mux.Use(middleware.RateLimit)

	mux.Handle("/mget", controller1.HTTPHandler("mget", controller1.MGetHandler))

}
