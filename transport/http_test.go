package transport

import (
	"fizzbuzz/endpoint"
	"fizzbuzz/service"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/go-kit/kit/log"
)

func Test_NewHTTPHandler(t *testing.T) {
	tests := map[string]struct {
		endpoints    *endpoint.Set
		path         string
		method       string
		body         io.Reader
		responseCode int
	}{
		"responds to POST on /single": {
			endpoints:    endpoint.New(service.New()),
			path:         "/single",
			method:       http.MethodPost,
			body:         http.NoBody,
			responseCode: http.StatusInternalServerError,
		},
		"doesn't respond to GET on /single": {
			endpoints:    endpoint.New(service.New()),
			path:         "/single",
			method:       http.MethodGet,
			body:         http.NoBody,
			responseCode: http.StatusMethodNotAllowed,
		},
		"responds to GET on /health": {
			endpoints:    endpoint.New(service.New()),
			path:         "/health",
			method:       http.MethodGet,
			body:         http.NoBody,
			responseCode: http.StatusOK,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			h := NewHTTPHandler(*tt.endpoints, log.NewNopLogger())
			req, err := http.NewRequest(tt.method, tt.path, tt.body)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.responseCode {
				t.Fatalf("expected %d, got %d", tt.responseCode, status)
			}
		})
	}
}
