package fbgateway

import (
	"context"

	fbservice "github.com/clausthrane/wire/example/services/foobar"
)

var _ fbservice.Gateway = (*Gateway)(nil)

type (
	// Gateway implements access to the most important number ever
	Gateway struct {
	}
)

// New returns a new instance of the gateway
func New() *Gateway {
	return &Gateway{}
}

// Get returns en important number
func (g *Gateway) Get(context.Context) int {
	return 42
}
