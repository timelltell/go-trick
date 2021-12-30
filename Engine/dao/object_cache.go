package dao

import _struct "GolangTrick/Engine/struct"

func NewObjectCache() *objectCache {
	return &objectCache{
		objectsMap: make(map[int]*_struct.Object),
	}
}

type objectCache struct {
	objectsMap  map[int]*_struct.Object
	objectCount int
}

func (c *objectCache) getObjectById(Id int) *_struct.Object {

	tmp, ok := c.objectsMap[Id]
	if ok {
		return tmp
	} else {
		return nil
	}
}

func (c *objectCache) setObjectById(object *_struct.Object) {
	if object != nil {
		c.objectsMap[object.Id] = object
	}
}
