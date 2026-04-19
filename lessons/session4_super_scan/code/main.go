package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	// Destination Territory
	target := "localhost"

	// THE WAITGROUP (Attendance Sheet)
	// This tells the main program to wait until all scouts (goroutines) are done.
	var wg sync.WaitGroup

	fmt.Println("==========================================")
	fmt.Printf("⚡  SUPER-SCAN: CONCURRENT DETECTOR\n")
	fmt.Printf("🌍  Target Territory: %s\n", target)
	fmt.Println("==========================================")

	start := time.Now() // Start our stopwatch

	// We are sending scouts for ports 1 to 1024
	for port := 1; port <= 1024; port++ {
		
		// Add 1 to our attendance sheet (one scout leaving)
		wg.Add(1)

		// Start a Goroutine (The 'go' keyword is the magic!)
		go func(p int) {
			// Mark as done when the scout returns (even if they fail)
			defer wg.Done()

			address := fmt.Sprintf("%s:%d", target, p)
			
			// We use a short timeout because we are checking MANY ports at once.
			conn, err := net.DialTimeout("tcp", address, 200*time.Millisecond)

			if err == nil {
				// SUCCESS: Specimen detected!
				fmt.Printf("[+] specimen found on Port %-5d | Status: ACTIVE\n", p)
				conn.Close()
			}
		}(port) // Pass the current 'port' loop value into the scout's hand
	}

	// WAIT for all scouts to check back in
	wg.Wait()

	// Calculate and show total time
	duration := time.Since(start)
	fmt.Println("==========================================")
	fmt.Printf("✅  Super-Scan Complete in %v!\n", duration)
	fmt.Println("==========================================")
}
