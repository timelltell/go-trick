package comm

import (
	_struct "GolangTrick/Engine/struct"
	"encoding/json"
)

var DefaultgetRequestErrRespFn = func() string {
	respStr, _ := GenerateHttpResponse(10, "internal error", nil)
	return respStr
}

func GenerateHttpResponse(errNo int64, errMsg string, data interface{}) (string, error) {
	m := Response{}
	m.ErrNo = errNo

	if errNo != 0 {
		m.ErrMsg = errMsg
	} else {
		m.ErrMsg = "success"
	}

	m.Data = data

	respStr, _ := json.Marshal(m)

	return string(respStr), nil
}

var DefaultGererateRespFn = func(resp *_struct.RespData, err error) string {
	if err != nil {
		respStr, _ := GenerateHttpResponse(10, "internal error", nil)
		return respStr
	}

	if resp != nil && resp.Resp != nil {
		respStr, _ := json.Marshal(resp.Resp)
		return string(respStr)
	}

	respStr, _ := GenerateHttpResponse(0, "success", nil)
	return string(respStr)
}
