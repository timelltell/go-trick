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
		IndexerMode                  string
		IndexerFile                  string
	}
}

var GConfig Config

func InitConfig(configFilePath string) error {
	_, err := toml.DecodeFile(configFilePath, &GConfig)
	GConfig.FilePath = configFilePath
	return err
}

type Object struct {
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
	Id               int64
	Scene            string
	Age              int
	CanvasInfo       map[int][]int
	BatchRequestFlag bool
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

type PopeIndexerResponse struct {
	Errcode int       `json:"errcode"`
	Errmsg  string    `json:"errmsg"`
	Data    []*Object `json:"data"`
}

func NewPopeIndexerResponse() *PopeIndexerResponse {
	return &PopeIndexerResponse{}
}

type TriggerParam struct {
	Type string
	Data interface{}
}
type Trigger struct {
	TriggerId int
	Key       string
	Params    map[string]TriggerParam
}

type ConditionParam struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
type Condition struct {
	Id        int                       `json:"condition_id"`
	Key       string                    `json:"key"`
	Expr      string                    `json:"expr"`
	Params    map[string]ConditionParam `json:"params"`
	ValueType string                    `json:"value_type"`
	Value     interface{}               `json:"value_data"`
}
type Action struct {
	Id               int
	Key              string
	Params           map[string]interface{}
	NeedResponseData int  `json:"NeedResponseData,omitempty"`
	Async            int  `json:"Async,omitempty"`
	Virtual          bool `json:"Virtual,omitempty"`
}

type ActionParam struct {
	Type string
	Data interface{}
}
type Step struct {
	Id             int
	ForwardStepIds []int
	PreviousStepId int
	Trigger        Trigger
	Conditions     []Condition
	Actions        []Action
	LastFlag       int
}
type ConditionData struct {
	Id          int
	Key         string
	Expr        string
	ValueType   string
	TargetValue interface{}
	RealValue   interface{}
	RealType    string
}

type UserOPStream struct {
	Uid    int64
	EvData EventData
	//CanvasId      int
	Object        Object
	Step          Step //will be replaced by Steps
	Steps         []Step
	ConditionData []ConditionData
}

// Attributes:
//  - TraceID
//  - Caller
//  - SpanID
//  - SrcMethod
//  - HintCode
//  - HintContent
type Trace struct {
	TraceID     string `thrift:"traceID,1,required" json:"traceID"`
	Caller      string `thrift:"caller,2,required" json:"caller"`
	SpanID      string `thrift:"spanID,3,required" json:"spanID"`
	SrcMethod   string `thrift:"srcMethod,4,required" json:"srcMethod"`
	HintCode    string `thrift:"hintCode,5,required" json:"hintCode"`
	HintContent string `thrift:"hintContent,6,required" json:"hintContent"`
}

//司机剩余积分
type TargetConditionDrvPoints struct {
	DrvPoints int
}
