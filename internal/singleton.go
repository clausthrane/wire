package internal

import (
	"sync"
)

type containerItem struct {
	syncer   sync.Once
	builderF FactoryMethod
	res      interface{}
}


// A SingletonRepository implements wire.Repository
//
// When registering your factory methods with the singleton provider, they are only ever called once to create a single
// instances which is shared between callers.
//
// If you require more than one instance it is likely you want to register it under a separate name
type SingletonRepository struct {
	beans map[string]*containerItem
}

// NewSingletonRepository returns an instance of a SingletonRepository
func NewSingletonRepository() *SingletonRepository {
	r := &SingletonRepository{}
	r.beans = make(map[string]*containerItem)
	return r
}

// Register registers a factory under the given name
//
// Factory methods are invoked once and the result is cached for reuse
func (c *SingletonRepository) Register(name string, f FactoryMethod) {
	c.beans[name] = &containerItem{sync.Once{}, f, nil}
	return
}

// Get returns an instance as constructed by the factory registered under the given name
func (c *SingletonRepository) Get(name string) interface{} {
	if item, ok := c.beans[name]; ok {
		item.syncer.Do(func() {
			item.res = item.builderF()
		})
		return item.res
	}
	return nil
}
