package testing

import "testing"

func TestMySum(t *testing.T) {
	// Simple test
	// want := 5
	// got := MySum(2, 3)

	// if got != want {
	// 	t.Errorf("Expected %d, got %d", want, got)
	// }

	// Table test - write a series of test to run, allowing us to test a variety of situations
	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{input: []int{2, 3}, want: 5},
		{input: []int{-1, 0}, want: -1},
		{input: []int{-2, -2}, want: -4},
		{input: []int{-2, 6}, want: 4},
	}

	for _, test := range tests {
		got := MySum(test.input...)
		want := test.want
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	}
}
