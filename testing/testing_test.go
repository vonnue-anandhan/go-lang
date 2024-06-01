package testing

import "testing"

func TestMySum(t *testing.T) {
	want := 5
	got := MySum(2, 3)

	if got != want {
		t.Errorf("Expected %d, got %d", want, got)
	}
}
