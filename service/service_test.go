package service

import (
	"testing"
)

func Test_service_Single(t *testing.T) {
	tests := map[string]struct {
		number  int
		output  string
		wantErr bool
	}{
		"zero": {
			number:  0,
			output:  "",
			wantErr: true,
		},
		"one": {
			number: 1,
			output: "1",
		},
		"two": {
			number: 2,
			output: "2",
		},
		"three": {
			number: 3,
			output: "fizz",
		},
		"five": {
			number: 5,
			output: "buzz",
		},
		"fifteen": {
			number: 15,
			output: "fizzbuzz",
		},
		"eighteen": {
			number: 18,
			output: "fizz",
		},
		"ninety": {
			number: 90,
			output: "fizzbuzz",
		},
		"ninety eight": {
			number: 98,
			output: "98",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			svc := &service{}
			r, err := svc.Single(tt.number)

			if (err != nil) != tt.wantErr {
				t.Fatal("expected error, got none")
			}

			if r != tt.output {
				t.Fatalf("expected \"%s\" with number %d, got \"%s\"", tt.output, tt.number, r)
			}
		})
	}
}
