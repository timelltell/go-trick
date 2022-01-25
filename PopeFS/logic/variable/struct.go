package variable1

import (
	driverAdapter1 "GolangTrick/PopeFS/logic/variable/driverAdapter"
	model1 "GolangTrick/PopeFS/model"
	"golang.org/x/net/context"
)

type IRequest interface {
	Init(context.Context)                   // init request param
	PreBuildParams(map[string]string) error // 参数预处理
	BuildRequestParams(*driverAdapter1.Driver, []string) (map[string]string, error)
}

type IResponse interface {
	ParseResponse(model1.ThriftFeatureMap, []string) error
	Logic(model1.ThriftFeatureMap) error
	GetValFromData() model1.ThriftFeatureMap
}

type Virtual struct {
	Request  IRequest
	Response IResponse
	Keys     []string
	Params   map[string]string
	Symbol   string
}
type Request struct {
	DriverID     string
	StepID       string
	CanvasID     string
	ParamsTarget map[string]string
}

type Response struct {
	ErrNo  int64
	ErrMsg string
	Data   model1.ThriftFeatureMap
}

type VirtualI interface {
	Init(ctx context.Context, keys []string, params map[string]string)

	// 主流程
	BeforeRead() error
	Read() (model1.ThriftFeatureMap, error)
	AfterRead(model1.ThriftFeatureMap) (model1.ThriftFeatureMap, error)

	// 函数式模式 入口
	Run(p VirtualI) (model1.ThriftFeatureMap, error)
}

func (o *Response) ParseResponse(featureMap model1.ThriftFeatureMap, varKeys []string) (errRet error) {

	//for _, varKey := range varKeys {
	//	featureName := driverAdapter.GetFeatureName(varKey)
	//	if _, ok := featureMap[featureName]; !ok {
	//		errRet = errors.Wrap(errRet, fmt.Sprintf("feature adpater fail! feature:%s", featureName))
	//		continue
	//	}
	//
	//	switch featureName {
	//	case driverAdapter.FEATURE_DRIVERFS, driverAdapter.FEATURE_REWARD_DIVE, driverAdapter.FEATURE_DRIVER_GRADE:
	//		// 解析特征数据
	//		if v,err := parseJsonData(featureMap[featureName].Val, varKey);err!=nil {
	//			errRet = errors.Wrap(errRet, fmt.Sprintf("feature parse fail! err:%v", err))
	//			continue
	//		}else{
	//			featureMap[varKey] = &v
	//		}
	//
	//	default:
	//		errRet = errors.Wrap(errRet, fmt.Sprintf("feature parse fail! feature:%s", featureName))
	//	}
	//
	//}
	//
	//// 删除原始数据
	//o.Data = make(model.ThriftFeatureMap)
	//for _, varKey := range varKeys {
	//	featureName := driverAdapter.GetFeatureName(varKey)
	//	if _, ok := featureMap[varKey]; ok {
	//		o.Data[varKey] = &popefsthrift.FeatureVal{
	//			Val:featureMap[varKey].Val,
	//			Type:featureMap[varKey].Type,
	//		}
	//	}
	//
	//	delete(featureMap, featureName)
	//}

	return
}

// todo  复合变量的计算应该统一从o.Data中读取
func (o *Response) Logic(featureMap model1.ThriftFeatureMap) (errRet error) {

	// step 复合变量
	// 计算奖励总额
	//if featureMap[driverAdapter.REWARD_MONEY_ORDER_STREAM]!=nil{
	//	f1 := o.Data[driverAdapter.REWARD_MONEY_ORDER_STREAM]
	//	f2 := o.Data[driverAdapter.REWARD_MONEY_ORDER]
	//	sum_orderstream,err := strconv.ParseFloat(f1.Val,64)
	//	sum_order,err := strconv.ParseFloat(f2.Val,64)
	//	if err == nil {
	//		if int(sum_orderstream + sum_order)%100 == 0 {
	//			featureMap[driverAdapter.REWARD_MONEY_ORDER_STREAM].Val = fmt.Sprintf("%.f",sum_orderstream + sum_order)
	//		}else {
	//			featureMap[driverAdapter.REWARD_MONEY_ORDER_STREAM].Val = fmt.Sprintf("%.2f",sum_orderstream + sum_order)
	//		}
	//
	//	}else{
	//		errRet = errors.Wrap(errRet, fmt.Sprintf("feature get logic fail! varkey:%s err:%v", driverAdapter.REWARD_MONEY_ORDER_STREAM,err))
	//	}
	//}
	//
	//// 计算奖励剩余天数
	//if featureMap[driverAdapter.REWARD_DATE_DAYLEFT]!=nil {
	//	f1 := o.Data[driverAdapter.REWARD_DATE_END]
	//	t,err := time.ParseInLocation("2006-01-02 15:04:05",f1.Val + " 23:59:59", time.Local)
	//	if err == nil {
	//		offset := math.Ceil(t.Sub(time.Now()).Seconds()/86400.0)
	//		if offset>0 {
	//			featureMap[driverAdapter.REWARD_DATE_DAYLEFT].Val = strconv.Itoa(int( offset))
	//		}else{
	//			featureMap[driverAdapter.REWARD_DATE_DAYLEFT].Val = ""
	//		}
	//
	//	}else{
	//		errRet = errors.Wrap(errRet, fmt.Sprintf("feature get logic fail! varkey:%s err:%v", driverAdapter.REWARD_DATE_DAYLEFT,err))
	//	}
	//}
	//// 计算奖励剩余单量
	//if featureMap[driverAdapter.REWARD_ORDER_LEFT]!=nil {
	//	f1 := o.Data[driverAdapter.REWARD_ORDER_TARGET]
	//	f2 := o.Data[driverAdapter.REWARD_ORDER_COUNT]
	//	order_target,err := strconv.Atoi(f1.Val)
	//	order_count,err := strconv.Atoi(f2.Val)
	//	order_left := 0
	//	if order_count < order_target{
	//		order_left = order_target - order_count
	//	}
	//	if err == nil{
	//		if order_left != 0 {
	//			featureMap[driverAdapter.REWARD_ORDER_LEFT].Val = strconv.Itoa(order_left)
	//		}else{
	//			featureMap[driverAdapter.REWARD_ORDER_LEFT].Val = ""
	//		}
	//	}else{
	//		errRet = errors.Wrap(errRet, fmt.Sprintf("feature get logic fail! varkey:%s err:%v", driverAdapter.REWARD_ORDER_LEFT,err))
	//	}
	//}
	//
	//
	//// step-1 业务逻辑
	//for varKey, FeatureVal := range featureMap {
	//
	//	desc, _ := driverAdapter.GetFeatureDesc(varKey)
	//	switch desc {
	//	case driverAdapter.PARSE_PATH_REGISTER_DOC:
	//		//证件审核逻辑
	//		registerLicLogic(FeatureVal, varKey)
	//	case driverAdapter.PARSE_PATH_REWARD_MONEY_ORDER,driverAdapter.PARSE_PATH_REWARD_MONEY_ORDER_STREAM:
	//		//滴分转换成小数点
	//		amount, _ := strconv.ParseFloat(FeatureVal.Val, 64)
	//		if int(amount)%100 != 0 {
	//			FeatureVal.Val = fmt.Sprintf("%.2f",amount / 100)
	//		}else {
	//			FeatureVal.Val = fmt.Sprintf("%.f",amount / 100)
	//		}
	//
	//	case driverAdapter.PARSE_PATH_REWARD_DATE_END:
	//		//去掉年份
	//		if varKey == driverAdapter.REWARD_DATE_END {
	//			timeformat := []byte(FeatureVal.Val)
	//			FeatureVal.Val = string(timeformat[5:])
	//		}
	//
	//	default:
	//		//errRet = errors.Wrap(errRet, fmt.Sprintf("feature get desc fail! varkey:%s", varKey))
	//	}
	//}

	return nil
}

func (o *Response) GetValFromData() model1.ThriftFeatureMap {
	return o.Data
}

func (o *Request) PreBuildParams(param map[string]string) (err error) {
	//if _,ok := param["context"];!ok {
	//	return nil
	//}
	//
	//if err := json.Unmarshal([]byte(param["context"]), &o); err != nil {
	//	return err
	//}
	//
	//o.DriverID = param["driver_id"]
	//o.CanvasID = param["canvas_id"]
	//o.StepID = param["step_id"]
	//
	//o.Extra.Data = make(map[string]string)
	//
	//if o.Extra.Temp != nil {
	//	for k, v := range o.Extra.Temp.(map[string]interface{}) {
	//		o.Extra.Data[k], err = utils.ConvertToString(v)
	//	}
	//}

	return
}

func (o *Request) Init(context.Context) {
	return
}

func (o *Request) BuildRequestParams(driver *driverAdapter1.Driver, keys []string) (paramsOut map[string]string, err error) {

	//paramsIn := map[string]string{
	//	"driver_id":  o.DriverID,
	//	"product_id": o.Extra.Data["product_id"],
	//}

	//if o.Event.CityId <= 0 {
	//	paramsIn["city_id"] = ""
	//}else{
	//	paramsIn["city_id"] = strconv.FormatInt(o.Event.CityId, 10)
	//}
	//
	//// 透传TCA外部业务参数 todo 不单独处理参数，规范TCA上下文
	//logicParams, err := datautil.ConvertStructToMap(&o.Trigger.Data)
	//if err == nil {
	//	for k, v := range logicParams {
	//		paramsIn[k] = v
	//	}
	//}
	//if code,ok := o.Extra.Data["canonical_country_code"]; ok {
	//	paramsIn["country_code"] = code
	//}
	//if status,ok := o.Extra.Data["grade_status"]; ok {
	//	paramsIn["grade_status"] = status
	//}
	//
	//paramsOut = driver.BuildParams(paramsIn, keys)

	return nil, nil
}
