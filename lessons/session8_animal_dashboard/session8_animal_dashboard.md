# 🚦 Session 8: The Live Dashboard (Mutexes & Safety)

In Session 3 and 4, we learned how to run many tasks at once. But there is a danger! What happens if two scouts try to write in the same scientific journal at the exact same millisecond?

The computer gets confused and **Crashes**. This is called a "Race Condition". Today, we learn how to use a **Mutex** (The "Talking Stick") to keep our data safe.

---

### 🧠 Brain Exercise: The Danger Detector
*Safety first! Let's practice using logic to flag dangerous situations.*

**The Problem:** An animal tracker sends a `dangerLevel` (int) from 1 to 10. If the level is `7` or higher, we need to sound an alarm.

**Your Task:** 
1. Create a variable `dangerLevel := 8`.
2. Write an `if` statement to check if it is `7` or higher.
3. If it is, print `"⚠️ ALARM: High danger detected!"`.
4. If it's lower, print `"✅ Status: Safe"`.

---

### Goal
Build a dashboard that tracks the locations of multiple animals simultaneously without crashing.

### 1. The Mutex (Mutual Exclusion)
Think of a Mutex like a **Locker Key**.
1. To write to the map, a scout must "Lock" the mutex.
2. If another scout tries to write, they must **wait** until the key is returned.
3. Once finished, the scout "Unlocks" it so the next person can use it.

### 2. The Tracking Code
Create `main.go` in your `code/` folder:

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Our shared database
var (
	// The map that holds 'Animal Name' -> 'Current Location'
	locations = make(map[string]string)
	
	// THE MUTEX: This is our digital "Talking Stick"
	mu sync.Mutex
)

func trackAnimal(name string) {
	for {
		newLoc := fmt.Sprintf("Sector %d", rand.Intn(100))

		// 1. LOCK the journal so no one else can write
		mu.Lock()
		locations[name] = newLoc
		// 2. UNLOCK it so others can use it
		mu.Unlock()

		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func main() {
	// Start 3 separate scouts at once!
	go trackAnimal("Lion 🦁")
	go trackAnimal("Zebra 🦓")
	go trackAnimal("Elephant 🐘")

	// The High-Speed Dashboard
	for {
		fmt.Println("\033[H\033[2J") // This clears the terminal screen
		fmt.Println("🛰️ SAFARI GPS DASHBOARD")
		fmt.Println("---------------------------")

		// 3. LOCK before reading the map too!
		mu.Lock()
		for animal, loc := range locations {
			fmt.Printf("%-12s | %s\n", animal, loc)
		}
		mu.Unlock()

		time.Sleep(1 * time.Second)
	}
}
```

---

### 🚀 Research Task
1. Run the code. You should see a live table updating in your terminal!
2. **The "Crash" Test:** Comment out `mu.Lock()` and `mu.Unlock()` in both places. Run it again. Depending on your computer speed, it might eventually crash with a "Concurrent Map Writes" error! 
3. **Question:** Why do we need to lock even when we are just **Reading** the data?

### 🎓 Task: The "Fast & Furious" Challenge
- Add 10 more animals to the tracker.
- Make the `time.Sleep` in `trackAnimal` much shorter (e.g., `500 * time.Millisecond`).
- Watch how hard the Mutex works to keep the data safe!

---

### Summary
1. **Race Conditions** happen when two goroutines touch the same data at once.
2. **sync.Mutex** ensures only one goroutine can access data at a time.
3. Always remember to `.Unlock()`! If you forget, your program will "Freeze" forever (this is called a Deadlock).
