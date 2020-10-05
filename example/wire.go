package main

import (
	"context"

	"github.com/clausthrane/wire"
	fbgateway "github.com/clausthrane/wire/example/gateways/foobar"
	fbservice "github.com/clausthrane/wire/example/services/foobar"
)

const (
	_fbService = "fbService"
	_fbGateway = "fbGateway"
)

// Application holds our repository of services that exist throughout the lifespan of the application.
type Application struct {
	wire.Repository
	ctx context.Context
}

// NewApplication initializes our application with the services needed
func NewApplication() (Application, context.CancelFunc) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	app := Application{
		Repository: wire.NewSingletonRepository(),
		ctx:        ctx,
	}

	app.Register(_fbService, func() interface{} {
		return fbservice.New(app.Gateway())
	})

	app.Register(_fbGateway, func() interface{} {
		return fbgateway.New()
	})

	return app, cancelFunc
}

// Context returns the application Context.
// when the application is closing the context is cancelled.
func (a *Application) Context() context.Context {
	return a.ctx
}

// Service provides access to a shared instance of a *fbservice.Service
func (a *Application) Service() *fbservice.Service {
	return a.Get(_fbService).(*fbservice.Service)
}

// Gateway provides access to a shared instance of a *fbgateway.Gateway
func (a *Application) Gateway() *fbgateway.Gateway {
	return a.Get(_fbGateway).(*fbgateway.Gateway)
}
