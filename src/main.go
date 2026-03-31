package main

import (
	"fmt"
	"net"
	"time"
)

func TCPPort(host string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp4", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func UDPPort(host string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("upd", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	var host string

	logo := `
  ██████   ██████  ██       █████  ███    ██  ██████      ██████   ██████  ██████  ████████ ███████  ██████  █████  ███    ██ 
██       ██    ██ ██      ██   ██ ████   ██ ██           ██   ██ ██    ██ ██   ██    ██    ██      ██      ██   ██ ████   ██ 
██   ███ ██    ██ ██      ███████ ██ ██  ██ ██   ███     ██████  ██    ██ ██████     ██    ███████ ██      ███████ ██ ██  ██ 
██    ██ ██    ██ ██      ██   ██ ██  ██ ██ ██    ██     ██      ██    ██ ██   ██    ██         ██ ██      ██   ██ ██  ██ ██ 
 ██████   ██████  ███████ ██   ██ ██   ████  ██████      ██       ██████  ██   ██    ██    ███████  ██████ ██   ██ ██   ████                                               
                                                                                                                              `
	fmt.Printf("%s\n", logo)
	fmt.Print("Enter IP address to scan: ")

	_, err := fmt.Scanln(&host)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	timeout := 1 * time.Second
	fmt.Printf("Scanning %s...\n", host)

	for port := 1; port <= 1024; port++ {
		if TCPPort(host, port, timeout) {
			fmt.Printf("TCP Port: %d is OPEN\n", port)
		}
	}

	for port := 1; port <= 1024; port++ {
		if UDPPort(host, port, timeout) {
			fmt.Printf("UDP Port: %d is OPEN\n", port)

		}
	}
}
