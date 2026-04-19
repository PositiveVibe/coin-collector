# 🌐 Session 9: The Digital Encyclopedia (Web APIs)

Science is about collaboration! In the digital world, apps collaborate by sharing data through **APIs** (Application Programming Interfaces). 

Today, we will turn our Safari Hub into a "Digital Encyclopedia" by fetching real-world animal facts from a website on the other side of the planet.

---

### 🧠 Brain Exercise: The Weighing Station
*API data often involves decimals! Let's make sure we're comfortable with floats.*

**The Problem:** You have two baby elephants. One weighs `120.5` kg and the other weighs `115.8` kg.

**Your Task:** 
1. Create two `float64` variables for the weights.
2. Add them together and store them in a variable called `totalWeight`.
3. Print: `"The two elephants weigh a total of: [number] kg"`.
4. **The "Strict" Challenge:** Try to add a weight (float64) to an age (int). Watch the Go compiler stop you! How do you fix it?

---

### Goal
Build a Go tool that talks to a real Web API and displays an animal fact.

### 1. What is an API?
Think of an API as a **Menu** at a restaurant.
1. You (The Client) look at the menu.
2. You ask the waiter for a specific dish (**Request**).
3. The kitchen (The API Server) prepares the dish.
4. The waiter brings it to your table (**Response**).

### 2. The JSON Handshake
Most APIs talk in JSON (which we learned in Session 6). To read the API response, we need to create a `struct` that matches their "Menu".

### 3. The Encyclopedia Code
We will use the "Cat Facts API". Create `main.go` in your `code/` folder:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// This struct matches the data the API sends us
type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func getAnimalFact() {
	// 1. Ask the API for a fact
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Println("Error: Couln't reach the encyclopedia!")
		return
	}
	defer resp.Body.Close()

	// 2. Unpack the JSON response
	var fact CatFact
	json.NewDecoder(resp.Body).Decode(&fact)

	// 3. Display it!
	fmt.Println("\n📖 NEW RESEARCH DISCOVERY:")
	fmt.Printf("\"%s\"\n", fact.Fact)
}

func main() {
	fmt.Println("🌐 Connecting to Great Library of Animalia...")
	
	// Create a loop that fetches a new fact every 5 seconds
	for i := 0; i < 3; i++ {
		getAnimalFact()
		time.Sleep(5 * time.Second)
	}
}
```

---

### 🚀 Research Task
1. Run the code. You should see a random cat fact appear every 5 seconds!
2. **The "URL" Discovery:** Copy the link `https://catfact.ninja/fact` and paste it into your web browser. Do you see the raw JSON data?
3. **Question:** Why do we use `json.NewDecoder` here instead of `json.Unmarshal`? (Hint: The data is coming through a "Stream" from the internet!)

### 🎓 Task: The "Species Expansion" Challenge
- There are many free APIs! Can you find one for Dog facts or Zoo animals?
- Try creating a new `struct` and a new URL to fetch a different kind of animal data.
- **Pro-Tip:** Most professional apps (like Uber, Netflix, and Pokémon GO) are mostly just many APIs talking to each other!

---

### Summary
1. **APIs** let your Go program talk to other computers over the internet.
2. **http.Get** is how we ask for data from a website.
3. **json.NewDecoder** is the best way to read JSON data that is being "streamed" from a server.
