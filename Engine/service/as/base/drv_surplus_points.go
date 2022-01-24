package basic_condition

import (
	_struct "GolangTrick/Engine/struct"
	"errors"
	"github.com/codegangsta/inject"
)

//司机积分
type DrvPoints int64

type ConditionDrvPoints struct{}

func (this ConditionDrvPoints) Judge(conditionInfo _struct.Condition, ioc inject.Invoker, eventData *_struct.EventData) (bool, error) {
	result := false

	if conditionInfo.Value == nil {
		return result, errors.New("")
	}

	//target类型转换
	realTargetVal, ok := conditionInfo.Value.(_struct.TargetConditionDrvPoints)
	if !ok {
		return result, errors.New("")
	}

	_, invokeErr := ioc.Invoke(func(drvPoints DrvPoints) {
		result = compareInt(int(drvPoints), int(realTargetVal.DrvPoints), conditionInfo.Expr)
	})

	//去获取drvPoints时出错
	if invokeErr != nil {
		result = false
		return result, invokeErr
	}

	return result, nil
}
