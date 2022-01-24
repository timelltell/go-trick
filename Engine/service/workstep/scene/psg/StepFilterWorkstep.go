package psg

import (
	"GolangTrick/Engine/lib/datagrip"
	"GolangTrick/Engine/service/as"
	"GolangTrick/Engine/service/workflow"
	_struct "GolangTrick/Engine/struct"
	"errors"
	"golang.org/x/net/context"
)

type StepFilterWorkstep struct {
}

func (cf StepFilterWorkstep) Key() string {
	return "FLOW_DIVERSION_STEP"
}

func (cf StepFilterWorkstep) Handle(ctx context.Context, in workflow.WorkStepData) (workflow.WorkStepData, error, string) {
	out := in

	opStreams := (in.OpStreams).([]_struct.UserOPStream)

	//Step condition filter
	outOpStreams := make([]_struct.UserOPStream, 0, len(opStreams))
	//1 先取出所有的condition的key，去data_installer获取provider list
	params := datagrip.QueryParam{EventData: &in.EventData, UosList: opStreams}
	dataProvider := datagrip.GetStepProvider(params)
	//2 调用fetch方法生成ioc
	ioc, err := datagrip.FetchData(ctx, &in.EventData, dataProvider...)
	if err != nil {
		return out, err, workflow.GOTO_NEXT
	}
	//3 遍历，逐个调用as.MatchStep返回step结果
	for _, opStream := range opStreams {
		tmpOpStream := opStream

		steps, err := as.NewAS().MatchStep(ctx, tmpOpStream, ioc)
		if nil != err {
			//logutil.AddErrorLog(ctx, logutil.MDU_STEP_FILTER, logutil.IDX_STEP_FILTER_FAILED, err.Error(), fmt.Sprintf("canvas_id=%d||userOPStream=%+v", tmpOpStream.Canvas.Id, tmpOpStream))
			continue
		}

		//只返回符合条件的step
		if len(steps) > 0 {
			tmpOpStream.Steps = steps
			tmpOpStream.Step = steps[0]
			outOpStreams = append(outOpStreams, tmpOpStream)
			//out.Canvases = append(outCanvases, tmpOpStream.Canvas)
		}
	}
	out.OpStreams = outOpStreams

	if 0 == len(outOpStreams) {
		return out, errors.New("no matched userOpStream"), workflow.GOTO_NEXT
	}
	return out, nil, workflow.GOTO_NEXT
}
