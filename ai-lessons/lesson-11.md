# Lesson 11: Retrieval-Augmented Generation (RAG)

**Duration:** 1.5 Hours
**Objective:** Give the NPC access to a "library" of game lore so it can answer questions about the game world without hallucinating.
**Recommended Reading:** [What is Retrieval-Augmented Generation (RAG)?](https://research.ibm.com/blog/retrieval-augmented-generation-RAG)

## 1. Concept: The Open-Book Test (30 mins)
* What is RAG? (Explaining that LLMs are like students taking a closed-book test; RAG gives them an open book).
* Why we can't just put the entire game history into the system prompt (Context limits and API cost).

## 2. Activity: Building the Knowledge Base (30 mins)
* Create a `lore.txt` file in the Go project containing facts about the `coin-collector` universe.
* Write a Go function that takes a keyword and does a simple search to find the most relevant sentences in the lore file.

**Example `lore.txt` content:**
```text
The Magic Sword was forged in Mount Doom by the ancient goblins.
The Bronze Coin is worth 1 point and is the most common currency.
The Golden Crown belonged to the Mad King and is hidden in the Whispering Woods.
```
