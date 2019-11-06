package endpoint

import (
	"context"
	s "fizzbuzz/service"
	"testing"
)

func Test_makeSingleEndpoint(t *testing.T) {
	tests := map[string]struct {
		svc     s.Service
		number  int
		wantErr bool
	}{
		"default service and valid input": {
			svc:     s.New(),
			number:  1,
			wantErr: false,
		},
		"default service and invalid input": {
			svc:     s.New(),
			number:  -1,
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			e := New(tt.svc)
			s, err := e.Single(context.Background(), SingleRequest{N: tt.number})

			if (err != nil) != tt.wantErr {
				t.Fatalf("wantErr %v, error %v", tt.wantErr, err)
			}

			if (s.(SingleResponse).S == "") != tt.wantErr {
				t.Fatalf("expected valid response, got empty response")
			}
		})
	}
}
