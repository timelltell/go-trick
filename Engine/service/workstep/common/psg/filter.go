package psg

import (
	"GolangTrick/Engine/dao"
	_struct "GolangTrick/Engine/struct"
	"errors"
	"golang.org/x/net/context"
)

func NewAgeFilter(age int, db *dao.DB) *ageFilter {
	return &ageFilter{
		age: age,
		db:  db,
	}
}

type ageFilter struct {
	age int
	db  *dao.DB
}

func (tf *ageFilter) Filter(ctx context.Context, canvases []_struct.Object) ([]_struct.Object, error) {
	var err error = nil
	if tf.age > 0 {
		canvasesFromIndexer, err := tf.db.GetObjectByAge(tf.age)
		if err != nil {
			//logutil.AddFilterLog(ctx, logutil.MDU_BASIC_SEARCH, logutil.IDX_TRIGGER_FILTER, bs.ERR_NO_MATCH_CANVASES_BY_EVENT, 0, 0, fmt.Sprint("trigger_id=", tf.triggerId))
			return nil, errors.New("1")
		}
		for _, iv := range canvasesFromIndexer {
			canvases = append(canvases, iv)
		}
		if len(canvases) < 1 {
			//logutil.AddFilterLog(ctx, logutil.MDU_BASIC_SEARCH, logutil.IDX_TRIGGER_FILTER, bs.ERR_NO_MATCH_CANVASES_BY_EVENT, 0, 0)
			return nil, errors.New("1")
		}
	} else {
		//logutil.AddFilterLog(ctx, logutil.MDU_BASIC_SEARCH, logutil.IDX_TRIGGER_FILTER, "trigger id is zero below", 0, 0)
		return nil, errors.New("1")
	}
	return canvases, err
}
