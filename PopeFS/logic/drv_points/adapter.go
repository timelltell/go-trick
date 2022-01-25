package drv_points1

import (
	model1 "GolangTrick/PopeFS/model"
	"golang.org/x/net/context"
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
