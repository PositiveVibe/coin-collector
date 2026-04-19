package main

import (
	"fmt"
	"net/http"
)

// homeHandler handles the requests to the root path "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// We use a raw string literal to define the HTML content directly in Go.
	// This makes it easy for the student to see the connection between Go and the Web.
	fmt.Fprint(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Safari Command Center</title>
    <!-- Tailwind CSS CDN - No setup required, perfect for learning! -->
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-slate-900 text-white font-sans selection:bg-emerald-500/30">

    <!-- NAVIGATION BAR -->
    <nav class="p-6 border-b border-slate-800 flex justify-between items-center backdrop-blur-md bg-slate-900/80 sticky top-0 z-50">
        <div class="flex items-center space-x-3">
            <span class="text-3xl">🦁</span>
            <h1 class="text-2xl font-bold bg-gradient-to-r from-emerald-400 to-cyan-400 bg-clip-text text-transparent">
                Safari Command Center
            </h1>
        </div>
        <div class="hidden md:flex space-x-6 text-slate-400 text-sm font-medium">
            <div class="flex items-center gap-2">
                <span class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
                <span class="hover:text-emerald-400 cursor-pointer transition-colors">Lab Status: Active</span>
            </div>
            <span class="hover:text-emerald-400 cursor-pointer transition-colors">Sector: 07-Bravo</span>
        </div>
    </nav>

    <!-- MAIN CONTENT -->
    <main class="max-w-6xl mx-auto p-8 md:p-12 mb-20">
        <header class="mb-16">
            <div class="inline-block px-3 py-1 rounded-full bg-emerald-500/10 border border-emerald-500/20 text-emerald-400 text-xs font-bold uppercase tracking-widest mb-4">
                Researcher Dashboard
            </div>
            <h2 class="text-5xl font-extrabold mb-6 tracking-tight">
                Welcome back, <span class="text-emerald-500">Officer.</span>
            </h2>
            <p class="text-xl text-slate-400 max-w-2xl leading-relaxed">
                The digital habitat is currently stable. All sensors are reporting normal activity across the northern grid.
            </p>
        </header>

        <!-- APP GRID -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            
            <!-- PROJECT 1: CENTRAL HUB (COMPLETED) -->
            <div class="group relative bg-slate-800/40 p-10 rounded-3xl border border-slate-700/50 hover:border-emerald-500/50 hover:bg-slate-800/60 transition-all duration-300 shadow-xl overflow-hidden">
                <div class="absolute -right-4 -top-4 text-8xl opacity-5 group-hover:rotate-12 transition-transform duration-500">🦁</div>
                <div class="relative z-10 text-5xl mb-6 group-hover:scale-110 transition-transform duration-300 inline-block">🏡</div>
                <h3 class="text-2xl font-bold mb-3 tracking-tight">Central Hub</h3>
                <p class="text-slate-400 leading-relaxed mb-6">The main control center for all digital safari operations and sensor management.</p>
                <div class="flex items-center text-emerald-400 text-sm font-bold gap-2">
                    <span>LAUNCH PROJECT</span>
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 transform group-hover:translate-x-1 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="9 5l7 7-7 7" />
                    </svg>
                </div>
            </div>

            <!-- PROJECT 2: NETWORK HABITAT (LOCKED) -->
            <div class="group relative bg-slate-900/20 p-10 rounded-3xl border border-dashed border-slate-700/50 opacity-60 hover:opacity-100 transition-opacity">
                <div class="text-5xl mb-6 grayscale group-hover:grayscale-0 transition-all">🔍</div>
                <h3 class="text-2xl font-bold mb-3 tracking-tight text-slate-300">Network Habitat</h3>
                <p class="text-slate-500 leading-relaxed italic mb-6">Discovery scan in progress... waiting for Session 3 deployment.</p>
                <div class="inline-block px-3 py-1 rounded-md bg-slate-800 text-slate-500 text-[10px] font-bold uppercase tracking-widest">
                    LOCKED: SECTOR 3
                </div>
            </div>

            <!-- PROJECT 3: SPECIES VAULT (LOCKED) -->
            <div class="group relative bg-slate-900/20 p-10 rounded-3xl border border-dashed border-slate-700/50 opacity-60 hover:opacity-100 transition-opacity">
                <div class="text-5xl mb-6 grayscale group-hover:grayscale-0 transition-all">🔐</div>
                <h3 class="text-2xl font-bold mb-3 tracking-tight text-slate-300">Species Vault</h3>
                <p class="text-slate-500 leading-relaxed italic mb-6">Encryption protocols loading... hardware key required (Session 5).</p>
                <div class="inline-block px-3 py-1 rounded-md bg-slate-800 text-slate-500 text-[10px] font-bold uppercase tracking-widest">
                    LOCKED: SECTOR 5
                </div>
            </div>

        </div>
    </main>

    <!-- FOOTER -->
    <footer class="fixed bottom-0 w-full p-4 text-center text-slate-600 text-[10px] uppercase tracking-[0.2em]">
        Research License: GO-SAFARI-2024-X11
    </footer>

</body>
</html>
	`)
}

func main() {
	// Register the handler
	http.HandleFunc("/", homeHandler)

	// Start the server
	fmt.Println("----------------------------------------------------------")
	fmt.Println("🦁  Safari Hub Version 1.0 (Tailwind Edition)")
	fmt.Println("🌍  Habitat Server: http://localhost:8080")
	fmt.Println("📝  Status: Researcher Active")
	fmt.Println("----------------------------------------------------------")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
