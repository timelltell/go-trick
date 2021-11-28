package tric

import (
	"errors"
	"fmt"
	"time"
)

type processFunc func([]interface{}) (interface{},error)

type RetryConfig struct {
	RetryCount      int
	ProductFunc     processFunc
	ProductFuncArgs []interface{}
	RetryInterval int
}


func Run(retryConfig *RetryConfig) (interface{},error) {
	var err error = nil
	defer func() {
		panicerr:=recover()
		if panicerr != nil {
			fmt.Println("panicerr:",panicerr)
			err=errors.New(panicerr.(string))
		}
	}()
	var count int = 0
	for range time.Tick(time.Second * time.Duration(retryConfig.RetryInterval)) {
		res,err:=retryConfig.ProductFunc(retryConfig.ProductFuncArgs)
		if err != nil {
			fmt.Println("err:",err)
		} else {
			return res,nil
		}
		count++
		if count < retryConfig.RetryCount {
			continue
		} else {
			return nil,errors.New("retry to many times")
		}
	}
	return nil,err
}


func NewCreateMsgDTO(options ...oprion) *RetryConfig {
	var createMsgDTO *RetryConfig = &RetryConfig{
		1,
		func(i []interface{}) (interface{},error){
			fmt.Println("i: ",i)
			return nil,nil
		},
		[]interface{}{1},
		1,
	}
	for _, option := range options {
		option(createMsgDTO)
	}
	return createMsgDTO
}

type oprion func(*RetryConfig)

func RetryCount(p int) oprion {
	return func(s *RetryConfig) {
		s.RetryCount = p
	}
}
func ProductFunc(p processFunc) oprion {
	return func(s *RetryConfig) {
		s.ProductFunc = p
	}
}
func ProductFuncArgs(p []interface{}) oprion {
	return func(s *RetryConfig) {
		s.ProductFuncArgs = p
	}
}
func RetryInterval(p int) oprion {
	return func(s *RetryConfig) {
		s.RetryInterval = p
	}
}