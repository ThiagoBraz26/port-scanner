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
	
	ports := scanner.Run(*host)
	
	for _, port := range ports {
		fmt.Println(port)
	}

	fmt.Printf("\n\n Tempo decorrido: %s", time.Since(start))
}