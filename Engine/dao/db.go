package dao

import (
	_struct "GolangTrick/Engine/struct"
	"time"
)

var db *DB

func InitDB() error {
	db = &DB{}
	return nil
}

func GetInstance() *DB {
	return db
}

type DB struct {
	originData       *_struct.Config
	ObjectManager    *objectManager
	exploreObjectIDs []int
	timeTick         <-chan time.Time
	config           *_struct.Config
}

func (d *DB) OriginData() *_struct.Config {
	return d.originData
}

func (d *DB) GetObjects() ([]_struct.Objects, error) {
	objectsFromMemory, err := d.ObjectManager.getObjects()

	if err != nil {
		return nil, err
	}
	realObjects := make([]_struct.Objects, 0, 100)
	for _, objectFromMemory := range objectsFromMemory {
		realObjects = append(realObjects, *objectFromMemory)
	}
	return realObjects, nil
}
