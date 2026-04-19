# 🧪 Session 2: The Library Lab (Mux & Embed)

In Session 1, we built a beautiful dashboard. But if we keep adding more code to one file, it will become a messy "spaghetti" of text! 

Today, we learn how to use **Libraries**—extra tools made by other engineers to make our lives easier.

---

### 🧠 Brain Exercise: The Lab Assistant
*Communication is key in the lab. Let's practice sending data to functions!*

**The Problem:** A scientist needs a function that takes two numbers and returns their **Product** (Multiplication).

**Your Task:** 
1. Create a function called `multiply(a int, b int) int`.
2. Inside, return `a * b`.
3. In `main`, call the function with `5` and `10` and print the result.

**The Challenge:** Can you name the types of the inputs and the return value correctly?

---

### Goal
Learn what "Libraries" are and use two of them to clean up our Safari Hub.

### 1. What are Libraries? (The Lego Analogy)
Imagine you are building a Lego castle. 
- **Standard Library:** These are the basic bricks that come in every Go box (like `fmt` and `net/http`). They are built by the Go team.
- **External Libraries:** These are special pieces (like a pre-made dragon or a catapult) that you "borrow" from other people. In Go, we use the `go get` command to download them.

### 2. Upgrading our "Map" (Gorilla Mux)
Standard Go is great, but a library called **Gorilla Mux** makes it easier to handle complex "routes" (different pages like `/explorer` or `/lab`).

Run this in your terminal to download it:
```bash
go get github.com/gorilla/mux
```

### 3. The "Scientific" Code Structure
We are going to move our HTML into its own file named `index.html`. This is how professional engineers do it! 

Create `index.html` and paste your HTML from Session 1 there. Then, update your `main.go`:

```go
package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gorilla/mux" // This is our new "Borrowed" library!
)

// This special command tells Go to "Pack" index.html into the app!
//go:embed index.html
var content embed.FS

func main() {
	// 1. Create a new "Router" using the Mux library
	r := mux.NewRouter()

	// 2. Tell the router what to do for the Homepage
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Read the packed HTML file
		data, _ content.ReadFile("index.html")
		fmt.Fprint(w, string(data))
	})

	// 3. New Route: The Research Lab
	r.HandleFunc("/lab", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Research Lab. Scopes are being calibrated...")
	})

	fmt.Println("🌳 Safari Hub is evolving! Running on http://localhost:8080...")
	http.ListenAndServe(":8080", r)
}
```

---

### 🚀 Why use `embed`?
The `embed` library is a "hidden superpower" of Go. It allows you to ship your Go program as a **single file**. You don't need to send the HTML file separately—it is "baked" into the Go code!

### 🎓 Task: The "New Wing" Challenge
1. Create a new route called `/camera`.
2. Make it return a message: `"📸 Live Safari Feed: Checking for Lions..."`.
3. Try to visit `http://localhost:8080/camera` in your browser.

**Pro-Tip:** If you see an error saying `content.ReadFile` isn't working, make sure `//go:embed index.html` (with no space after the slashes) is right above your `var content` line!

---

### Summary
1. **Libraries** are toolboxes that add new features to Go.
2. **Standard Libraries** are built-in; **External Libraries** are downloaded with `go get`.
3. **Mux** helps us manage navigation easily.
4. **Embed** packs our extra files into our Go program.
