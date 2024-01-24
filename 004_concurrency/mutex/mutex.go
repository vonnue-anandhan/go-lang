package mutexu

import (
	"fmt"
	"sync"
)

func test() {
	fmt.Println("Race condition - learncodeonline.in")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var scores = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()

		fmt.Println("One r")

		mut.Lock()
		scores = append(scores, 1)
		mut.Unlock()

	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()

		fmt.Println("Two r")

		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()

	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()

		fmt.Println("Three r")

		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()

		fmt.Println("Four r")

		mut.RLock()
		fmt.Println(scores)
		mut.RUnlock()

	}(wg, mut)

	wg.Wait()
	fmt.Println(scores)
}
