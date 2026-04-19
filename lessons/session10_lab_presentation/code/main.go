package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// LabStatus tracks our graduation progress and research activity
type LabStatus struct {
	TotalEntries int
	StartTime    time.Time
	CurrentTask  string
	mu           sync.Mutex
}

var stats = LabStatus{
	TotalEntries: 0,
	StartTime:    time.Now(),
	CurrentTask:  "Calibrating Sensors...",
}

// simulateResearch runs in the background like in Session 8
func simulateResearch() {
	tasks := []string{
		"Optimizing Goroutines...",
		"Syncing Database...",
		"Encrypting Location Data...",
		"Upgrading WebSocket Feeds...",
		"Polling Global APIs...",
	}
	
	i := 0
	for {
		time.Sleep(4 * time.Second)
		
		stats.mu.Lock()
		stats.TotalEntries++
		stats.CurrentTask = tasks[i%len(tasks)]
		stats.mu.Unlock()
		
		i++
	}
}

func main() {
	// Start the background researcher
	go simulateResearch()

	// HOMEPAGE: The Scientific Symposium Dashboard
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		stats.mu.Lock()
		count := stats.TotalEntries
		uptime := time.Since(stats.StartTime).Round(time.Second)
		task := stats.CurrentTask
		stats.mu.Unlock()

		// FINAL GRADUATION UI
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>🦁 Safari Graduation Symposium</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Outfit:wght@400;700;900&display=swap" rel="stylesheet">
    <style>
        body { font-family: 'Outfit', sans-serif; }
    </style>
</head>
<body class="bg-[#020617] text-white min-h-screen flex items-center justify-center p-6">

    <!-- GRADUATION FRAME -->
    <div class="max-w-4xl w-full bg-slate-900/50 border border-emerald-500/20 p-12 rounded-[3rem] backdrop-blur-3xl shadow-[0_0_100px_rgba(16,185,129,0.1)] relative overflow-hidden text-center">
        
        <!-- DECORATIVE BACKGROUND GLOW -->
        <div class="absolute -top-24 -left-24 w-64 h-64 bg-emerald-500/10 blur-[100px] rounded-full"></div>
        <div class="absolute -bottom-24 -right-24 w-64 h-64 bg-cyan-500/10 blur-[100px] rounded-full"></div>

        <header class="relative z-10 mb-12">
            <div class="text-6xl mb-6">🦁</div>
            <h1 class="text-7xl font-black tracking-tight mb-4 bg-gradient-to-br from-white to-slate-400 bg-clip-text text-transparent italic">
                GRADUATED.
            </h1>
            <p class="text-emerald-400 font-bold tracking-[0.4em] uppercase text-sm">
                Digital Safari Series Complete
            </p>
        </header>

        <!-- STATS GRID -->
        <div class="relative z-10 grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
            
            <div class="bg-slate-800/50 p-6 rounded-2xl border border-slate-700/50">
                <div class="text-[10px] text-slate-500 font-black uppercase tracking-widest mb-2">Total Logs</div>
                <div class="text-4xl font-bold text-emerald-400">%d</div>
            </div>

            <div class="bg-slate-800/50 p-6 rounded-2xl border border-slate-700/50">
                <div class="text-[10px] text-slate-500 font-black uppercase tracking-widest mb-2">System Uptime</div>
                <div class="text-4xl font-bold text-cyan-400">%s</div>
            </div>

            <div class="bg-slate-800/50 p-6 rounded-2xl border border-slate-700/50">
                <div class="text-[10px] text-slate-500 font-black uppercase tracking-widest mb-2">Active Task</div>
                <div class="text-sm font-bold text-slate-300 truncate">%s</div>
            </div>

        </div>

        <div class="relative z-10 p-1 rounded-full bg-gradient-to-r from-emerald-500 to-cyan-500 inline-block mb-8">
            <div class="px-8 py-3 rounded-full bg-slate-950 hover:bg-transparent transition-colors duration-500">
                <span class="text-xs font-black tracking-widest uppercase">Go Systems Engineer: Certified</span>
            </div>
        </div>

        <p class="text-slate-500 text-xs mt-8">
            Series: Digital Safari & Lab (2024) • Verification ID: GO-SAFARI-FINAL-010
        </p>

    </div>

    <script>
        // This simple script tells the browser to refresh every 3 seconds
        // so we can see the 'Total Logs' grow in real-time!
        setTimeout(() => {
            window.location.reload();
        }, 3000);
    </script>

</body>
</html>
		`, count, uptime, task)
	})

	fmt.Println("----------------------------------------------------------")
	fmt.Println("🏆  GRADUATION DAY: SYMPOSIUM IS LIVE")
	fmt.Println("🌍  Final Report URL: http://localhost:8080")
	fmt.Println("📝  Status: Researcher Promoted to Lead Engineer")
	fmt.Println("----------------------------------------------------------")
	
	http.ListenAndServe(":8080", nil)
}
