package errors

import "fmt"

func myPanic() {
	panic("something went wrong")
}

func TestRecover() {
	// recover must be called within a deferred function.
	// When the enclosing function panics, the defer will activate and a recover call within it will catch the panic.
	defer func() {
		// The return value of recover is the error raised in the call to panic.
		if r := recover(); r != nil {
			fmt.Println("Recovered, error:\n", r)
			fmt.Println("After recover....")
		}
	}()

	myPanic()

	fmt.Println("After myPanic()")

}
