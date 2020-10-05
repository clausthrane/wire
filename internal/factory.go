package internal

// FactoryMethod is a method type used to provide instances to the ServiceRepository
type FactoryMethod func() interface{}
