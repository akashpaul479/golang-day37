package goroutines

import (
	"fmt"
	"sync"
)

func Worker1(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
	}
}

func WorkerPool() {
	jobs := make(chan int, 5)
	var wg sync.WaitGroup
	// workers
	fmt.Println("processing started!")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go Worker1(i, jobs, &wg)
	}
	// jobs
	for j := 1; j <= 10; j++ {
		jobs <- j
	}

	// Close jons
	close(jobs)

	wg.Wait()
	fmt.Println("All job processed!")
}
