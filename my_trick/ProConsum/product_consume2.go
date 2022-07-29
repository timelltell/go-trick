package ProConsum

import (
	"fmt"
	"sync"
)


type concurrentConfig struct {
	ConcurrentNum   int
	ProductFunc     func([]interface{}) interface{}
	ProductFuncArgs [][]interface{}
	ConsumeFunc     func(interface{})
}

type Option func(*concurrentConfig)

func ConcurrentNum(p int) Option {
	return func(s *concurrentConfig) {
		s.ConcurrentNum = p
	}
}
func ProductFunc(p func([]interface{}) interface{}) Option {
	return func(s *concurrentConfig) {
		s.ProductFunc = p
	}
}
func ProductFuncArgs(p [][]interface{}) Option {
	return func(s *concurrentConfig) {
		s.ProductFuncArgs = p
	}
}
func ConsumeFunc(p func(interface{})) Option {
	return func(s *concurrentConfig) {
		s.ConsumeFunc = p
	}
}

func NewConcurrentConfig(options ...Option) *concurrentConfig {
	var retryConfig *concurrentConfig = &concurrentConfig{}
	for _, option := range options {
		option(retryConfig)
	}
	return retryConfig
}

type ConcurrentModeler interface {
	Run(c *concurrentConfig)
}

type ConcurrentPCModel struct {
}

type FixedConcurrentPCModel struct {
}

func (tmp *ConcurrentPCModel) Run(conf *concurrentConfig) {
	task := make(chan interface{})
	quit := make(chan int, 2)
	defer func() {
		close(quit)
		err := recover()
		if err != nil {
		}
	}()

	var wg1 sync.WaitGroup
	go func() {
		product := conf.ProductFunc
		for _, args := range conf.ProductFuncArgs {
			wg1.Add(1)
			go func(args []interface{}) {
				task <- product(args)
				wg1.Done()
			}(args)
		}
		wg1.Wait()
		close(task)
		quit <- 1
	}()

	go func() {
		var wg2 sync.WaitGroup
		for {
			select {
			case canvasId, ok := <-task:
				if !ok {
					wg2.Wait()
					quit <- 1
					return
				}
				wg2.Add(1)
				go func() {
					conf.ConsumeFunc(canvasId)
					wg2.Done()
				}()
			}
		}
	}()
	<-quit
	<-quit
}


func (tmp *FixedConcurrentPCModel) Run(conf *concurrentConfig) {
	task := make(chan interface{}, conf.ConcurrentNum)
	quit := make(chan int, 2)
	defer func() {
		close(quit)
		err := recover()
		if err != nil {
		}
	}()

	go func() {
		var wg1 sync.WaitGroup
		product := conf.ProductFunc
		for _, args := range conf.ProductFuncArgs {
			wg1.Add(1)
			go func(args []interface{}) {
				task <- product(args)
				wg1.Done()
			}(args)
		}
		wg1.Wait()
		close(task)
		quit <- 1
	}()

	go func() {
		var wg2 sync.WaitGroup
		for i := 0; i < conf.ConcurrentNum; i++ {
			wg2.Add(1)
			go func() {
				for {
					select {
					case canvasId, ok := <-task:
						if !ok {
							wg2.Done()
							return
						}
						conf.ConsumeFunc(canvasId)
					}
				}
			}()
		}
		wg2.Wait()
		quit <- 1
	}()
	<-quit
	<-quit
}

func test() {
	var m sync.Map
	productfunc := func(i []interface{}) interface{} {
		return i[0]
	}
	productfuncargs := make([][]interface{}, 0, 1)
	for i := 0; i < 1000; i++ {
		productfuncargs = append(productfuncargs, []interface{}{i})
	}

	//var mCanvasIdSummary sync.Map
	ConsumeFunc := func(i interface{}) {
		canvasId := i.(int)
		m.Store(i, []int{canvasId})
	}

	conf := &concurrentConfig{
		(10),
		(productfunc),
		(productfuncargs),
		(ConsumeFunc),
	}
	var c ConcurrentModeler= new(FixedConcurrentPCModel)
	c.Run(conf)
	lens1 := 0
	m.Range(func(key, value interface{}) bool {
		//fmt.Printf("s1 key : %v, s1 value : %+v\n", key, value)
		lens1++
		return true
	})
	if lens1 != 1000 {
		fmt.Println("lens1: ", lens1)
	}
}