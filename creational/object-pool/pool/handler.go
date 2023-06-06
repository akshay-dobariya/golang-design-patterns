package pool

import (
	"fmt"
	"sync"

	"github.com/akshay-dobariya/golang-design-patterns/creational/object-pool/poolobject"
)

type Pool struct {
	idleObjects   []poolobject.PoolObjectInterface
	activeObjects []poolobject.PoolObjectInterface
	capacity      int
	lock          sync.Mutex
}

func InitPool(poolObjects []poolobject.PoolObjectInterface) (*Pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("atleast one pool object is required")
	}
	acticeObjects := []poolobject.PoolObjectInterface{}
	return &Pool{
		idleObjects:   poolObjects,
		activeObjects: acticeObjects,
		capacity:      len(poolObjects),
		lock:          sync.Mutex{},
	}, nil
}

func (p *Pool) GetPoolObject() (*poolobject.PoolObjectInterface, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if len(p.idleObjects) == 0 {
		return nil, fmt.Errorf("no idle pool object")
	}
	poolObject := p.idleObjects[0]
	p.idleObjects = p.idleObjects[1:]
	p.activeObjects = append(p.activeObjects, poolObject)
	return &poolObject, nil
}

func (p *Pool) FreePoolObject(poolObject *poolobject.PoolObjectInterface) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	length := len(p.activeObjects)
	for i, obj := range p.activeObjects {
		if obj.GetId() == (*poolObject).GetId() {
			p.activeObjects[length-1], p.activeObjects[i] = p.activeObjects[i], p.activeObjects[length-1]
			p.activeObjects = p.activeObjects[:length-1]
			p.idleObjects = append(p.idleObjects, obj)
			return nil
		}
	}
	return fmt.Errorf("poolObject with id: '%d' is not part of the this pool", (*poolObject).GetId())
}
