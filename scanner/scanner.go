package scanner

import (
	"sync"
	"time"
)

func Run(host string, workers int, timeout time.Duration) []int {
	var wg sync.WaitGroup
	ports := make(chan DialResult, 65536)
	jobs := make(chan int, 65536)
	var results []int

	for i := range 65536 {
		jobs <- i
	}
	close(jobs)

	for range workers {
		wg.Add(1)
		go worker(host, jobs, ports, timeout, &wg)
	}

	wg.Wait()

	close(ports)

	for result := range ports {
		if result.Err == nil {
			results = append(results, result.Port)
		}
	}

	return results
}
