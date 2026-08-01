# Lesson 6: Concurrency with Goroutines

**Duration:** 1.5 Hours
**Objective:** Learn how to make slow network calls without freezing the main application.
**Recommended Reading:** [Go Concurrency: The Full Guide to Goroutines and Beyond](https://dev.to/pixperk/go-concurrency-the-full-guide-to-goroutines-and-beyond-2ajh)

## 1. Concept: Blocking vs. Asynchronous (30 mins)
* API calls take 1-3 seconds. If we do this in the `coin-collector` game loop, the entire game will freeze.
* Introduction to Goroutines (`go doSomething()`) and Go Channels (the "walkie-talkie" pattern).

## 2. Activity: Simulating a Game Loop (30 mins)
* Write a simple infinite loop that prints "Game Running..." every 500ms.
* Add an API call to the loop without a goroutine. Watch the "game" freeze.

**Example Prompt to intentionally cause a delay:**
*   "Write a highly detailed, 10-paragraph story about the ancient history of the first coin ever minted in the universe." *(This will take the AI a long time to generate).*

## 3. Activity: Non-Blocking Calls (30 mins)
* Wrap the API call in a Goroutine.
* Use a Go Channel (`chan string`) to pass the AI's response back to the main thread once it's ready.
* Watch the "game" continue to run smoothly while the AI thinks in the background.