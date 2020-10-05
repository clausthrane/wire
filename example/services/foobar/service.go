package fbservice

import "context"

type (
	// Gateway defines the interface of a subsystem of this service type
	Gateway interface {
		Get(context.Context) int
	}

	// Service implements some business logic
	Service struct {
		gateway Gateway
	}
)

// New returns a new Service instances initalized with the provided gateway dependency
func New(gateway Gateway) *Service {
	return &Service{
		gateway: gateway,
	}
}

// GetValue returns some business critical value
func (s *Service) GetValue(ctx context.Context) int {
	return s.gateway.Get(ctx)
}
