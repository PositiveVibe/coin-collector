package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SearchLore scans a text file for lines containing a specific keyword
func SearchLore(keyword string) string {
	file, err := os.Open("lore.txt")
	if err != nil {
		return "Lore library is currently closed."
	}
	defer file.Close()

	var results []string
	scanner := bufio.NewScanner(file)

	// Read file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Case-insensitive search
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			results = append(results, line)
		}
	}

	if len(results) == 0 {
		return "No relevant lore found."
	}

	// Combine all found sentences into one block of context
	return strings.Join(results, " ")
}

func BuildNPCPrompt(userQuestion string, searchKeyword string) string {
	// 1. Retrieve the context (The "R" in RAG)
	loreContext := SearchLore(searchKeyword)

	// 2. Augment the prompt (The "A" in RAG)
	systemPrompt := fmt.Sprintf(`You are the Wisdom Statue. 
Use the following game lore to answer the player's question. 
If the answer is not in the lore, say "The ancient texts are silent on this matter."
Do not make up facts.

GAME LORE:
%s

PLAYER ASKS:
%s`, loreContext, userQuestion)

	// 3. (Generation) Send this augmented string to the LLM...
	return systemPrompt
}

func main() {
	// Example: A player asks about the Golden Crown
	prompt := BuildNPCPrompt("Where is the Golden Crown?", "crown")
	fmt.Println("=== Generated Prompt for LLM ===")
	fmt.Println(prompt)
}
