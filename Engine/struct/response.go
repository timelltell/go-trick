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

type FeatureValType int64

// Attributes:
//  - Val
//  - Type
type FeatureVal struct {
	Val  string         `thrift:"val,1,required" json:"val"`
	Type FeatureValType `thrift:"type,2,required" json:"type"`
}
