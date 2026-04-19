# 🔍 Session 3: Network Discovery (Scanning the Habitat)

Today, we leave the "Base Camp" and head into the field! In a real safari, you look for animal tracks. In systems engineering, we look for **Open Ports** and **IP Addresses**.

Every computer on your WiFi is like an animal in the safari. To talk to them, we need to find them first.

---

### 🧠 Brain Exercise: The Resource Counter
*Before we trek into the network, let's practice counting!*

**The Problem:** You need to count from `1` up to `20` and print each number.

**Your Task:** 
1. Write a `for` loop that starts at `1`.
2. The loop should stop when it reaching `21`.
3. Inside the loop, print the number.

**Bonus Tip:** Remember the `i++` trick? It's the shortest way to say "add one to i".

---

### Goal
Build a Go tool that scans a computer to see which "doors" (ports) are open.

### 1. The Territory (IPs) and the Burrow (Ports)
- **IP Address:** This is the "GPS Location" of a computer (e.g., `192.168.1.5`).
- **Port:** This is the "Doorway" to a specific app. 
    - Port `80` or `8080` is usually a Website.
    - Port `22` is a Secure Shell (Remote Control).
    - Port `443` is a Secure Website.

### 2. The TCP Handshake (The "Hello")
To check if a port is open, Go sends a "TCP Handshake". It's like knocking on a door:
1. We say: "Hi, anyone home?" (**SYN**)
2. They say: "Yes! Come in!" (**SYN-ACK**)
3. We say: "Thanks, I'm just checking!" (**ACK**)

If they don't answer, the port is **Closed**.

### 3. The Habitat Scanner
Create `main.go` in your `code/` folder:

```go
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// The target "Digital Territory" (usually your own computer)
	target := "localhost"

	fmt.Printf("🔍 Starting Habitat Scan on: %s\n", target)
	fmt.Println("------------------------------------------")

	// We are going to scan ports 1 to 1024 (the most common ones)
	for port := 1; port <= 1024; port++ {
		
		// Create a "Address" string like "localhost:80"
		address := fmt.Sprintf("%s:%d", target, port)

		// Try to "Knock" on the port. We wait only 100 milliseconds.
		conn, err := net.DialTimeout("tcp", address, 100*time.Millisecond)

		if err == nil {
			// SUCCESS! The door opened!
			fmt.Printf("[+] PORT %d: OPEN (Specimen Detected!)\n", port)
			conn.Close() // Don't forget to close the door!
		}
	}

	fmt.Println("------------------------------------------")
	fmt.Println("✅ Scan Complete. No more specimens found.")
}
```

---

### 🚀 Running the Scan
1. Run `go run main.go`.
2. It might seem slow... that's because it's checking one port at a time! (We will fix this with "Super-Speed" in the next session).
3. **Question:** Does it find port `8080` from your Session 1 project? (Make sure your Session 1 server is running!)

### 🎓 Task: The "High-Level" Challenge
Most servers use ports in the thousands. 
- Try changing the loop to scan ports `8000` to `9000`. 
- Can you find your Safari Hub from Session 2?
- **Pro-Tip:** If you scan your router's IP (usually `192.168.1.1`), you might find open doors for its settings!

---

### Summary
1. **IP Addresses** identify the device; **Ports** identify the app.
2. **net.DialTimeout** is how we knock on a digital door.
3. If `err == nil`, the port is open and a "Specimen" is listening there.
