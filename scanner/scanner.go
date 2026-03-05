package scanner

import (
	"sync"
)

func Run(host string, workers int) ([]int) {
	var wg sync.WaitGroup
	chports := make(chan DialResult, 65536)
	chjobs := make(chan int, 65536)
	var ports []int

	for i := 0; i < 65536; i++ {
		chjobs <- i
	}
	close(chjobs)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(host, chjobs, chports, &wg)
	}

	wg.Wait()

	close(chports)

	for result := range chports {
		if(result.Err == nil) {
			ports = append(ports, result.Port)
		}
	}

	return ports
}