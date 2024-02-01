package mutex

import (
	"fmt"
	"sync"
)

func mutex() {
	gs := 100
	counter := 0

	mut := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mut.Lock()

			counter++

			mut.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
