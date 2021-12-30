package _struct

import (
	"GolangTrick/Engine/constant"
	"context"
	"github.com/BurntSushi/toml"
	"net/url"
	"time"
)

type ServerConf struct {
	Common string
	Log    string
	Global *Config
}

type Duration struct {
	time.Duration
}

type Config struct {
	FilePath string
	Proxy    struct {
		HTTPPort               string
		TimeOut                Duration
		PriceTimeOut           Duration
		HTTPServerReadTimeout  Duration
		HTTPServerWriteTimeout Duration
	}
	MQconsumer struct {
		Subscribe         bool
		EventGroup        string
		EventGoroutineNum int
		EventTopic        string
		ProxyListNum      int
	}
	PopeEditor struct {
		IndexerReloadIntervalSeconds int
		LogDir                       string
		LogName                      string
	}
}

var GConfig Config

func InitConfig(configFilePath string) error {
	_, err := toml.DecodeFile(configFilePath, &GConfig)
	GConfig.FilePath = configFilePath
	return err
}

type Objects struct {
	Id   int
	Name string
	Age  int
}

type MqMessage struct {
	Key        string            `thrift:"key,1" json:"key"`
	Value      []byte            `thrift:"value,2" json:"value"`
	Tag        string            `thrift:"tag,3" json:"tag"`
	Offset     int64             `thrift:"offset,4" json:"offset"`
	Properties map[string]string `thrift:"properties,5" json:"properties"`
}

type ReqDataer interface {
	SetPayload(string)
	GetPayload() string
	GetQueryForm() url.Values
	SetQueryForm(url.Values)
	GetPostForm() url.Values
	SetPostForm(url.Values)
}

type EventData struct {
	Id         int64
	Scene      string
	CanvasInfo map[int][]int
}

func NewEventData(Id int64, scene string) *EventData {
	return &EventData{
		Id:    Id,
		Scene: scene,
	}
}
func GetEventDataFromStdParams(ctx context.Context, stdParams *StandardParamsStruct) *EventData {

	eventData := NewEventData(stdParams.Id, stdParams.Scene)

	canvasStepsMap := make(map[int][]int, 0)
	canvasStepsMap[int(123)] = []int{int(4565)} //mock值

	return eventData
}

type StandardParamsStruct struct {
	Id     int64  `json:"app_id"`
	IsSync bool   `json:"is_sync"`
	Scene  string `json:"scene"`
}

func NewStandardParams() *StandardParamsStruct {
	return &StandardParamsStruct{
		Id: constant.DEFAULT_UNSET_INT,
	}
}
