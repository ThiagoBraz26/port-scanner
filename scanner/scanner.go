package scanner

import (
	"fmt"
	"net"
)

func Run(host string) ([]string) {
	var ports []string

	for i := 0; i <= 65536; i++ {
		conn, err := net.Dial("tcp", net.JoinHostPort(host, fmt.Sprintf("%d", i)))
		if err != nil {
			ports = append(ports, fmt.Sprintf("Porta: %-6d FECHADA", i))
			continue
		}

		ports = append(ports, fmt.Sprintf("Porta: %-6d ABERTA", i))
		conn.Close()
	}
	
	return ports
}