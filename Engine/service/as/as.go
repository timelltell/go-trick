package as

import (
	"GolangTrick/Engine/constant"
	basic_condition "GolangTrick/Engine/service/as/base"
	_struct "GolangTrick/Engine/struct"
	"github.com/codegangsta/inject"
	"golang.org/x/net/context"
	"strconv"
	"strings"
)

//开放condition服务， conditionKey -> stepId+ index  -> true/false
type OpenServiceResp map[string]map[string]bool

type BasicConditionIf interface {
	Judge(conditionInfo _struct.Condition, ioc inject.Invoker, eventData *_struct.EventData) (bool, error)
}

type ConditionFuncsRelManager struct {
	basicCondition map[int]BasicConditionIf
}

var CondFuncsRelationManager ConditionFuncsRelManager

type AdvancedSearch struct {
	relationManager ConditionFuncsRelManager
}

func init() {
	CondFuncsRelationManager.basicCondition = map[int]BasicConditionIf{
		constant.CONDITION_Drv_Points: basic_condition.ConditionDrvPoints{},
	}
}

//New AdvancedSearch
func NewAS() *AdvancedSearch {
	return &AdvancedSearch{relationManager: CondFuncsRelationManager}
}

//根据condition判断获取对应的action
func (as *AdvancedSearch) MatchStep(ctx context.Context, userOpStream _struct.UserOPStream, ioc inject.Invoker) (result []_struct.Step, err error) {
	//没有step的情况，直接返回成功
	if len(userOpStream.Steps) < 1 {
		return result, err
	}

	//遍历step list
	for _, stepInfo := range userOpStream.Steps {
		//处理一个step中的多个condition
		stepRes, stepErr := as.handleConditionsAnd(stepInfo, ioc, &userOpStream.EvData, ctx, userOpStream.Object.Id)
		if stepErr != nil {
			err = stepErr
			continue
		}

		if stepRes {
			result = append(result, stepInfo)
		}
	}

	return result, err
}

//每个condition之间按且的关系处理
func (as *AdvancedSearch) handleConditionsAnd(stepInfo _struct.Step, ioc inject.Invoker, evData *_struct.EventData, ctx context.Context, canvasId int) (bool, error) {
	result := false
	var err error

	//无condition
	if len(stepInfo.Conditions) == 0 {
		result = true
		return result, err
	}

	//有condition
	for index, condV := range stepInfo.Conditions {
		tmpRes := false
		var tmpErr error
		//logutil.AddInfoLog(ctx, logutil.MDU_OPEN_CONDITION, logutil.IDX_CONDITION_JUDGE, fmt.Sprintf("condKey：%v", condV.Key))
		if strings.HasPrefix(condV.Key, "pope.common") {
			resultMap := getResultMapFromIoc(ctx, ioc)
			tmpRes, tmpErr = judgeOpenCondition(condV, stepInfo.Id, index, resultMap)
		} else {
			tmpRes, tmpErr = as.handleSingleCondition(&condV, ioc, evData, ctx)
		}

		if tmpErr != nil {
			result = false
			err = tmpErr
			//logutil.AddErrorLog(ctx, logutil.MDU_ADVANCE_SEARCH, logutil.IDX_ADVANCE_SEARCH_CONDITION_FAILED, err.Error(), fmt.Sprintf("stepId=%v||conditionId=%v", stepInfo.Id, condV.Id))
			break
		}
		//condition, _ :=json.Marshal(condV)
		if !tmpRes {
			result = false
			//logutil.AddFilterLog(ctx, logutil.MDU_STEP_FILTER, logutil.IDX_CONDITION,
			//	string(condition),
			//	canvasId,
			//	stepInfo.Id,
			//	fmt.Sprintf("condition_info:%+v", condV),
			//	fmt.Sprint("pid=", evData.PassengerInfo.PassengerId),
			//	fmt.Sprint("phone=", evData.PassengerInfo.Telphone),
			//)
			break
		} else {
			result = true
			//logutil.AddInfoLog(ctx, logutil.MDU_STEP_FILTER, logutil.IDX_ADVANCE_SEARCH_COND_RESULT, "passed", fmt.Sprintf("stepId=%v||conditionId=%v", stepInfo.Id, condV.Id))
		}
	}
	return result, err
}

//判断单个condition
func (as *AdvancedSearch) handleSingleCondition(conditionInfo *_struct.Condition, ioc inject.Invoker, evData *_struct.EventData, ctx context.Context) (bool, error) {
	result := false
	conditionObj, _ := as.relationManager.basicCondition[conditionInfo.Id]

	//if err != nil {
	//	return result, err
	//}
	result, _ = conditionObj.Judge(*conditionInfo, ioc, evData)
	//return result, err
	return result, nil
}

func getResultMapFromIoc(ctx context.Context, ioc inject.Invoker) map[string]map[string]bool {
	resultMap := make(map[string]map[string]bool)
	_, invokeErr := ioc.Invoke(func(commonResp OpenServiceResp) {
		resultMap = commonResp
	})
	if invokeErr != nil {
		//logutil.AddErrorLog(ctx, logutil.MDU_OPEN_CONDITION, logutil.IDX_INVOKE_FAILED, fmt.Sprintf("err:%v", invokeErr))
	}
	return resultMap
}

//开放能力judge开放condition
func judgeOpenCondition(condv _struct.Condition, stepId int, index int, resultMap map[string]map[string]bool) (bool, error) {
	condRes, ok := resultMap[condv.Key]
	if !ok {
		return false, nil
	}
	res, ok := condRes[IntConvertStr(stepId, index)]
	if !ok {
		return false, nil
	}
	return res, nil
}

//stepId和index唯一标识一个condition
func IntConvertStr(stepId, index int) string {
	step := strconv.Itoa(stepId)
	ind := strconv.Itoa(index)
	return step + "_" + ind
}
