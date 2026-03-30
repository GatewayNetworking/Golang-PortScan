package main

import (
	"fmt"
	"net"
	"time"
)

func scanPort(host string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp4", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	var input int
	input, err := fmt.Scan("%.2f", &input)

	host := "120"
	timeout := 1 * time.Second

	for port := 1; port <= 1024; port++ {
		if scanPort(host, port, timeout) {
			fmt.Printf("Port: %d is OPEN\n", port)
		}
	}
}
