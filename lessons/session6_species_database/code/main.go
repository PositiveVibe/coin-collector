package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Animal represents an animal record in our scientific database.
// The text inside `json:"..."` tells Go how to name the field in the JSON file.
type Animal struct {
	Name     string `json:"name"`
	Species  string `json:"species"`
	Location string `json:"location"`
	SeenAt   string `json:"seen_at"`
	Danger   int    `json:"danger_level"` // 1 (Safe) to 5 (Extremely Dangerous)
}

func main() {
	filename := "safari_database.json"

	fmt.Println("==========================================")
	fmt.Println("📁  SPECIES DATABASE: FILE PERSISTENCE")
	fmt.Println("==========================================")

	// 1. Let's create some initial research data
	researchLog := []Animal{
		{
			Name:     "Leo",
			Species:  "Lion",
			Location: "Waterhole Alpha",
			SeenAt:   "08:45 AM",
			Danger:   4,
		},
		{
			Name:     "Spot",
			Species:  "Giraffe",
			Location: "Acacia Grove",
			SeenAt:   "10:20 AM",
			Danger:   1,
		},
	}

	// 2. Turn our Go data into JSON (Marshalling)
	// MarshalIndent makes it pretty-printed so humans can read it easily.
	jsonData, err := json.MarshalIndent(researchLog, "", "    ")
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}

	// 3. Save it to our hard drive
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
	fmt.Printf("✅ SUCCESS: Research log saved to %s\n", filename)

	// -------------------------------------------------------------------------
	
	// 4. Now, let's READ it back from the disk!
	fmt.Println("\n[READING DATABASE FROM DISK...]")
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 5. Convert JSON text back into our Go []Animal slice (Unmarshalling)
	var loadedLog []Animal
	err = json.Unmarshal(fileBytes, &loadedLog)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return
	}

	// 6. Display the results
	for i, a := range loadedLog {
		fmt.Printf("%d. [%s] %s the %s was spotted at %s\n", i+1, a.SeenAt, a.Name, a.Species, a.Location)
		if a.Danger >= 4 {
			fmt.Println("   ⚠️ WARNING: HIGH DANGER! KEEP DISTANCE.")
		}
	}

	fmt.Println("==========================================")
}
