# 🦁 Session 1: The Digital Safari Hub (Go + Tailwind)

Welcome to the **Advanced Systems Series**! Over the next 10 sessions, we aren't just building "scripts"—we are building a full **Digital Safari Command Center**. This is a dashboard where you can monitor "digital animals" (devices on your network), protect research data, and see live migration feeds.

Today, we build the "Base Camp": Your central dashboard.

---

### 🧠 Brain Exercise: The Species Decoder
*Before we start, let's test your memory on Go Data Types!*

**The Problem:** Look at the data below. What "Type" should each variable be? (Choose from: `string`, `int`, `bool`, `float64`)
1. `animalName := "Simba"` -> **Type?**
2. `isEndangered := true` -> **Type?**
3. `populationCount := 250` -> **Type?**
4. `averageWeight := 190.5` -> **Type?**

**Your Task:** Write a small `warmup.go` file that creates these 4 variables and prints them using `fmt.Println`. If you get it right, you're ready for the dashboard!

---

### Goal
Create a professional Go web server that uses **Tailwind CSS** to look like a premium app.

### 1. Project Setup
We are starting fresh! Open your terminal and create a new directory for our Command Center:
```bash
mkdir safari-hub
cd safari-hub
go mod init safari-hub
```

### 2. The "Base Camp" Server
Create a file named `main.go`. This server will serve our dashboard.

```go
package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// We are using a 'Raw String' (backticks) to send a whole website!
	fmt.Fprint(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Safari Command Center</title>
    <!-- We are importing Tailwind CSS - the industry favorite for styling! -->
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-slate-900 text-white font-sans">

    <!-- NAVIGATION BAR -->
    <nav class="p-6 border-b border-slate-800 flex justify-between items-center">
        <h1 class="text-2xl font-bold bg-gradient-to-r from-emerald-400 to-cyan-400 bg-clip-text text-transparent">
            🦁 Safari Command Center
        </h1>
        <div class="space-x-4 text-slate-400 text-sm">
            <span class="hover:text-emerald-400 cursor-pointer">Status: Online</span>
            <span class="hover:text-emerald-400 cursor-pointer">Lab: Sector 7</span>
        </div>
    </nav>

    <!-- MAIN CONTENT -->
    <main class="max-w-6xl mx-auto p-12">
        <header class="mb-12">
            <h2 class="text-4xl font-extrabold mb-4 text-emerald-500">Welcome, Researcher.</h2>
            <p class="text-slate-400">Your digital tools are ready for field deployment.</p>
        </header>

        <!-- APP GRID (This is where our future projects will live!) -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            
            <!-- HUB CARD (The one we are building today) -->
            <div class="bg-slate-800 p-8 rounded-2xl border border-slate-700 hover:border-emerald-500 transition-all cursor-pointer group">
                <div class="text-4xl mb-4 group-hover:scale-110 transition-transform">🏡</div>
                <h3 class="text-xl font-bold mb-2">Central Hub</h3>
                <p class="text-slate-400 text-sm">The main control center for all safari operations.</p>
            </div>

            <!-- FUTURE CARDS (PLACEHOLDERS) -->
            <div class="bg-slate-800/50 p-8 rounded-2xl border border-dashed border-slate-700 opacity-50">
                <div class="text-4xl mb-4">🔍</div>
                <h3 class="text-xl font-bold mb-2">Network Habitat</h3>
                <p class="text-slate-400 text-sm">Scanning for digital life in the local environment...</p>
            </div>

            <div class="bg-slate-800/50 p-8 rounded-2xl border border-dashed border-slate-700 opacity-50">
                <div class="text-4xl mb-4">🔐</div>
                <h3 class="text-xl font-bold mb-2">Species Vault</h3>
                <p class="text-slate-400 text-sm">Securing sensitive discovery logs...</p>
            </div>

        </div>
    </main>

</body>
</html>
	`)
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("🌳 Safari Command Center starting on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
```

---

### 🚀 Decoding the "Tailwind" Magic
Take a look at the `class` names in the code. This is **Tailwind CSS**. Instead of writing separate CSS files, we use "Utility Classes":
- `bg-slate-900`: Sets the background to a dark grey/blue.
- `text-emerald-500`: Makes the text a vibrant green.
- `grid grid-cols-3`: Automatically arranges our "App Cards" into 3 columns.
- `hover:border-emerald-500`: This makes the card glow green when you touch it!

### 🎓 Task: The "Dark Mode" Challenge
Try changing the `bg-slate-900` to `bg-black`. Notice how much more "Pro" it looks! Also, try changing the `emerald` color to `amber` or `rose`. 

**The Goal:** Pick a "Scientific Theme" (Green, Blue, or Orange) and make the whole dashboard match!

---

### Pro-Tip: The Backtick (`)
In Go, we use `backticks` to create multi-line strings. This is perfect for HTML because it lets us paste the entire website inside our Go code without any errors!
