# 📡 Session 7: WebSocket Field-Feeds (Real-Time Updates)

In our previous lessons, if you wanted new data, you had to "Ask" the server (by refreshing or clicking a button). This is called **Pulling**.

Today, we learn how to **Push**. We will use **WebSockets** to create a "Live Pipe" between the safari field and your dashboard. When a researcher sees an animal, your screen will update instantly!

---

### 🧠 Brain Exercise: The Tracker Map
*Before we push live data, we need a place to store it on the server!*

**The Problem:** You need a way to store which animal is in which sector. Use a **Map** for this.

**Your Task:** 
1. Create a `map[string]string` called `tracker`.
2. Add a lion to Sector 5: `tracker["Lion"] = "Sector 5"`.
3. Add a Zebra to Sector 2.
4. Print the location of the Lion by calling its name: `fmt.Println(tracker["Lion"])`.

---

### Goal
Build a real-time feed that "pushes" animal sightings to your browser using WebSockets.

### 1. The Upgrade (HTTP vs WebSocket)
- **HTTP:** Like a letter. You send it, they reply, and the connection is CLOSED.
- **WebSocket:** Like a phone call. Once you connect, the line stays OPEN for as long as you want. Both sides can talk at any time.

### 2. The Tools (Gorilla WebSocket)
We will use another external library to handle the "upgrade". Run this:
```bash
go get github.com/gorilla/websocket
```

### 3. The Live Feed Code
Create `main.go` in your `code/` folder:

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket" // Our real-time library
)

// We need an 'Upgrader' to turn a normal website visitor into a WebSocket caller
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow everyone
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Upgrade the connection
	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	fmt.Println("📡 New tracking collar connected!")

	// 2. Start pushing fake "Live Data" every 2 seconds
	animals := []string{"🐘 Elephant", "🦓 Zebra", "🦒 Giraffe", "🐆 Leopard"}
	
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("LIVE SIGHTING: %s detected at sector %d", animals[i%4], i+1)
		
		// 3. Send the message through the pipe
		conn.WriteMessage(websocket.TextMessage, []byte(message))
		
		time.Sleep(2 * time.Second)
	}
}

func main() {
	http.HandleFunc("/ws", socketHandler)
	
	// We also serve a tiny HTML page to listen to the feed
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			<h1>Safari Live Feed</h1>
			<div id="output" style="font-family: monospace; background: #000; color: #0f0; padding: 20px;"></div>
			<script>
				const socket = new WebSocket("ws://localhost:8080/ws");
				socket.onmessage = (event) => {
					document.getElementById("output").innerHTML += event.data + "<br>";
				};
			</script>
		`)
	})

	fmt.Println("📡 Live Feed Server starting on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
```

---

### 🚀 Research Task
1. Run the code and visit `http://localhost:8080`.
2. Watch the screen. Every 2 seconds, a new animal appears **automatically** without you clicking anything!
3. **The "Multiple Eyes" Test:** Open the website in two differnet browser tabs at the same time. Do they both see the same feed?

### 🎓 Task: The "Emergency" Challenge
- Can you change the code so that if the animal is a "Leopard", it adds a `⚠️ DANGER` tag to the message?
- **Pro-Tip:** WebSockets are what power the "Typing..." bubbles in iMessage and the live price charts on stock markets.

---

### Summary
1. **WebSockets** allow for two-way, always-on communication.
2. **Upgrading** is the process of turning an HTTP request into a WebSocket.
3. **conn.WriteMessage** is how the server "pushes" data to the student's browser.
