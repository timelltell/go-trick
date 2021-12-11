package _struct

import (
	"time"
	"github.com/BurntSushi/toml"
)

type Duration struct {
	time.Duration
}

type Config struct {
	FilePath    string
	Proxy struct {
		HTTPPort               string
		TimeOut                Duration
		PriceTimeOut           Duration
		HTTPServerReadTimeout  Duration
		HTTPServerWriteTimeout Duration
	}
}

var GConfig Config

func InitConfig(configFilePath string) error{
	_,err:=toml.DecodeFile(configFilePath,&GConfig)
	GConfig.FilePath = configFilePath
	return err
}

type Objects struct {
	name string
	age int
}