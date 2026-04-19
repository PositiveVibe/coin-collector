# ⚡ Session 4: Super-Scan (The Power of the Pack)

In Session 3, our scanner was slow. It was like one scout walking the whole safari alone. Today, we are going to send the whole **Pack** out at once.

We will use **Goroutines**—Go’s "secret weapon" for doing thousands of things at the same time.

---

### 🧠 Brain Exercise: Parallel Scouting
*Sending multiple scouts at once requires careful coordination!*

**The Problem:** Look at this code snippet. What is the difference between these two lines?
1. `fmt.Println("Scout 1 heading out!")`
2. `go fmt.Println("Scout 2 heading out!")`

**Your Task:** 
1. Create a function called `shout(msg string)`.
2. In `main`, call it twice: once normally, and once with the `go` keyword.
3. **Observation:** Why does the "Go" version sometimes not show up in the terminal? (Hint: The program finishes before the scout has time to shout!)

---

### Goal
Use Goroutines to scan 1,000 ports in less than 1 second.

### 1. The "Pack" (Goroutines)
A Goroutine is a "Mini-Task" that runs in the background. 
- Normal Code: `Scan(port 1)` -> Wait -> `Scan(port 2)` -> Wait...
- Concurrent Code: `go Scan(port 1)` + `go Scan(port 2)` + `go Scan(port 3)`... all at once!

### 2. The Attendance Sheet (WaitGroup)
If we start 1,000 scouts at once, we need to know when they are **all** finished before we can head back to base. In Go, we use a `sync.WaitGroup` to track our scouts.

### 3. The Super-Scanner Code
Create `main.go` in your `code/` folder:

```go
package main

import (
	"fmt"
	"net"
	"sync" // New Library: This helps us sync our scouts!
	"time"
)

func main() {
	target := "localhost"
	var wg sync.WaitGroup // Our attendance sheet

	fmt.Printf("⚡ Starting Super-Scan on: %s\n", target)
	start := time.Now() // Start the timer!

	for port := 1; port <= 1024; port++ {
		
		// 1. Mark that one scout is heading out
		wg.Add(1)

		// 2. The 'go' keyword turns this function into a Goroutine!
		go func(p int) {
			// 3. Mark that this scout returned when finished
			defer wg.Done()

			address := fmt.Sprintf("%s:%d", target, p)
			conn, err := net.DialTimeout("tcp", address, 100*time.Millisecond)

			if err == nil {
				fmt.Printf("[+] PORT %d: OPEN\n", p)
				conn.Close()
			}
		}(port) // We pass the 'port' number into the scout's hand
	}

	// 4. Wait for EVERY scout to check in
	wg.Wait()

	duration := time.Since(start)
	fmt.Printf("------------------------------------------\n")
	fmt.Printf("✅ Super-Scan Complete in %v!\n", duration)
}
```

---

### 🚀 Performance Check
1. Run `go run main.go`.
2. Notice how much faster it is! You should see the scan finish in just a fraction of a second.
3. **Compare:** The old scanner (Session 3) took over 1 minute. This one takes less than 1 second. That is the power of Go!

### 🎓 Task: The "Thousand-Scout" Challenge
1. Open your Session 1 server.
2. Change the scanner to scan ports `8000` to `9000`.
3. Does the Super-Scanner find your server instantly?
4. **Pro-Tip:** If you see an error saying "too many open files," it means you are being **too fast** for your computer! We will learn how to "Limit the Pack" later.

---

### Summary
1. **Goroutines** (`go func()`) let you run many tasks simultaneously.
2. **sync.WaitGroup** makes the main program wait for everyone to finish.
3. Concurrency is what makes Go the favorite language for professional server engineers.
