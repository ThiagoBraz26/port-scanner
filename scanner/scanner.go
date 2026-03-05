package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
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
		go func(host string, ports chan int) {
			for port := range ports {
				conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, fmt.Sprintf("%d", port)), 1*time.Second)
				if err == nil {
					conn.Close()
				}
				chports <- DialResult{err, port}
			}
			wg.Done()
		}(host, chjobs)
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