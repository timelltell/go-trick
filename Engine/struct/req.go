package _struct

import (
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

func GetReqFromBody(req *http.Request) (ReqDataer, error) {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	} else {
		reqData := &ReqData{
			Payload: bytes.NewBuffer(result).String(),
		}
		return reqData, nil
		//return bytes.NewBuffer(result).String(), nil
	}
}
