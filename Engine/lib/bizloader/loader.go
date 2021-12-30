package bizloader

import (
	"context"
	"errors"
)

type taskIf interface {
	sign() int
	context() context.Context
	run() error
}

type BizLoader struct {
	listeners []listenerIf
	tasks     map[int]taskIf
	queue     chan int
	initSigns []int
}

var (
	bizLoader *BizLoader
)

func GetLoader() *BizLoader {
	return bizLoader
}

func newBizLoader() *BizLoader {
	return &BizLoader{
		listeners: make([]listenerIf, 0, 3),
		tasks:     make(map[int]taskIf),
		queue:     make(chan int, 100),
		initSigns: make([]int, 0, 10),
	}
}

func (l *BizLoader) addListener(listener listenerIf) *BizLoader {
	l.listeners = append(l.listeners, listener)
	return l
}

func (l *BizLoader) trigger(s int) {
	l.queue <- s
}

func (l *BizLoader) addTask(task taskIf) error {
	_, ok := l.tasks[task.sign()]
	if ok {
		return errors.New("DUPLICATE")
	}
	l.tasks[task.sign()] = task
	return nil
}

func (l *BizLoader) initSign(sign []int) {
	l.initSigns = sign
}

func (l *BizLoader) forceLoad() error {
	var err error
	for _, sign := range l.initSigns {
		err := l.runTask(sign)
		if err != nil {
			return err
		}
	}
	return err
}

func (l *BizLoader) runTask(s int) error {
	task, err := l.getTask(s)
	if err != nil {
		return err
	} else {
		//ctx := context.Background()
		//ctx = logutil.SetLogScene(ctx, logutil.DL_LOG_SCENE)
		return task.run()
	}

}

func (l *BizLoader) getTask(s int) (taskIf, error) {
	task, ok := l.tasks[s]
	if ok {
		return task, nil
	} else {
		return nil, errors.New("err_task_sign_not_found")
	}
}
