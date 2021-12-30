package bizloader

import (
	_struct "GolangTrick/Engine/struct"
	"golang.org/x/net/context"
	"time"
)

func Init(config *_struct.Config) (*BizLoader, error) {
	bizLoader = newBizLoader()

	bizLoader.addListener(&psgBizListener{
		freq:    time.Second,
		trigger: bizLoader.trigger,
		config:  config,
	})

	bizLoader.addTask(&psgBizTask{
		fetcher: newPsgFetcher(config),
		ctx:     context.Background(),
	})

	bizLoader.initSign([]int{1})
	err := bizLoader.forceLoad()
	if err != nil {
		return nil, err
	} else {
		return bizLoader, nil
	}
}
