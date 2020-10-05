package wire

import "github.com/clausthrane/wire/internal"

// FactoryMethod is a method type used to provide instances to the ServiceRepository
type FactoryMethod = internal.FactoryMethod

// Repository provides access to instances factories registered during bootstrapping
//
// Current implementations are a singleton repository constructed via wire.SingletonRepository
//
// Register stores a factory method for the provided name. The factory method is invoked accoring to the semantics of
// the chosen repository type e.g. for singleton this is only ever invoked once during the lifecycle of the application
//
// Get provides named access to instances produced by the factory method previously registered. Depending on the
// chosen implementation, this may be a singletons or prototypes, and can either create objects lazily on demand or
// eagerly on startup
//
type Repository interface {
	Register(name string, factory FactoryMethod)
	Get(name string) interface{}
}

// NewSingletonRepository returns an instance of a SingletonRepository
var NewSingletonRepository = internal.NewSingletonRepository

// SingleDependencyRepository returns an instance SingletonRepository populated with a single dependency
func SingleDependencyRepository(key string, factory FactoryMethod) Repository {
	r := NewSingletonRepository()
	r.Register(key, func() interface{} {
		return factory()
	})
	return r
}
