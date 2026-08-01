# Lesson 7: Function Calling (Theory & Schema)

**Duration:** 1.5 Hours
**Objective:** Teach the AI how to use tools to get real-time information about our game state.
**Recommended Reading:** [Function Calling (OpenAI API Docs)](https://developers.openai.com/api/docs/guides/function-calling)

## 1. Concept: Reversing the Flow (30 mins)
* Normally, we ask the AI a question, and it answers.
* With function calling, the AI realizes it needs data and asks *us* to run a function.

## 2. Activity: Defining a Tool Schema (30 mins)
* Look at the required JSON structure for defining a "Tool" for the API.
* Write a JSON schema for a function called `get_player_coins()`.

**Example System Prompt to introduce the tool:**
*   "You are a helpful game assistant. You have access to a tool called `get_player_coins`. If the user asks about their money or inventory, use this tool to find the exact amount."

## 3. Activity: Catching the Tool Call (30 mins)
* Send the payload to the API.

**Example User Prompt to trigger the tool:**
*   "Hey, how many coins do I have right now?"
*   *Expected output:* Instead of text, the API should return a `tool_call` request asking for `get_player_coins`.