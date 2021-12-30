package comm

//
//func StandardDispatcher(ctx context.Context, reqData ReqDataIf, pushEventHandler PushEventHandler) (*RespData, error) {
//	reqStr := reqData.GetPayload()
//	fmt.Println(reqStr)
//	//stdParams := structs.NewStandardParams()
//	//err := json.Unmarshal([]byte(reqStr), stdParams)
//	//if err != nil {
//	//	logutil.AddErrorLog(ctx, logutil.MDU_BIZ_FRAME, logutil.IDX_INVALID_REQ_FORMAT, err.Error())
//	//	return nil, err
//	//}
//
//	respData := &RespData{
//		TriggerId: 1,
//		Phone:     "1",
//		Pid:       1,
//	}
//	respData.Ctx = ctx
//
//	// 参数校验
//	//errNo := frame_comm.ValidateByTriggerId(ctx, stdParams.TriggerInfo.TriggerId, stdParams)
//	//if errNo != 0 {
//	//	respData.Resp = &structs.Response{
//	//		ErrNo:  int64(errNo),
//	//		ErrMsg: constant.GetErrMsg(errNo),
//	//	}
//	//	return respData, errors.New("param error")
//	//}
//
//	//基本过滤
//	//if stdParams.TriggerInfo.TriggerId == constant.DEFAULT_UNSET_INT || stdParams.ProductId == constant.DEFAULT_UNSET_INT ||
//	//	stdParams.EventInfo.CityId == constant.DEFAULT_UNSET_INT ||
//	//	stdParams.EventInfo.EventType == "" ||
//	//	stdParams.EventInfo.Source == "" ||
//	//	stdParams.InternationalInfo.UtcOffset == constant.DEFAULT_UNSET_INT ||
//	//	stdParams.InternationalInfo.Timestamp == constant.DEFAULT_UNSET_INT ||
//	//	(stdParams.EventInfo.Role == constant.ROLE_PASSENGER && stdParams.PassengerInfo.PassengerId == constant.DEFAULT_UNSET_INT && stdParams.PassengerInfo.Telphone == "") {
//	//	logutil.AddFilterLog(ctx, logutil.MDU_BIZ_FRAME, logutil.IDX_PARAMS_MISSING, "", 0, 0,
//	//		fmt.Sprint("pid=", stdParams.PassengerInfo.PassengerId),
//	//		fmt.Sprint("phone=", stdParams.PassengerInfo.Telphone),
//	//		fmt.Sprint("trigger_id=", stdParams.TriggerInfo.TriggerId),
//	//	)
//	//	return respData, nil
//	//}
//
//	//eventData := genEventDataFromStdParams(ctx, stdParams)
//	//respData.Pid = eventData.PassengerInfo.PassengerId
//	//respData.Phone = eventData.PassengerInfo.Telphone
//	//respData.TriggerId = eventData.TriggerInfo.TriggerId
//
//	//if !frame_comm.IsValidRequest(ctx, eventData) {
//	//	logutil.AddFilterLog(ctx, logutil.MDU_BIZ_FRAME, logutil.IDX_DEGRADE_FILTER, "", 0, 0)
//	//	return respData, nil
//	//}
//
//	//logutil.FileLog(&EventDataLog{Ctx: ctx, EventData: EventData})
//	//eventTime := eventData.EventInfo.EventTime.Format("2006-01-02 15:04:05")
//	//eventKey := "default"
//	//value, ok := eventData.ExtraInfo["event_info"].(map[string]interface{})
//	//if ok {
//	//	v, ok := value["key"].(string)
//	//	if ok {
//	//		eventKey = v
//	//	}
//	//}
//
//	//publicContext := logutil.NewPublicContext(1, eventData.PassengerInfo.PassengerId, eventData.PassengerInfo.Telphone, eventData.PassengerInfo.UserId,
//	//	eventData.OrderInfo.OrderId, eventData.TriggerInfo.TriggerId, eventTime, eventData.EventInfo.EventTime.Unix(), eventKey, eventData.EventInfo.CityId,
//	//	eventData.ProductInfo.EngProductId, 0)
//	//ctx = logutil.SetPublicContext(ctx, publicContext)
//
//	var eventData *_struct.EventData//MOCK
//	resp, err := pushEventHandler(ctx, eventData)
//	if err != nil {
//		//psg_core_log.AddNormalLog(ctx, *eventData, constant.LOG_EVENT_STATUS_PROCESS_FAILED)
//		return respData, err
//	}
//	respData.Resp = resp
//
//	return respData, nil
//}
