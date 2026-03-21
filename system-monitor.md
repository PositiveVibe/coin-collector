# 📊 Project: macOS System Monitor (Go + Web)

In this project, we are going to build a **Live Dashboard** that monitors your Mac’s brain (CPU), memory (RAM), and storage (Disk). This is a professional-grade utility used by engineers to keep servers running smoothly.

---

<details>
<summary><b>🛠️ Session 1: Installing the "Sensors"</b></summary>

### Goal
In Go, we use professional libraries to talk to hardware. We will use `gopsutil`, a famous library for system monitoring.

### 1. Project Setup
Open your terminal on your Mac and run:
```bash
mkdir mac-monitor
cd mac-monitor
go mod init mac-monitor
```

### 2. Download the Libraries
Run these three commands to get the "sensors" for your code:
```bash
go get github.com/shirou/gopsutil/v3/cpu
go get github.com/shirou/gopsutil/v3/mem
go get github.com/shirou/gopsutil/v3/disk
```

### 3. The RAM Test
Create `main.go` and see if you can read your Mac's memory:
```go
package main

import (
    "fmt"
    "github.com/shirou/gopsutil/v3/mem"
)

func main() {
    v, _ := mem.VirtualMemory()
    fmt.Printf("Your Mac is using %.2f%% of its RAM!\n", v.UsedPercent)
}
```

</details>

---

<details>
<summary><b>🧠 Session 2: The Stats Engine (Data Structs)</b></summary>

### Goal
Create a "Data Package" (Struct) that holds all your system info at once.

We will define a `Stats` struct. This is like a container for our CPU, RAM, and Disk numbers. Note the **JSON Tags**—they tell Go how to talk to a website!

```go
type Stats struct {
    CPU  float64 `json:"cpu"`
    RAM  float64 `json:"ram"`
    Disk float64 `json:"disk"`
}

func getStats() Stats {
    // 1. Get CPU (takes 1 second to measure)
    c, _ := cpu.Percent(0, false)
    
    // 2. Get RAM
    v, _ := mem.VirtualMemory()
    
    // 3. Get Disk (looks at your main hard drive)
    d, _ := disk.Usage("/")

    return Stats{
        CPU:  c[0],
        RAM:  v.UsedPercent,
        Disk: d.UsedPercent,
    }
}
```

</details>

---

<details>
<summary><b>🌐 Session 3: The Data API (The Bridge)</b></summary>

### Goal
Turn your Go program into a "Data Provider" that a website can talk to.

We will create a `/api/stats` route. Instead of sending text, we are sending **JSON**—the universal language of apps like YouTube and TikTok.

```go
func statsHandler(w http.ResponseWriter, r *http.Request) {
    data := getStats()
    
    // Tell the browser we are sending JSON data
    w.Header().Set("Content-Type", "application/json")
    
    // Convert our Go Struct into a JSON string!
    json.NewEncoder(w).Encode(data)
}
```

**Testing it:** Run your code and visit `http://localhost:8080/api/stats`. You should see live hardware numbers appearing on your screen!

</details>

---

<details>
<summary><b>🎨 Session 4: Visual Dashboard (HTML & CSS)</b></summary>

Goal
Make the data look like a professional dashboard. We’ll use HTML for the structure and CSS for the "high-tech" look.
<details>
<summary><b>⚙️ How Go "Presents" HTML</b></summary>

The Magic of the ResponseWriter
In Go, the http.ResponseWriter (usually called w) is like a direct pipe to the user's browser. When you "write" to it, Go sends that information across the network.

Using Backticks (`) for Raw Strings
Normally, Go strings use double quotes (" "), but those are hard to use for HTML because HTML has its own quotes inside it. Instead, we use Backticks.

Backticks allow you to:

Multiple Lines: You can paste a whole website without errors.

No Escaping: You don't have to worry about quotes inside the HTML.

```Go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    // We send the whole HTML file through the 'w' pipe
    fmt.Fprint(w, `
        <!DOCTYPE html>
        <html>
            <h1>Go is serving this!</h1>
        </html>
    `)
}
```
</details>

1. The Structure (HTML)
We use a div with a class of "card" to group each stat. Inside, we have a "bar-bg" (the empty grey bar) and a "bar-fill" (the part that actually grows).

```HTML
<div class="card">
    <h3>CPU Usage</h3>
    <div class="bar-bg">
        <div id="cpu-bar" class="bar-fill"></div>
    </div>
</div>
```
2. The Style (CSS)
This makes the dashboard look like a dark-mode app. Notice the transition property—this makes the bars move smoothly rather than jumping!

```CSS
.card {
    background: #1e1e1e;
    padding: 20px;
    border-radius: 12px;
    width: 320px;
}

.bar-bg {
    background: #333;
    height: 12px;
    border-radius: 6px;
}

.bar-fill {
    background: #00ff88;
    height: 100%;
    width: 0%; /* JavaScript will change this! */
    transition: width 0.5s ease;
}
```
3. The Live Updates (JavaScript)
We use fetch to grab the JSON from Go, then we change the CSS width of the bars in real-time.

```JavaScript
async function update() {
    const res = await fetch('/api/stats');
    const data = await res.json();
    
    // Update the width to match the percentage from Go
    document.getElementById('cpu-bar').style.width = data.cpu + '%';
}
setInterval(update, 2000); // Do this every 2 seconds
```
</details>


---

### 🎓 Professional Skills Learned:
1. **Third-Party Libraries:** How to expand Go using `go get`.
2. **JSON APIs:** The industry standard for data transfer.
3. **Hardware Interfacing:** Reading real CPU and RAM data from macOS.
4. **Asynchronous Web:** Using `fetch` and `setInterval` for live updates.
