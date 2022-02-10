package drv_points1

import (
	model1 "GolangTrick/PopeFS/model"
	"fmt"
	"git.xiaojukeji.com/falcon/pope-fs/util/errutil"
	"golang.org/x/net/context"
	"reflect"
	"strconv"
)

type driverRequest struct {
	UserID     int64  `tag:"user_id"`
	PointsType string `tag:"points_type"`
}

func GetDriverPoints(ctx context.Context, keys []string, params map[string]string) (model1.ThriftFeatureMap, []string, error) {
	res := make(model1.ThriftFeatureMap)
	//service, err := config.GetServiceMeta("drv_points")
	//if err != nil {
	//	return nil, nil, err
	//}
	//var queryParam driverRequest
	//err = datautil.ConvertMapToStruct(params, &queryParam)
	//if err != nil {
	//	return nil, nil, err
	//}
	//client := rpcsdk.NewDrvPointsClient(ctx, service.Timeout)
	//userIDStr := strconv.FormatInt(queryParam.UserID, 10)
	//data, err := client.GetDrvPointsByUserID(userIDStr)
	//if err != nil {
	//	return nil, nil, err
	//}
	//val := popefsthrift.NewFeatureVal()
	//val.Type = popefsthrift.FeatureValType_JSON
	//iData, _ := json.Marshal(data)
	//val.Val = string(iData)
	//res[keys[0]] = val
	return res, nil, nil
}

// ConvertMapToStruct convert map to struct
func ConvertMapToStruct(params map[string]string, val interface{}) error {
	structVal := reflect.ValueOf(val).Elem()
	structType := structVal.Type()
	for i := 0; i < structVal.NumField(); i++ {
		fieldType := structType.Field(i)
		field := structVal.FieldByName(fieldType.Name)
		name := fieldType.Tag.Get("tag")
		if name == "" {
			continue
		}
		ignore := fieldType.Tag.Get("ignore")
		def := fieldType.Tag.Get("default")
		if ignore == "true" {
			continue
		}
		data, ok := params[name]
		if !ok {
			if def == "" {
				return errutil.New(errutil.ErrMissParam, fmt.Sprintf("need params: %v", name))
			} else {
				data = def
			}
		}
		err := convertStringToVal(data, field)
		if err != nil {
			return err
		}
	}
	return nil
}

func convertValToString(val reflect.Value) (string, error) {
	kind := getKind(val)
	switch kind {
	case reflect.String:
		return val.String(), nil
	case reflect.Int:
		return fmt.Sprint(val.Int()), nil
	case reflect.Float64:
		return fmt.Sprint(val.Float()), nil
	default:
		return "", errutil.New(errutil.ErrInvalidParam, fmt.Sprintf("not support data type:%v", kind))
	}
}

func convertStringToVal(data string, val reflect.Value) error {
	kind := getKind(val)
	switch kind {
	case reflect.String:
		val.SetString(data)
		return nil
	case reflect.Int:
		num, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return errutil.New(errutil.ErrInvalidParam, fmt.Sprintf("parse int is err data:%v", data))
		}
		val.SetInt(num)
		return nil
	case reflect.Float64:
		num, err := strconv.ParseFloat(data, 64)
		if err != nil {
			return errutil.New(errutil.ErrInvalidParam, fmt.Sprintf("parse float is err data:%v", data))
		}
		val.SetFloat(num)
		return nil
	case reflect.Bool:
		b, err := strconv.ParseBool(data)
		if err != nil {
			return errutil.New(errutil.ErrInvalidParam, fmt.Sprintf("parse bool is err data:%v", data))
		}
		val.SetBool(b)
		return nil
	default:
		return errutil.New(errutil.ErrInvalidParam, fmt.Sprintf("no suppert data and type:%v data:%v", kind, data))
	}
}

func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float64
	default:
		return kind
	}
}
