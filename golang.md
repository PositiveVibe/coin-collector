
# 🚀 Project: Build Your Own URL Shortener in Go

Welcome! Over the next few weeks, we are moving away from "gaming" and diving into **Systems Engineering**. We are using **Go** (or Golang), a language created by Google to power the world's biggest servers. 

Our goal is to build a tool that takes a long website address and turns it into a short code—exactly like Bitly or TinyURL.

---

<details>
<summary><b>📦 Session 1: The "Go" Foundations (Hello, Compiler!)</b></summary>

### Goal
Get your environment ready and see why Go is different from Python or JavaScript.

### 1. Project Setup
Open your terminal and run these commands to create your workspace:
```bash
mkdir go-shortener
cd go-shortener
go mod init shortener

```

### 2. Your First Struct

In Go, we use `structs` to create templates for data. It's much stricter than a Python dictionary! Create a file named `main.go`:

```go
package main

import "fmt"

// This is our template for a Link
type Link struct {
    ID      string
    LongURL string
    Clicks  int
}

func main() {
    // Create a new link
    myLink := Link{
        ID:      "goog",
        LongURL: "[https://google.com](https://google.com)",
        Clicks:  0,
    }

    // Pro-Tip: Use %+v to see the labels (ID, LongURL, etc.) in the terminal
    fmt.Printf("Link Created: %+v\n", myLink)
}

```

**The "Strict" Challenge:** Try to change `Clicks` to a string like `"ten"`. Watch the compiler stop you! This "Strictness" is why Go is so fast and reliable.

</details>

---

<details>
<summary><b>🌐 Session 2: Turning your PC into a Server</b></summary>

### Goal

Learn how the internet "talks" using HTTP.

### 1. The Basic Listener

Replace your `main.go` code with this to start a web server. This makes your computer "listen" for visitors.

```go
package main

import (
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    // w = ResponseWriter (what we send back to the user)
    // r = Request (what the user sent to us)
    fmt.Fprintf(w, "<h1>Welcome to the URL Vault</h1>")
}

func main() {
    http.HandleFunc("/", homeHandler)
    fmt.Println("Server starting on http://localhost:8080...")
    http.ListenAndServe(":8080", nil)
}

```

**Task:** Run `go run main.go` and visit `http://localhost:8080` in your browser. You just hosted your first website!

</details>

---

<details>
<summary><b>🧠 Session 3: The "Brain" (The Redirect Logic)</b></summary>

### Goal

Use a "Map" to store links and make the browser "jump" to a new site automatically.

### 1. The URL Map

We need a place to store our shortcuts. In Go, a `map` stores keys and values.

```go
// This is like a dictionary: short code -> long URL
var urlStore = map[string]string{
    "google": "[https://google.com](https://google.com)",
    "wiki":   "[https://wikipedia.org](https://wikipedia.org)",
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    // r.URL.Path[1:] grabs the text after the slash (e.g., "google")
    shortCode := r.URL.Path[1:]

    targetURL, exists := urlStore[shortCode]

    if exists {
        // This is the "Magic" that moves the browser to the new site!
        http.Redirect(w, r, targetURL, http.StatusFound)
    } else {
        fmt.Fprintf(w, "Error: Code '%s' not found.", shortCode)
    }
}

```

**Task:** In your `main()` function, add `http.HandleFunc("/", redirectHandler)`. Now, going to `localhost:8080/google` should instantly take you to Google!

</details>

---

<details>
<summary><b>🛠️ Session 4: Building the "Shorten" Tool (The API)</b></summary>

### Goal

Let people create their own short links without ever touching the code.

### 1. The Shorten Handler

We will use **URL Parameters** to send new data to our server.

```go
func shortenHandler(w http.ResponseWriter, r *http.Request) {
    // Example usage: /shorten?id=yt&url=[https://youtube.com](https://youtube.com)
    newID := r.URL.Query().Get("id")
    newURL := r.URL.Query().Get("url")

    if newID != "" && newURL != "" {
        urlStore[newID] = newURL
        fmt.Fprintf(w, "Success! Created link: localhost:8080/%s", newID)
    } else {
        fmt.Fprintf(w, "Error: You need an 'id' and a 'url'!")
    }
}

```

### 2. Final Order

In your `main()`, make sure `/shorten` is registered **above** the redirect handler:

```go
func main() {
    http.HandleFunc("/shorten", shortenHandler)
    http.HandleFunc("/", redirectHandler)
    http.ListenAndServe(":8080", nil)
}

```

**Final Boss Challenge:** Create a short link for your favorite website using the URL parameters in your browser, then try to visit it!

</details>

---

### 🎓 What you learned:

1. **Go Syntax:** How to use Structs, Maps, and strict data types.
2. **Web Servers:** How to turn a computer into a server using `net/http`.
3. **The Internet's Language:** How Redirects, Status Codes, and Query Parameters work.

---

