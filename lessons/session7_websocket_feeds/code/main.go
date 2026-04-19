package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// The Upgrader tells Go how to transition from HTTP to a WebSocket.
var upgrader = websocket.Upgrader{
	// This setting allows the browser to connect to our server.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// wsHandler handles the live data connection
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 1. "Upgrade" the connection from HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()

	fmt.Println("📡 [NEW CONNECTION] Tracking collar calibrated.")

	animals := []string{"Elephant 🐘", "Zebra 🦓", "Giraffe 🦒", "Lion 🦁", "Rhino 🦏"}
	locations := []string{"Waterhole", "Savannah", "North Gate", "Acacia Grove"}

	for {
		// 2. Pick a random animal and location
		animal := animals[rand.Intn(len(animals))]
		loc := locations[rand.Intn(len(locations))]
		
		msg := fmt.Sprintf("[%s] Live Alert: %s spotted at %s!", 
			time.Now().Format("15:04:05"), animal, loc)

		// 3. PUSH the message to the browser
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			fmt.Println("Connection closed by browser.")
			break
		}

		// 4. Wait a few seconds before the next update
		time.Sleep(3 * time.Second)
	}
}

func main() {
	// Route for the Live Socket
	http.HandleFunc("/ws", wsHandler)

	// Simple Homepage with JavaScript to listen to the socket
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Live Safari Feed</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-slate-900 text-white p-12 font-sans">
    <div class="max-w-xl mx-auto">
        <h1 class="text-3xl font-bold mb-6 text-emerald-400">📡 Live Tracking Feed</h1>
        <div id="status" class="mb-4 text-sm text-slate-500 italic">Connecting to field sensors...</div>
        
        <div id="feed" class="bg-black/50 p-6 rounded-xl border border-slate-700 h-[400px] overflow-y-auto space-y-2 font-mono text-sm">
            <!-- MESSAGES APPEAR HERE -->
        </div>
        
        <p class="mt-6 text-slate-500 text-xs uppercase tracking-widest text-center">Researcher ID: SAFARI-ALPHA</p>
    </div>

    <script>
        const feed = document.getElementById('feed');
        const status = document.getElementById('status');
        
        // Connect to our Go WebSocket
        const socket = new WebSocket("ws://localhost:8080/ws");

        socket.onopen = () => {
            status.innerText = "ONLINE: Receiving live field data.";
            status.classList.remove('text-slate-500');
            status.classList.add('text-emerald-500');
        };

        socket.onmessage = (event) => {
            const div = document.createElement('div');
            div.className = "p-2 border-l-2 border-emerald-500 bg-emerald-500/5 animate-pulse";
            div.innerText = event.data;
            feed.prepend(div); // Add new messages at the top
        };

        socket.onclose = () => {
            status.innerText = "OFFLINE: Connection lost.";
            status.className = "text-red-500 font-bold";
        };
    </script>
</body>
</html>
		`)
	})

	fmt.Println("🌳 Safari Field-Feed is LIVE at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
