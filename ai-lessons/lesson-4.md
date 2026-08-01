# Lesson 4: Memory and Context

**Duration:** 1.5 Hours
**Objective:** Build a CLI chatbot that remembers what you said by managing an array of messages.
**Recommended Reading:** [Understanding LLM Context Windows](https://www.pinecone.io/learn/llm-context-window/)

## 1. Concept: AI Has No Memory (30 mins)
* APIs are stateless. To have a conversation, we must send the *entire* chat history every single time.
* Discuss token limits (the "backpack" of memory the AI can carry).

## 2. Activity: Setting up the CLI Loop (30 mins)
* Use `bufio.Reader` and `os.Stdin` to create an infinite loop that waits for user input.
* Create a Go slice of type `[]Message` to hold the history.

## 3. Activity: The Chat Loop (30 mins)
* Append the user's input to the `[]Message` slice. Send the whole slice to the API. Append the AI's response to the slice.

**Example Prompt Sequence to test memory:**
*   *User Input 1:* "Hi, I am a knight in the coin-collector game. My favorite color is green."
*   *AI Response 1:* "Greetings, brave green knight!"
*   *User Input 2:* "What was my favorite color again?"
*   *AI Response 2 (if code works):* "Your favorite color is green!" *(If it fails, it means the slice isn't updating correctly!)*