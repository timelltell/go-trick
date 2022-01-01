package dao

import (
	_struct "GolangTrick/Engine/struct"
	"errors"
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

func (d *DB) GetObjects() ([]_struct.Object, error) {
	objectsFromMemory, err := d.ObjectManager.getObjects()

	if err != nil {
		return nil, err
	}
	realObjects := make([]_struct.Object, 0, 100)
	for _, objectFromMemory := range objectsFromMemory {
		realObjects = append(realObjects, *objectFromMemory)
	}
	return realObjects, nil
}

func (d *DB) GetObjectAgeById(id int) (*_struct.Object, error) {
	objectsFromMemory, err := d.ObjectManager.getObjectById(id)
	if err != nil {
		return nil, err
	}
	return objectsFromMemory, nil
}

func (d *DB) GetObjectByAge(age int) ([]_struct.Object, error) {
	objectsFromMemory, err := d.ObjectManager.getObjects()

	if err != nil {
		return nil, err
	}
	realObjects := make([]_struct.Object, 0, 100)
	for _, objectFromMemory := range objectsFromMemory {
		if objectFromMemory.Age == age {
			realObjects = append(realObjects, *objectFromMemory)
		}
	}
	return realObjects, nil
}

func (d *DB) Fill(apiRepData *_struct.PopeIndexerResponse) error {
	if apiRepData.Errcode != 0 {
		return errors.New("err")
	}
	canvasManager := NewObjectManager()
	err := canvasManager.load(apiRepData.Data)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	d.ObjectManager = canvasManager
	return nil
}
