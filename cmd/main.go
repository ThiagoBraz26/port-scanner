package main

import (
	"flag"
	"fmt"
	"time"

	"port-scanner/scanner"
)

func main() {
	start := time.Now()

	host := flag.String("host", "localhost", "Target host")
	workers := flag.Int("workers", 500, "Numbers of workers")
	timeout := flag.Duration("timeout", 1*time.Second, "Timeout limit")
	flag.Parse()

	results := scanner.Run(*host, *workers, *timeout)

	for _, result := range results {
		fmt.Printf("Porta %-6d ABERTA\n", result)
	}

	fmt.Printf("\n\n Tempo decorrido: %s", time.Since(start))
}
