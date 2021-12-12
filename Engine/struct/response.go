package _struct

import "context"

type Response struct {
	ErrNo   int64                  `json:"errno"`
	ErrMsg  string                 `json:"errmsg"`
	Data    interface{}            `json:"data,omitempty"`
	Context map[string]interface{} `json:"context,omitempty"`
}

type RespData struct {
	Ctx  context.Context
	Id   int64
	Resp *Response
}
