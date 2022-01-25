package controller1

import (
	logic1 "GolangTrick/PopeFS/logic"
	"encoding/json"
	"fmt"
	"git.xiaojukeji.com/falcon/pope-fs/logic/monitor"
	"git.xiaojukeji.com/falcon/pope-fs/model"
	"git.xiaojukeji.com/falcon/pope-fs/util/datautil"
	"git.xiaojukeji.com/falcon/pope-fs/util/errutil"
	"git.xiaojukeji.com/falcon/pope-fs/util/requtil"
	"git.xiaojukeji.com/gobiz/logger"
	"golang.org/x/net/context"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

// BaseController 基础 controller
type BaseController struct {
	name          string
	isRawResponse bool
	//HandleFunc 格式必要为func (ctx context.Context, req interface{}) (interface{}, error)
	handleFunc interface{}
}

// Request 存储各种类型的 request
type Request interface{}

// Response 存储各种类型的 response
type Response interface{}

//ResponseData 相应数据结构体
type ResponseData struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	LogID  string      `json:"logid,omitempty"`
	Data   interface{} `json:"data"`
}

//ContextKey 存储键的类型
type ContextKey string

//ResponseJSON 处理控制层返回的结构体 为 json字符串
func ResponseJSON(ctx context.Context, w http.ResponseWriter, respStruct interface{}, err error) string {
	respData := &ResponseData{
		ErrNo:  errutil.Success,
		ErrMsg: errutil.GetMessage(errutil.Success),
		LogID:  datautil.TransToString(requtil.GetLogRecord(ctx).LogID),
		Data:   respStruct,
	}
	return ResponseRawJSON(ctx, w, respData, err)
}

func resolveErr(err error) (int, string) {
	var errno int
	var errmsg string

	if detailErr, ok := err.(*errutil.ErrorInfo); ok {
		errno = detailErr.ErrNo
		errmsg = detailErr.ErrMsg
	} else {
		errno = errutil.ErrUnknown // default errno
		errmsg = errutil.GetMessage(errutil.ErrUnknown, err)
	}
	return errno, errmsg
}

//ResponseRawJSON 处理控制层返回的结构体 为 json字符串，不处理返回结构体的data包装
func ResponseRawJSON(ctx context.Context, w http.ResponseWriter, respData Response, err error) string {
	oldResp, assertOk := respData.(*ResponseData)
	// 判断是否发生错误
	if err != nil {
		errNo, errMsg := resolveErr(err)
		// 重新封装新的返回对象
		if assertOk {
			respData = &ResponseData{
				ErrNo:  errNo,
				ErrMsg: errMsg,
				LogID:  datautil.TransToString(requtil.GetLogRecord(ctx).LogID),
				Data:   oldResp.Data,
			}
		} else {
			respData = &ResponseData{
				ErrNo:  errNo,
				ErrMsg: errMsg,
				LogID:  datautil.TransToString(requtil.GetLogRecord(ctx).LogID),
			}
		}
	}
	jsonStr, err := json.Marshal(respData)
	if err != nil {
		var newResponseData *ResponseData
		if assertOk {
			newResponseData = &ResponseData{
				ErrNo:  errutil.ErrJSONMarshalFailed,
				ErrMsg: errutil.GetMessage(errutil.ErrJSONMarshalFailed),
				LogID:  datautil.TransToString(requtil.GetLogRecord(ctx).LogID),
				Data:   oldResp.Data,
			}
		} else {
			newResponseData = &ResponseData{
				ErrNo:  errutil.ErrJSONMarshalFailed,
				ErrMsg: errutil.GetMessage(errutil.ErrJSONMarshalFailed),
				LogID:  datautil.TransToString(requtil.GetLogRecord(ctx).LogID),
			}
		}

		jsonStr = []byte(formToJSON(newResponseData))
		if err != nil {
			logger.Errorf(ctx, logger.DLTagUndefined, "write data to response writer error")
		}
	}

	w.Header().Set("Content-Type", "application/json")

	// 将返回数据发送到客户端
	_, err = w.Write(jsonStr)
	if err != nil {
		logger.Errorf(ctx, logger.DLTagUndefined, "write data to response writer error")
		// 没有数据返回成功
		return ""
	}
	return string(jsonStr)
}

func formToJSON(r *ResponseData) string {
	return fmt.Sprintf("{\"errno\":%v,\"errmsg\":\"%v\",\"logid\":\"%v\",\"data\": \"\"}", r.ErrNo, r.ErrMsg, r.LogID)
}

// HandleHTTP ： BaseController 的 HandleHTTP 方法
func (bc *BaseController) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var responseStr string

	start := time.Now()
	defer func() {
		end := time.Now()
		lantencyNs := end.Sub(start)
		procTimeMs := datautil.SinceInMs(start)
		errCode := errutil.Success
		if err != nil {
			if errInfo, ok := err.(*errutil.ErrorInfo); ok {
				errCode = errInfo.ErrNo
			} else {
				errCode = errutil.ErrUnknown
			}
		}
		//捕捉panic
		if err := recover(); err != nil {
			//返回未知错误，以免客户端收不到信息
			responseStr = ResponseJSON(r.Context(), w, nil, errutil.New(errutil.ErrUnknown, errutil.GetMessage(errutil.ErrPanic)))
			errCode = errutil.ErrPanic
			logger.Errorf(r.Context(), logger.DLTagHTTPFailed, "Panic Error||errno=%d||err=%v||lantency=%d||proc_time=%.2f||stack=%s",
				errCode, err, lantencyNs, procTimeMs, string(debug.Stack()))
		} else if errCode != errutil.Success {
			//打印错误统计信息
			logger.Errorf(r.Context(), logger.DLTagRequestOut, "request=%v||response=%v||errno=%v||err=%v||lantency=%d||proc_time=%.2f",
				"request", responseStr, errCode, err, lantencyNs, procTimeMs/1000)
		}
		logger.Infof(r.Context(), logger.DLTagRequestOut, "request=%v||response=%v||errno=%v||lantency=%d||proc_time=%.2f",
			"request", responseStr, errCode, lantencyNs, procTimeMs/1000)
		//发送监控数据，不关心返回
		_ = monitor.GetStatistic().Response(strings.ToLower(bc.name), lantencyNs, errCode)
	}()

	//绑定入参
	var request Request
	var response Response
	if request, _, err = bc.bindRequest(r); err == nil {
		handleFunc := reflect.ValueOf(bc.handleFunc)
		//业务处理
		returnValues := handleFunc.Call([]reflect.Value{reflect.ValueOf(r.Context()), reflect.ValueOf(request)})
		//返回参数处理
		response = returnValues[0].Interface()
		if returnValues[1].Interface() != nil {
			err = returnValues[1].Interface().(error)
		}
	}

	//返回json
	if bc.isRawResponse {
		responseStr = ResponseRawJSON(r.Context(), w, response, err)
	} else {
		responseStr = ResponseJSON(r.Context(), w, response, err)
	}
}

// HTTPHandler 将 HandlerFunc 和 name 封装为新的 http.HandlerFunc
func HTTPHandler(name string, handleFunc interface{}) http.HandlerFunc {
	// 检查传入的参数
	//checkHandleFunc(handleFunc)
	// 构建返回对象
	baseController := &BaseController{name: name, isRawResponse: false, handleFunc: handleFunc}
	return baseController.HandleHTTP
}

// MGetHandler 批量获取特征入口
func MGetHandler(ctx context.Context, req *model.MGetReq) (*model.MGetResp, error) {

	convertedParams := convertMSIToMSS(req.Params)
	//logic.PrepareParam(ctx, convertedParams)
	res, err := logic1.Dispatch(ctx, req.Features, convertedParams)
	return &model.MGetResp{
		FeatureMap:  res.FeatureMap,
		FailedList:  res.FailedList,
		NoDataList:  res.NoDataList,
		InvalidList: res.InvalidList,
	}, err
}

// map[string]interface{} => map[string]string
func convertMSIToMSS(input map[string]interface{}) map[string]string {
	var output = make(map[string]string, len(input))
	for k, v := range input {
		output[k] = fmt.Sprintf("%v", v)
	}
	return output
}

// GetHandler 获取特征入口
func GetHandler(ctx context.Context, req *model.GetReq) (*model.GetResp, error) {
	return nil, nil
}

func (bc *BaseController) bindRequest(r *http.Request) (request Request, requestStr string, err error) {
	request = reflect.New(reflect.TypeOf(bc.handleFunc).In(1).Elem()).Interface()
	if err = requtil.Bind(r, request); err == nil {
		formJSONByte, _ := json.Marshal(r.Form)
		beanJSONByte, _ := json.Marshal(request)
		requestStr = string(beanJSONByte)
		logger.Infof(r.Context(), logger.DLTagRequestIn, "original=%s||request=%s", string(formJSONByte), requestStr)
	} else {
		formJSONByte, _ := json.Marshal(r.Form)
		logger.Infof(r.Context(), logger.DLTagRequestIn, "Request Data bind failed||original=%s||err=%s", string(formJSONByte), err)
	}
	return
}
