package lib

import _struct "GolangTrick/Engine/struct"

func InitConsumer() error {
	if !_struct.GConfig.MQconsumer.Subscribe {
		return nil
	}
	//消费程序
	//eventConsumer = consumer.NewCsdCarreraConsumer(&consumer.Config{
	//这个是消费函数
	//	MsgProceedFunc: comm.ProcessMqMessage,

	//	GoroutineNum:   _struct.GConfig.MQconsumer.EventGoroutineNum,
	//	Group:          _struct.GConfig.MQconsumer.EventGroup,
	//	CsdEnv:         common.GetMqEnv(),
	//	PullGoroutineNum: pullGoroutineNum,
	//})

	return nil
}
