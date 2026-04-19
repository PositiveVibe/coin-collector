package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SafariManager handles our animal location data safely
type SafariManager struct {
	// A map to store Animal Name -> Sector
	locations map[string]string
	
	// A Mutex to prevent "Race Conditions"
	// This is like a lock on a diary so only one person writes at a time.
	mu sync.Mutex
}

// UpdateLocation safely changes an animal's sector
func (sm *SafariManager) UpdateLocation(animal string, sector string) {
	// 1. LOCK the data
	sm.mu.Lock()
	
	// 2. Schedule the UNLOCK to happen automatically when we are done
	defer sm.mu.Unlock()

	// 3. Make the change
	sm.locations[animal] = sector
}

// GetReport returns a copy of the current sightings
func (sm *SafariManager) GetReport() map[string]string {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// We create a copy so the dashboard can read it without holding the lock too long
	reportCopy := make(map[string]string)
	for k, v := range sm.locations {
		reportCopy[k] = v
	}
	return reportCopy
}

func scout(sm *SafariManager, name string) {
	for {
		// Pick a random sector 1-50
		sector := fmt.Sprintf("Sector-%d", rand.Intn(50)+1)
		
		sm.UpdateLocation(name, sector)

		// Wait 1-3 seconds before moving again
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	}
}

func main() {
	manager := &SafariManager{
		locations: make(map[string]string),
	}

	fmt.Println("🛰️  INITIALIZING GLOBAL TRACKING SYSTEM...")
	fmt.Println("------------------------------------------")

	// Start four concurrent "Scouts" tracking different animals
	animals := []string{"Elephant 🐘", "Lion 🦁", "Gazelle 🦌", "Rhino 🦏"}
	for _, a := range animals {
		go scout(manager, a)
	}

	// THE DASHBOARD: Show updates every second
	for {
		// Clear screen command (works on most terminals)
		fmt.Print("\033[H\033[2J") 

		fmt.Println("==========================================")
		fmt.Println("🏁  LIVE SAFARI GPS DASHBOARD")
		fmt.Println("==========================================")
		fmt.Println("ANIMAL       |  LAST SEEN LOCATION")
		fmt.Println("------------------------------------------")

		report := manager.GetReport()
		if len(report) == 0 {
			fmt.Println("Waiting for scout reports...")
		}

		for animal, loc := range report {
			fmt.Printf("%-12s |  %s\n", animal, loc)
		}

		fmt.Println("------------------------------------------")
		fmt.Println("Press Ctrl+C to exit monitoring mode.")
		
		time.Sleep(1 * time.Second)
	}
}
