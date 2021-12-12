package Engine

import (
	"GolangTrick/Engine/comm"
	_struct "GolangTrick/Engine/struct"
	"flag"
	"fmt"
	"os"
)

var (
	env string // 程序启动环境

	ConfigPath    = flag.String("comconf", "conf/develop.toml", "common conf")
	logConfigPath = flag.String("logconf", "conf/develop.conf", "log conf")
)

var EngineServer *comm.EngineServer

func main() {
	flag.StringVar(&env, "env", "dev",
		"change runtime environment, default value: dev\n")

	flag.Parse()

	serverConf := &_struct.ServerConf{
		Common: *ConfigPath,
		Log:    *logConfigPath,
	}
	var err error

	EngineServer, err = comm.NewEngineServer(serverConf)
	//日志的初始化配置
	EngineServer.AddInitFuncs(func() error {
		//
		//cfg, err := config.New(*logConfigPath)
		//if nil != err {
		//	return err
		//}
		return err
	})
	//mq的初始化
	EngineServer.AddInitFuncs(func() error {
		//err = mq.InitMQ()
		return err
	})

	EngineServer.AddInitFuncs(func() error {
		loader, err := bizloader.Init(EngineServer.Conf.Global)
		if err != nil {
			return err
		}
		return loader.Serv()
	})

	//初始化mq回调函数
	EngineServer.AddInitFuncs(func() error {
		//err = mqconsumer.InitConsumer()
		return err
	})

	EngineServer.AddInitFuncs(func() error {
		//errCh := make(chan error)
		go func() {
			comm.ServeHTTP()
			//errCh <- biz_frame.ServeHTTP()
		}()
		return nil
	})

	err = EngineServer.Serv()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
