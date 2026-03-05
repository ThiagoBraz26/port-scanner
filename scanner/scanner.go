package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func Run(host string) ([]int) {
	var wg sync.WaitGroup
	chports := make(chan DialResult, 65535)
	var ports []int

	for port := 0; port <= 65535; port++ {
		wg.Add(1)
		go func(host string, port int) {
			defer wg.Done()
			conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, fmt.Sprintf("%d", port)), 1*time.Second)
			if err == nil {
				conn.Close()
			}
			
			chports <- DialResult{err, port}
		}(host, port)
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