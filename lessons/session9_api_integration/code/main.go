package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// CatFact represents the structure of the data returned by the API.
// Check it out in your browser: https://catfact.ninja/fact
type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// fetchFact reaches out to the internet to get a new discovery.
func fetchFact(client *http.Client) {
	url := "https://catfact.ninja/fact"

	// 1. Send the Request
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("⚠️  CONNECTION ERROR: %v\n", err)
		return
	}
	
	// 'defer' ensures we close the connection even if the code fails later
	defer resp.Body.Close()

	// 2. Read the binary data from the response stream
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("⚠️  READ ERROR: %v\n", err)
		return
	}

	// 3. Unpack (Unmarshal) the JSON into our Go struct
	var myFact CatFact
	err = json.Unmarshal(body, &myFact)
	if err != nil {
		fmt.Printf("⚠️  DECODE ERROR: %v\n", err)
		return
	}

	// 4. Print the discovery
	fmt.Println("----------------------------------------------------------")
	fmt.Printf("🕒 Time: %s\n", time.Now().Format("15:04:05"))
	fmt.Printf("📖 RESEARCH NOTE: %s\n", myFact.Fact)
	fmt.Printf("📊 Data Size: %d characters\n", myFact.Length)
}

func main() {
	fmt.Println("==========================================")
	fmt.Println("🌐  SAFARI ENCYCLOPEDIA: WEB API MODE")
	fmt.Println("==========================================")
	fmt.Println("Connecting to the Global Research Database...")

	// We create a custom 'Client' with a timeout. 
	// This ensures our program doesn't hang forever if the internet is slow!
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Fetch 5 facts, waiting 5 seconds between each
	for i := 1; i <= 5; i++ {
		fmt.Printf("\nDiscovery #%d in progress...", i)
		fetchFact(client)
		
		if i < 5 {
			time.Sleep(5 * time.Second)
		}
	}

	fmt.Println("\n==========================================")
	fmt.Println("✅  Research Session Complete.")
	fmt.Println("==========================================")
}
