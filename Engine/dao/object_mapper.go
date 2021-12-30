package dao

import (
	_struct "GolangTrick/Engine/struct"
	"errors"
)

func NewObjectManager() *objectManager {
	return &objectManager{}
}

type objectManager struct {
	objectCache *objectCache
}

func (this *objectManager) load(objects []*_struct.Object) error {
	this.objectCache = NewObjectCache()
	var object *_struct.Object
	totalCount := len(objects)
	if totalCount > 0 {
		for index := 0; index < totalCount; index++ {
			object = objects[index]
			if object == nil {
				continue
			}
			this.objectCache.setObjectById(object)
		}
	}
	this.objectCache.objectCount = totalCount
	return nil
}

func (this *objectManager) getObjects() (map[int]*_struct.Object, error) {
	if this.objectCache.objectsMap == nil || len(this.objectCache.objectsMap) < 1 {
		return nil, errors.New("no object")
	}
	return this.objectCache.objectsMap, nil
}

func (this *objectManager) getObjectById(Id int) (*_struct.Object, error) {
	object := this.objectCache.getObjectById(Id)
	if object == nil {
		return nil, errors.New("no object")
	}
	return object, nil
}
