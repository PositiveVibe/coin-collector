# Lesson 5: Structured AI Output (JSON)

**Duration:** 1.5 Hours
**Objective:** Force the AI to respond in strict JSON so our Go code can actually use the data programmatically.
**Recommended Reading:** [A practical guide to OpenAI JSON Mode](https://www.eesel.ai/blog/openai-json-mode)

## 1. Concept: Text vs. Data (30 mins)
* If an AI returns a paragraph of text, our game can't use it easily.
* If an AI returns JSON, we can use it to instantly spawn enemies, change stats, or trigger events.

## 2. Activity: Prompting for JSON (30 mins)
* Update our API call to include instructions forcing JSON output.
* Use the provider's "JSON Mode" if available.

**Example Prompt:**
*   "Generate a random enemy for my game. You MUST respond in valid JSON matching this exact structure: `{"name": "string", "health": integer, "element": "fire|water|earth"}`. Do not include any greeting, markdown formatting, or other text."

## 3. Activity: Unmarshaling AI Data (30 mins)
* Create a Go struct with `Name`, `Health`, and `Element` fields.
* Run the AI response through `json.Unmarshal`.
* Prove it worked by printing: `fmt.Printf("A wild %s appeared with %d health!\n", enemy.Name, enemy.Health)`