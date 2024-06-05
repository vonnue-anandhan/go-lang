package testing

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("James")
	want := "Welcome my dear James"
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
