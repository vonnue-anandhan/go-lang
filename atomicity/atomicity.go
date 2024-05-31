package atomicExample

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicExample() {
	var sum int64
	wg := &sync.WaitGroup{}

	const gs = 100
	wg.Add(100)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&sum, 1)
			fmt.Println(atomic.LoadInt64(&sum))
			wg.Done()
		}()
	}

	wg.Wait()
}
