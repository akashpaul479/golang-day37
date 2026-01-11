package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func FetchData(source string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fetching from:", source)
	time.Sleep(2 * time.Second)
	fmt.Println("Done:", source)
}

func SimulatedAPICall() {
	var wg sync.WaitGroup

	sources := []string{"API-1", "API-2", "API-3"}

	fmt.Println("fetching started!")
	for _, src := range sources {
		wg.Add(1)
		go FetchData(src, &wg)
	}
	wg.Wait()
	fmt.Println("All data fetched!")
}
