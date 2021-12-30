package common

import (
	"GolangTrick/Engine/dao"
	_struct "GolangTrick/Engine/struct"
	"context"
)

type Objectfilter interface {
	Filter(context.Context, []_struct.Object) ([]_struct.Object, error)
}

type Bs struct {
	db      *dao.DB
	filter  []Objectfilter
	objects []_struct.Object
}
