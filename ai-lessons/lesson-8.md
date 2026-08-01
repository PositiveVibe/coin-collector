# Lesson 8: Function Calling (Wiring to the Game)

**Duration:** 1.5 Hours
**Objective:** Connect the AI's tool requests to the actual `coin-collector` state.
**Recommended Reading:** [Function Calling (OpenAI API Docs)](https://developers.openai.com/api/docs/guides/function-calling)

## 1. Concept: The Three-Step Handshake (30 mins)
* Step 1: User asks a question.
* Step 2: AI asks for function execution. Go runs the function and returns the data.
* Step 3: AI uses the data to answer the original question.

## 2. Activity: Executing the Tool in Go (30 mins)
* Write the Go logic to intercept the `tool_call` from Lesson 7.
* Write a mock function that checks a variable (e.g., `playerCoins = 42`) and returns it as a string to the AI.

## 3. Activity: Completing the Loop (30 mins)
* Send the function result back to the API.

**Example Prompts for the full loop:**
*   *User:* "Can I afford the Magic Sword? It costs 50 coins."
*   *AI (internal logic):* Calls `get_player_coins()`.
*   *Go Code:* Returns `42`.
*   *AI (final answer to user):* "You only have 42 coins! You need 8 more to afford the Magic Sword."