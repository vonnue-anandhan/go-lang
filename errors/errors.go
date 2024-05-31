package errors

import (
	"errors"
	"fmt"
	"log"
)

func TestErrors() {
	err := getErrors()
	defer deferredFunction()

	if err != nil {
		// log.Fatalln(err.Error()) // - deferred functions ignored

		log.Panicln(err.Error()) // - deferred functions considered
	}

}

func getErrors() error {
	return errors.New("something went wrong")
}

func deferredFunction() {
	fmt.Println("*************** deferred function ran ***************")
}
