package comm

import "golang.org/x/net/context"

type Response struct {
	ErrNo   int64                  `json:"errno"`
	ErrMsg  string                 `json:"errmsg"`
	Data    interface{}            `json:"data,omitempty"`
	Context map[string]interface{} `json:"context,omitempty"`
}

type RespData struct {
	Ctx       context.Context
	Pid       int64
	Did       int64
	Phone     string // psg phone
	DrvPhone  string // drv phone
	TriggerId int64
	Resp      *Response
}
