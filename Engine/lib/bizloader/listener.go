package bizloader

import (
	_struct "GolangTrick/Engine/struct"
	"time"
)

type listenerIf interface {
	listen()
	handle()
	setTrigger(func(int))
}

type psgBizListener struct {
	freq    time.Duration
	trigger func(s int)
	config  *_struct.Config
}

func (l *psgBizListener) setTrigger(t func(int)) {
	l.trigger = t
}

func (l *psgBizListener) listen() {
	tick := time.Tick(l.freq)
	for {
		select {
		case <-tick:
			if time.Now().Unix()%int64(l.config.PopeEditor.IndexerReloadIntervalSeconds) == 0 {
				l.handle()
			}
		}
	}
}

func (l *psgBizListener) handle() {
	if l.trigger == nil {

	} else {
		l.trigger(1)
	}
}
