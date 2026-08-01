# Lesson 9: The NPC Integration

**Duration:** 1.5 Hours
**Objective:** Bring everything together into the actual `coin-collector` codebase by creating an interactive NPC.
**Recommended Reading:** [Understanding LLMs for Game NPCs: An exploration of opportunities](https://www.diva-portal.org/smash/get/diva2:1938971/FULLTEXT01.pdf)

## 1. Concept: System Design (30 mins)
* Where does the NPC code live in our repo structure?
* How does the player trigger the NPC? (e.g., walking up to a "Wisdom Statue" and pressing the 'E' key).

## 2. Activity: Game Loop Integration (45 mins)
* Add the trigger logic to the game loop.
* Use the Goroutines and Channels from Lesson 6 to fire off the API call when the player interacts.
* Pass the player's current actual score to the AI using the Function Calling logic from Lesson 8.

**Example Game-Context Prompts:**
*   *Player Action:* Presses 'E' near the statue.
*   *Background Code:* Sends prompt: "The player just approached you. Give them a cryptic hint about exploring the map."

## 3. Activity: UI Updates (15 mins)
* When the channel receives the AI's response, render it onto the screen using the game engine's text drawing functions.