# Lesson 3: The LLM Handshake

**Duration:** 1.5 Hours
**Objective:** Write a Go script that successfully sends a prompt to an LLM API and retrieves the response.
**Recommended Reading:** [Go: How to send POST HTTP requests with a JSON body](https://www.practical-go-lessons.com/post/go-how-to-send-post-http-requests-with-a-json-body-cbhvuqa220ds70kp2lkg)

## 1. Concept: Authentication & Payloads (30 mins)
* How API keys work and why we keep them secret (`.env` files).
* Understanding the JSON payload required by the LLM (model name, messages array, temperature).

## 2. Activity: Building the Request (30 mins)
* Set up a new Go file: `ask_ai.go`.
* Build a struct for the LLM request payload.
* Use `net/http` to send a POST request with the API key in the headers.

**Example Prompt to hardcode in your Go script:**
*   "Write a 3-sentence backstory for a goblin who loves shiny gold coins but is afraid of the dark. Make it funny."

## 3. Activity: Unpacking the Response (30 mins)
* Look at the complex nested JSON returned by the AI provider.
* Build the corresponding Go structs to dig out just the `content` string.
* Run the program and successfully print the AI's answer to the terminal.