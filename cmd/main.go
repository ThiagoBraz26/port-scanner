package main

import (
	"flag"
	"fmt"
	"time"
	
	"port-scanner/scanner"
)

func main() {
	start := time.Now()
	host := flag.String("host", "localhost", "Host alvo")
	flag.Parse()

	results := scanner.Run(*host)

	for _, result := range results {
		fmt.Printf("Porta %-6d aberta\n", result)
	}

	fmt.Printf("\n\n Tempo decorrido: %s", time.Since(start))
}