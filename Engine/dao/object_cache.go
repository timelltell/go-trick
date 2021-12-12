package dao

import _struct "GolangTrick/Engine/struct"

func NewObjectCache() *objectCache {
	return &objectCache{
		objectsMap: make(map[int]*_struct.Objects),
	}
}

type objectCache struct {
	objectsMap  map[int]*_struct.Objects
	objectCount int
}

func (c *objectCache) getObjectById(Id int) *_struct.Objects {

	tmp, ok := c.objectsMap[Id]
	if ok {
		return tmp
	} else {
		return nil
	}
}

func (c *objectCache) setObjectById(object *_struct.Objects) {
	if object != nil {
		c.objectsMap[object.Id] = object
	}
}
