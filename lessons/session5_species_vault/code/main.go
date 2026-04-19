package main

import (
	"bufio"
	"fmt"
	"os"
)

// encryptDecrypt takes a string and a single byte key to scramble/unscramble data.
// Since it uses XOR, the same function works for both!
func encryptDecrypt(input string, key byte) string {
	data := []byte(input)

	for i := 0; i < len(data); i++ {
		// XOR (^) flips the bits of the data based on the key
		data[i] = data[i] ^ key
	}

	return string(data)
}

func main() {
	// A special key used to "Lock" the data. 
	// This must be a number between 1 and 255.
	var secretKey byte = 137 

	fmt.Println("==========================================")
	fmt.Println("🔐  DIGITAL SAFARI: SPECIES VAULT")
	fmt.Println("==========================================")

	// We use a Reader to allow the student to type in a whole sentence
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter secret animal location to protect: ")
	text, _ := reader.ReadString('\n')

	fmt.Println("\n[1] SECURING DATA...")
	scrambled := encryptDecrypt(text, secretKey)
	fmt.Printf("🕵️  SCRAMBLED: %s\n", scrambled)

	fmt.Println("\n[2] READY FOR STORAGE.")
	fmt.Println("Wait... do you want to see the original? (Type 'yes')")
	var choice string
	fmt.Scanln(&choice)

	if choice == "yes" {
		unscrambled := encryptDecrypt(scrambled, secretKey)
		fmt.Printf("🔓  RESTORED: %s", unscrambled)
	} else {
		fmt.Println("Keeping the vault locked. Good choice.")
	}

	fmt.Println("==========================================")
}
