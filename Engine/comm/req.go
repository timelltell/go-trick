package comm

import (
	_struct "GolangTrick/Engine/struct"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ReqData struct {
	Payload   string
	Queryform url.Values
	Postform  url.Values
}

type ReqDataIf interface {
	SetPayload(string)
	GetPayload() string
	GetQueryForm() url.Values
	SetQueryForm(url.Values)
	GetPostForm() url.Values
	SetPostForm(url.Values)
}

func GetReqFromBody(req *http.Request) (*_struct.ReqData, error) {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	} else {
		reqData := &_struct.ReqData{
			Payload: bytes.NewBuffer(result).String(),
		}
		return reqData, nil
		//return bytes.NewBuffer(result).String(), nil
	}
}

func (rd *ReqData) SetPayload(payload string) {
	rd.Payload = payload
}

func (rd *ReqData) GetPayload() string {
	return rd.Payload
}

func (rd *ReqData) GetQueryForm() url.Values {
	return rd.Queryform
}

func (rd *ReqData) SetQueryForm(v url.Values) {
	rd.Queryform = v
}

func (rd *ReqData) GetPostForm() url.Values {
	return rd.Postform
}

func (rd *ReqData) SetPostForm(v url.Values) {
	rd.Postform = v
}
