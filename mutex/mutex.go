package mutexExample

import (
	"fmt"
	"sync"
)

func MutexExample() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	counter := 0
	const limit = 100

	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go func() {
			mut.Lock()
			counter++
			mut.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Count:", counter)
}
