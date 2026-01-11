package goroutines

import (
	"fmt"
	"sync"
)

func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker", id, "started")
}

func WaitGroup() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished!")
}
