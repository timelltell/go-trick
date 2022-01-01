package common

import (
	"GolangTrick/Engine/dao"
	_struct "GolangTrick/Engine/struct"
	"context"
	"errors"
)

type Objectfilter interface {
	Filter(context.Context, []_struct.Object) ([]_struct.Object, error)
}

type Bs struct {
	db      *dao.DB
	filters []Objectfilter
	objects []_struct.Object
}

func NewBS() *Bs {
	return &Bs{
		filters: make([]Objectfilter, 0, 8),
	}
}

func (b *Bs) AddFilter(filter Objectfilter) {
	b.filters = append(b.filters, filter)
}
func (b *Bs) Db() *dao.DB {
	return b.db
}

func (b *Bs) SetSourceObjects(ctx context.Context, objects []_struct.Object) *Bs {
	if len(objects) < 1 {
		objects = make([]_struct.Object, 0, 32)
	}
	b.objects = objects

	objectsIds := make([]int, 0, len(objects))
	for _, c := range b.objects {
		objectsIds = append(objectsIds, c.Id)
	}
	return b
}

//get all match Objects by filter
func (b *Bs) MatchedObjects(ctx context.Context) ([]_struct.Object, error) {
	var err error
	for _, filter := range b.filters {

		b.objects, err = filter.Filter(ctx, b.objects)
		if err != nil {
			return nil, err
		}
		if len(b.objects) == 0 {
			break
		}
	}
	if len(b.objects) < 1 {
		return nil, errors.New("ERR_NO_CANVASES")
	}
	return b.objects, nil
}
