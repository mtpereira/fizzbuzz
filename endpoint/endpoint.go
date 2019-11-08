package endpoint

import (
	"context"
	s "fizzbuzz/service"

	"github.com/go-kit/kit/endpoint"
)

// Set is an helper struct to share all the service's endpoints.
type Set struct {
	Single      endpoint.Endpoint
	HealthCheck endpoint.Endpoint
}

// New returns an Endpoints struct with all the svc's endpoints.
func New(svc s.Service) *Set {
	return &Set{
		Single:      makeSingleEndpoint(svc),
		HealthCheck: makeHealthCheckEndpoint(svc),
	}
}

// SingleRequest represents a request to the Single endpoint.
type SingleRequest struct {
	N int `json:"n"`
}

// SingleResponse represents a response from the Single endpoint.
type SingleResponse struct {
	S   string `json:"s"`
	Err string `json:"err,omitempty"`
}

func makeSingleEndpoint(svc s.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(SingleRequest)
		s, err := svc.Single(req.N)
		if err != nil {
			return SingleResponse{s, err.Error()}, err
		}
		return SingleResponse{s, ""}, nil
	}
}

// HealthCheckRequest represents a request to the HealthCheck endpoint.
type HealthCheckRequest struct{}

// HealthCheckResponse represents a response from the HealthCheck endpoint.
type HealthCheckResponse struct {
	Status bool `json:"status"`
}

func makeHealthCheckEndpoint(svc s.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		s := svc.HealthCheck()
		return HealthCheckResponse{s}, nil
	}
}
