package middleware

import (
	"context"
	fbe "fizzbuzz/endpoint"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

// Middleware is meant to be used on Endpoints and enables calling middlewares on before and after processing incoming requests.
type Middleware func(endpoint.Endpoint) endpoint.Endpoint

// Logging returns a Middleware for logging on Endpoints.
func Logging(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("params", request.(fbe.SingleRequest).N)
			return next(ctx, request)
		}
	}
}
