# Lesson 1: The AI Pair Programmer

**Duration:** 1.5 Hours
**Objective:** Learn how to use AI tools effectively to read, write, and debug Go code.
**Recommended Reading:** [Prompt Engineering for Developers: A Practical Guide](https://dev.to/dimagi_sihilel_0d6234fd02/prompt-engineering-for-developers-a-practical-guide-269f)

## 1. Concept: Context is Everything (30 mins)
* How LLMs process instructions and why they need specific constraints.
* The difference between a "Zero-Shot" prompt and a highly structured prompt.

**Example Prompts:**
*   **Bad (Zero-Shot):** "Write a Go function for a game." *(The AI will guess the genre, mechanics, and style, usually giving something you don't need.)*
*   **Good (Structured):** "Act as an expert Go developer. I am building a 2D game called coin-collector. Write a function that takes an integer `score` and returns a string rank (Bronze, Silver, Gold). Use a switch statement."

## 2. Activity: The Explainer (30 mins)
* Take a complex piece of code from the `coin-collector` repo (like the main game loop or collision detection).
* Feed it into an LLM and ask it to explain the code.

**Example Prompts:**
*   "Explain this Go code line-by-line to a beginner. Use analogies related to video games."
*   "What does the `&` symbol do in this specific function, and why did we use a pointer here instead of copying the value?"

## 3. Activity: Break and Fix (30 mins)
* Intentionally break a working Go function in the game (e.g., cause an index out of bounds or a syntax error).
* Copy the broken code and the terminal error output into the LLM.

**Example Prompts:**
*   "I am getting a 'panic: runtime error: index out of range' when my character hits a wall. Here is my collision function and the error trace. How do I fix it?"