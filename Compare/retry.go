package Compare

import (
	"fmt"
	"time"
)

type processFunc func(interface{}) (interface{},error)

type RetryConfig struct {
	RetryCount   int
	ProductFunc     processFunc
	ProductFuncArgs []interface{}
	RetryInterval int
}


func Run(retryConfig *RetryConfig) (interface{},error) {
	var count int = retryConfig.RetryCount
	for range time.Tick(time.Second * time.Duration(retryConfig.RetryInterval)) {
		res,err:=retryConfig.ProductFunc(retryConfig.ProductFuncArgs)
		if err != nil {
			fmt.Println("err:",err)
		} else {
			return res,nil
		}
		count++
		if count < 5 {
			continue
		} else {
			return nil,errors.New("retry to many times")
		}
	}
}


func NewCreateMsgDTO(options ...Option) *RetryConfig {
	var createMsgDTO *RetryConfig = &RetryConfig{
		1,
		func(i interface{}) (interface{},error){
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

type Option func(*RetryConfig)

func RetryCount(p int) Option {
	return func(s *RetryConfig) {
		s.RetryCount = p
	}
}
func ProductFunc(p processFunc) Option {
	return func(s *RetryConfig) {
		s.ProductFunc = p
	}
}
func ProductFuncArgs(p []interface{}) Option {
	return func(s *RetryConfig) {
		s.ProductFuncArgs = p
	}
}
func RetryInterval(p int) Option {
	return func(s *RetryConfig) {
		s.RetryInterval = p
	}
}