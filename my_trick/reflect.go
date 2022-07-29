package my_trick

import (
	"fmt"
	"reflect"
	"strings"
)

func AllTypes() {
	for i := reflect.Invalid; i <= reflect.UnsafePointer; i++ {
		fmt.Println("i: ", i.String())
	}
}

type FilterLogModel struct {
	//请求id
	RequestId string `json:"request_id"`
	//时间戳
	Ts int64 `json:"ts"`
	//业务线id
	Bid string `json:"bid"`
	//yy-MM-dd 分区用
	EventTime string `json:"event_time"`
	//国家码country_code
	CC string `json:"c_c"`
	//
	Uid int64 `json:"uid"`
	//过滤原因标志
	Index string `json:"index"`
	//traceid
	Tid string `json:"t_id"`
	//span_id
	Sid string `json:"s_id"`
	//资源位ID,resource_id
	Rid string `json:"r_id"`
	//计划ID
	PlanId string `json:"plan_id"`
	//applo分组key
	Aid string `json:"a_id"`
	//物料id
	Mid string `json:"m_id"`
	//手机号
	Tel string `json:"tel"`
	//司乘角色
	Role string `json:"role"`
	//每种过滤的特殊字段
	Msg string `json:"msg"`
	//结果
	Res string `json:"res"`
}

func GetFormatStr() (res string) {
	t := reflect.TypeOf(new(FilterLogModel)).Elem()
	strList := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		strList = append(strList, strings.Join([]string{t.Field(i).Tag.Get("json"), "=v%"}, ""))
	}
	res = strings.Join(strList, "||")
	return res
}
