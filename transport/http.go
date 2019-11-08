package transport

import (
	"context"
	"encoding/json"
	"fizzbuzz/endpoint"
	"net/http"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// NewHTTPHandler builds and returns an HTTP handler with all the routing configured.
func NewHTTPHandler(endpoints endpoint.Set, logger kitlog.Logger) http.Handler {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
	}

	singleHandler := kithttp.NewServer(
		endpoints.Single,
		decodeSingleRequest,
		encodeResponse,
		options...,
	)

	healthCheckHandler := kithttp.NewServer(
		endpoints.HealthCheck,
		decodeHealthCheckRequest,
		encodeResponse,
		options...,
	)
	r := mux.NewRouter()
	r.Handle("/single", singleHandler).Methods(http.MethodPost)
	r.Handle("/health", healthCheckHandler).Methods(http.MethodGet)

	return r
}

func decodeSingleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.SingleRequest
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req)
	return req, err
}

func decodeHealthCheckRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoint.HealthCheckRequest{}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
