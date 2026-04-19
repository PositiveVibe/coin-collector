package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// The target "Digital Territory" 
	// "localhost" means your own computer.
	// You can also try your router IP (e.g., "192.168.1.1")
	target := "localhost"

	fmt.Println("==========================================")
	fmt.Printf("🔍  SAFARI SCANNER: PORT DETECTOR\n")
	fmt.Printf("🌍  Target Territory: %s\n", target)
	fmt.Println("==========================================")
	fmt.Println("Status: Scanning for digital life...")

	count := 0

	// We start our "trek" from port 1 to 1024
	// These are the "System Ports" where the big apps live.
	for port := 1; port <= 1024; port++ {
		
		// Create a connection address like "localhost:80"
		address := fmt.Sprintf("%s:%d", target, port)

		// net.DialTimeout is like a "Quick Knock".
		// We only wait 100 milliseconds for an answer.
		// If we wait too long, the scan will be very slow!
		conn, err := net.DialTimeout("tcp", address, 100*time.Millisecond)

		if err == nil {
			// PORT IS OPEN! 
			fmt.Printf("[+] specimen found on Port %-5d | Status: ACTIVE\n", port)
			conn.Close()
			count++
		}
	}

	fmt.Println("==========================================")
	fmt.Printf("✅  Scan Complete. %d specimens detected.\n", count)
	fmt.Println("==========================================")
}
