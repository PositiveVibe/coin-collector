# Lesson 2: Introduction to APIs in Go

**Duration:** 1.5 Hours
**Objective:** Understand how computers talk to each other over the internet before we start talking to AI.
**Recommended Reading:** [Go REST Guide. The Standard Library](https://www.jetbrains.com/guide/go/tutorials/rest_api_series/stdlib/)

## 1. Concept: The Restaurant Analogy (30 mins)
* What is an API? (The Waiter).
* What is JSON? (The Language / The Menu).
* HTTP Methods: GET vs. POST.

## 2. Activity: Fetching Public Data (30 mins)
* Use a public, free API (like the Pokémon API or a random joke API) to see how data is requested.
* Write a basic Go script using `net/http` to make a GET request.

**Example API Requests to test in the browser:**
*   `https://pokeapi.co/api/v2/pokemon/pikachu` (Look at the raw JSON data that comes back).

## 3. Activity: Parsing the JSON (30 mins)
* Create Go `structs` that match the shape of the JSON.
* Use `encoding/json` to unmarshal the response into your structs.
* Print specific fields to the terminal.

**Example JSON parsing goal:**
*   Extract just the "name" and "base_experience" from the API response and ignore the hundreds of other lines of data.