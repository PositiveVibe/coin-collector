# 📁 Session 6: The Species Database (Keeping Records)

In Session 5, we scrambled our data. But as soon as we turned off the computer, the data was gone! Professional researchers use **Databases** to save their work forever.

Today, we will use **JSON files** to store our species list. JSON is the language of the internet—it's how modern apps talk to each other.

---

### 🧠 Brain Exercise: The Species Blueprint
*Before we save data to files, we need a template for our animals.*

**The Problem:** You want to describe a **Bird**. It has a `Name` (string), `Wingspan` (float64), and `CanFly` (bool).

**Your Task:** 
1. Create a `type Bird struct` with those 3 fields.
2. In `main`, create a new bird named "Falcon" with a wingspan of `1.5` and `CanFly: true`.
3. Print the whole bird using `fmt.Printf("%+v\n", myBird)`.

---

### Goal
Build a Go program that can save animal data to a file and read it back later.

### 1. The "Record" Template (Structs)
We need a standard way to describe an animal. In Go, we use a `struct` with **JSON Tags** to tell Go how to name the data in the file.

```go
type Animal struct {
	Name     string `json:"name"`
	Species  string `json:"species"`
	Location string `json:"location"`
	Count    int    `json:"count"`
}
```

### 2. Saving to File (Marshalling)
- **Marshalling:** Turning a Go Struct into a JSON string.
- **Unmarshalling:** Turning a JSON string back into a Go Struct.

### 3. The Database Code
Create `main.go` in your `code/` folder:

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Animal struct {
	Name     string `json:"name"`
	Species  string `json:"species"`
	Location string `json:"location"`
	Count    int    `json:"count"`
}

func main() {
	// 1. Create a list of animals (our database)
	database := []Animal{
		{Name: "Simba", Species: "Lion", Location: "Pride Rock", Count: 1},
		{Name: "Balu", Species: "Bear", Location: "Jungle", Count: 5},
	}

	// 2. Convert (Marshal) our list into JSON
	// We use 'Indent' to make the file easy for humans to read!
	jsonData, _ := json.MarshalIndent(database, "", "  ")

	// 3. Save it to a file named 'species.json'
	os.WriteFile("species.json", jsonData, 0644)
	fmt.Println("💾 Database saved to species.json!")

	// 4. READ it back!
	fileData, _ := os.ReadFile("species.json")
	var loadedData []Animal
	json.Unmarshal(fileData, &loadedData)

	fmt.Println("\n🔎 Loading records from file...")
	for _, a := range loadedData {
		fmt.Printf("- %s the %s was seen at %s\n", a.Name, a.Species, a.Location)
	}
}
```

---

### 🚀 Research Task
1. Run the code. A new file named `species.json` should appear in your folder!
2. Open `species.json`. Can you change the `Count` of an animal manually in the file?
3. **The "Persistent" Test:** Delete everything inside `main()` *except* the code that reads the file. Run it again. Does Go still remember your animals? (Yes! Because they are on the disk now!)

### 🎓 Task: The "New Arrival" Challenge
- Can you add a new animal to the list in your code (e.g., a "Rhino") and save it?
- **Pro-Tip:** In the real world, JSON is used by Spotify to send you your playlists and by Instagram to send you pictures. You're learning the "Industry Standard"!

---

### Summary
1. **JSON** is a format for storing and sending data.
2. **json.Marshal** turns Go code into JSON text.
3. **os.WriteFile** and **os.ReadFile** let Go talk to your hard drive.
