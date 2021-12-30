package bizloader

import "context"

type psgBizTask struct {
	fetcher *psgFetcher
	ctx     context.Context
}

func (t *psgBizTask) sign() int {
	return 1
}

func (t *psgBizTask) context() context.Context {
	return t.ctx
}

func (t *psgBizTask) run() error {
	return t.fetcher.fetch(t.context())
}
