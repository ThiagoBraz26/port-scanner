package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func worker(host string, ports chan int, results chan DialResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for port := range ports {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, fmt.Sprintf("%d", port)), 1*time.Second)
		if err == nil {
			conn.Close()
		}
		results <- DialResult{err, port}
	}
}
