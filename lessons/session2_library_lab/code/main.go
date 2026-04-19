package main

import (
	"embed"
	"fmt"
	"net/http"

	// 1. We imported the Gorilla Mux library (External)
	// You will need to run 'go get github.com/gorilla/mux' in your terminal!
	"github.com/gorilla/mux"
)

// 2. We use the 'embed' package to pack our HTML directly into the app.
// This means we only have to share the .exe file, and the website is inside it!
//go:embed index.html
var content embed.FS

func main() {
	// 3. Initialize the Gorilla Mux router
	r := mux.NewRouter()

	// 4. THE HOMEPAGE ROUTE
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// We read 'index.html' from our packed content
		data, err := content.ReadFile("index.html")
		if err != nil {
			http.Error(w, "Failed to read embedded HTML", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	// 5. THE LAB ROUTE
	// Notice how easy it is to add new "pages" with Mux!
	r.HandleFunc("/lab", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>🧪 Research Lab</h1><p>Calibration in progress. Check back in Session 3.</p>")
	})

	// 6. Start the server using our new router 'r'
	fmt.Println("----------------------------------------------------------")
	fmt.Println("🦁  Safari Hub Version 2.0 (Modular Edition)")
	fmt.Println("📦  Libraries Installed: Gorilla Mux, Embed")
	fmt.Println("🌍  Habitat Server: http://localhost:8080")
	fmt.Println("----------------------------------------------------------")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("Startup Error: %v\n", err)
	}
}
