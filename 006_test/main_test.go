package main

import "testing"

var tests = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{name: "valid-data", divident: 100.0, divisor: 10.0, expected: 10.0, isErr: false},
	{name: "invalid-data", divident: 10.0, divisor: 0.0, expected: 0.0, isErr: true},
	{name: "expect-5", divident: 50.0, divisor: 10.0, expected: 5.0, isErr: false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.divident, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f, but got %f", tt.expected, got)
		}
	}
}
