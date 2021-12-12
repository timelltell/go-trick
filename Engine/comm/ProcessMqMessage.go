package comm

import (
	"GolangTrick/Engine/service/workflow"
	_struct "GolangTrick/Engine/struct"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	//"runtime/debug"
	"time"
)

func MsgProceedFn(ctx context.Context, msg *_struct.MqMessage) bool {
	//startTime := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), _struct.GConfig.Proxy.TimeOut.Duration)
	defer func() {
		if err := recover(); err != nil {
			//logutil.AddErrorLog(ctx, logutil.ENGINE, logutil.IDX_MQ_CONSUMER_FAILED, "consume mq panic", string(debug.Stack()), err)
		}
	}()

	defer cancel()
	go func() {
		select {
		case <-time.After(_struct.GConfig.Proxy.TimeOut.Duration):
			//logutil.AddErrorLog(ctx, logutil.ENGINE, logutil.IDX_MQ_CONSUMER_TIMEOUT, "")
		case <-ctx.Done():
			return
		}

	}()
	reqData := &_struct.ReqData{}
	reqData.SetPayload(string(msg.Value))

	//if ctx.Topic ==  _struct.GConfig.MQconsumer.EventTopic {
	//	logutil.AddRequestInLog(ctx, logutil.ENGINE, logutil.IDX_MQ_CONSUMER_IN, "", "mqconsumer_event", string(msg.Value), nil, mqData)
	//} else {
	//	logutil.AddRequestInLog(ctx, logutil.ENGINE, logutil.IDX_MQ_CONSUMER_IN, "other topic", ctx.Topic, string(msg.Value), nil, mqData)
	//}

	resp, err := StandardDispatcher(ctx, reqData, PushEvent)
	if err != nil {
		return false
	}
	respStr := []byte{}
	if resp == nil {
		resp = &_struct.RespData{
			Ctx: ctx,
		}
	}
	if resp.Resp != nil {
		respStr, _ = json.Marshal(resp.Resp)
	}
	fmt.Println(respStr)

	return true
}

func StandardDispatcher(ctx context.Context, reqData *_struct.ReqData, pushEventHandler PushEventHandler) (*_struct.RespData, error) {
	reqStr := reqData.GetPayload()
	stdParams := _struct.NewStandardParams()
	err := json.Unmarshal([]byte(reqStr), stdParams)

	if err != nil {
		//logutil.AddErrorLog(ctx, logutil.ENGINE, logutil.IDX_INVALID_REQ_FORMAT, err.Error())
		return nil, err
	}

	respData := &_struct.RespData{}
	respData.Ctx = ctx
	eventData := _struct.GetEventDataFromStdParams(ctx, stdParams)

	resp, err := pushEventHandler(ctx, eventData)
	if err != nil {
		return respData, err
	}
	respData.Resp = resp

	return respData, nil
}

type PushEventHandler func(context.Context, *_struct.EventData) (*_struct.Response, error)

func PushEvent(ctx context.Context, eventData *_struct.EventData) (r *_struct.Response, err error) {
	result, err := pushEventHandler(ctx, eventData)
	if nil == result {
		return nil, err
	}
	return result.(*_struct.Response), err
}

func pushEventHandler(ctx context.Context, eventData *_struct.EventData) (interface{}, error) {
	es := GetEngineServer()
	var err error
	var gotoStep workflow.GoNext
	workStepIn := workflow.WorkStepData{
		EventData: *eventData,
	}

	workStepOut := workflow.WorkStepData{}
	workStep := es.workFlow.HeadStep(eventData.Scene)
	if workStep == nil {
		return nil, nil
	}

	finishFlag := false
	for !finishFlag {
		select {
		case <-ctx.Done():
			return nil, errors.New(fmt.Sprintf("context timeout %s", ctx.Err()))
		default:
		}

		workStepOut, err, gotoStep = workStep.Handle(ctx, workStepIn)
		if err != nil {
			errStep := es.workFlow.GotoErr(eventData.Scene, workStep.Key())
			if errStep != nil {
				_, errStepErr, _ := errStep.Handle(ctx, workStepOut)
				if errStepErr != nil {
					//logutil.AddErrorLog(ctx, logutil.ENGINE, logutil.IDX_PUSH_EVENT_STEP_FAILED, "", fmt.Sprintf("workStep=%+v||stepErr=%+v", errStep.Key(), errStepErr))
				} else {
					//logutil.AddInfoLog(ctx, logutil.ENGINE, logutil.IDX_PUSH_EVENT_INFO, "", fmt.Sprintf("workStep=%+v", errStep.Key()))
				}
			}
			if _, ok := err.(workflow.BizErr); ok {
				logutil.AddInfoLog(ctx, logutil.ENGINE, logutil.IDX_BIZ_INFO, err.Error())
				return nil, nil
			}
			return nil, err
		}

		workStepIn = workStepOut
		if gotoStep == workflow.GOTO_NEXT {
			workStep = es.workFlow.Next(eventData.Scene, workStep.Key())
		} else {
			workStep = es.workFlow.Goto(eventData.Scene, gotoStep)
		}
		if workStep == nil {
			finishFlag = true
		}
	}

	return workStepOut.ExtraData, err
}
