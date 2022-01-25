package PopeFS

import (
	"fmt"
	"git.xiaojukeji.com/falcon/pope-fs/util/requtil"
	"git.xiaojukeji.com/gobiz/config"
	"git.xiaojukeji.com/gobiz/logger"
	"github.com/go-chi/chi"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"syscall"
)

// initLogger 日志初始化
func initLogger(file string) error {
	fmt.Println("[pope-fs]server_start||init logger")
	// 绑定生成 LogID 的方法 参数从 context 中获取 ，于 middleware.header 中存储
	genLogHeader := func(ctx context.Context) string {
		r := requtil.GetLogRecord(ctx)
		return r.String()
	}

	configer, err := config.New(file)
	if err != nil {
		return err
	}

	err = logger.NewLoggerWithConfig(configer)
	if err != nil {
		return err
	}
	logger.RegisterContextFormat(genLogHeader)
	return nil
}

/*
	initHTTPServer : 初始化 web server
*/
func initHTTPServer(port string) error {
	//logger.Info("server_start", "init http server")
	fmt.Println("[pope-fs]server_start||init http server")
	mux := chi.NewMux()
	if port[0] != ':' {
		port = ":" + port
	}
	initHandler(mux)
	pid := os.Getpid()
	go func() {
		if err := http.ListenAndServe(port, mux); err != nil {
			logger.Fatal("init", "start_http_server_failed", err.Error())
			process := os.Process{Pid: pid}
			err := process.Signal(syscall.SIGINT)
			if err != nil {
				logger.Fatal("init", "shutdown_http_server_failed")
			}
		}
	}()
	return nil
}
