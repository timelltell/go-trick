package comm

import _struct "GolangTrick/Engine/struct"

var engineServer *EngineServer

type EngineServerFunc func() error
type EngineServer struct {
	initFuncs []EngineServerFunc
	Conf *_struct.Config
	signal chan int
	workFlow
}
