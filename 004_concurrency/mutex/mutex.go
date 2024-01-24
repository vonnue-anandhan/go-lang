package mutex

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

var signals = []string{"test"}

var mut sync.Mutex

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for website %s\n", res.StatusCode, endpoint)
	}
}

func Test() {
	websites := []string{"https://lco.dev", "https://go.dev", "https://google.com", "https://fb.com", "https://github.com"}

	for _, website := range websites {
		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)
}
