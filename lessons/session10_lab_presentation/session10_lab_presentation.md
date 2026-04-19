# 🏆 Session 10: The Lab Presentation (Grand Finale)

Congratulations, Lead Researcher! You have completed the **Advanced Systems Series**. You started with a simple website and ended with a tool that can scan networks, encrypt data, and talk to real-world APIs.

Today is the **Digital Safari Symposium**. We are going to build one final, polished app that brings all your skills together.

---

### 🧠 Brain Exercise: The Project List
*You've built so much! Let's practice managing a list of all your projects.*

**The Problem:** You want to keep a list of your 3 favorite projects in a **Slice**.

**Your Task:** 
1. Create a `[]string` slice called `myProjects`.
2. Add your top 3 projects (e.g., "Hub", "Scanner", "Vault").
3. Use a `for range` loop to print each one with its number (e.g., `1. Hub`, `2. Scanner`).
4. **Pro-Tip:** Slices are like arrays, but they can grow! Try using `append(myProjects, "Calculator")` to add a fourth one.

---

### Goal
Integrate your "Safari Skills" into a final "Scientific Report" dashboard.

### 1. The "Big Picture" (Integration)
Wait, why did we learn all these things?
- **Web Servers:** To show our data to the world.
- **Tailwind:** To make it look like a million-dollar app.
- **Concurrency:** To handle many animal trackers at once.
- **APIs:** To connect our lab to the rest of the planet.

### 2. The Project Scope
Today’s code is the most complex yet! It includes a "Live Background Worker" (like Session 8) and a "Web Interface" (like Session 1).

### 3. The Final "Symposium" Code
Create `main.go` in your `code/` folder:

```go
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Our scientific data
type LabStatus struct {
	SpecimensFound int
	LastMovement   string
	mu             sync.Mutex // For safety!
}

var stats = LabStatus{SpecimensFound: 0}

func backgroundScanner() {
	for {
		time.Sleep(3 * time.Second)
		stats.mu.Lock()
		stats.SpecimensFound++
		stats.LastMovement = time.Now().Format("15:04:05")
		stats.mu.Unlock()
	}
}

func main() {
	// Start our background researcher
	go backgroundScanner()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		stats.mu.Lock()
		count := stats.SpecimensFound
		last := stats.LastMovement
		stats.mu.Unlock()

		// FINAL TAILWIND DASHBOARD
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Safari Symposium</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-black text-white flex items-center justify-center min-h-screen font-sans">
    <div class="text-center p-12 border-2 border-emerald-500 rounded-3xl shadow-[0_0_50px_rgba(16,185,129,0.2)]">
        <h1 class="text-6xl font-black mb-4">🏆 SERIES COMPLETE</h1>
        <p class="text-emerald-400 uppercase tracking-[0.5em] mb-12">Digital Safari Command Center</p>
        
        <div class="grid grid-cols-2 gap-8 text-left">
            <div class="bg-emerald-500/10 p-6 rounded-xl border border-emerald-500/30">
                <div class="text-xs text-emerald-500 font-bold mb-1">TOTAL SPECIMENS</div>
                <div class="text-4xl font-mono">%d</div>
            </div>
            <div class="bg-emerald-500/10 p-6 rounded-xl border border-emerald-500/30">
                <div class="text-xs text-emerald-500 font-bold mb-1">LAST SIGHTING</div>
                <div class="text-4xl font-mono">%s</div>
            </div>
        </div>
        
        <p class="mt-12 text-slate-500 italic">"Scientific knowledge is the greatest safari of all."</p>
    </div>

    <script>
        // Refresh every 3 seconds to see live background data
        setTimeout(() => { location.reload(); }, 3000);
    </script>
</body>
</html>
		`, count, last)
	})

	fmt.Println("🏆 Symposium is LIVE at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
```

---

### 🚀 Graduation Task
1. Run the code. You should see a "Glowing" dashboard.
2. Every 3 seconds, the "Total Specimens" will grow—even if you don't refresh the page! (Wait... how? Because the background worker is running while the website is waiting!)
3. **The Final Challenge:** Can you add a button to the dashboard that links back to the **Safari Hub** you built in Session 1?

### 🎓 You are now a Go Developer!
You have completed:
- ✅ **Web Engineering** (Servers & UI)
- ✅ **Systems Engineering** (TCP & Networks)
- ✅ **Security** (Encryption)
- ✅ **Data Engineering** (JSON & Persistence)
- ✅ **Concurrency** (Goroutines & Mutexes)

**Keep exploring, researcher! 🦁🌳**
