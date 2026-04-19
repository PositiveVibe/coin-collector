# 🔐 Session 5: The Species Vault (Basic Encryption)

In a real safari, researchers have to keep the locations of endangered animals secret so that poachers can't find them. In the digital world, we do this with **Encryption**.

Encryption is the art of turning a message into "Digital Gibberish" that only people with a **Key** can read.

---

### 🧠 Brain Exercise: The Digitizer
*Before we scramble data, we need to turn letters into numbers!*

**The Problem:** Computers don't see the letter "A", they see the number `65`.

**Your Task:** 
1. Create a `string` variable named `secret := "ABC"`.
2. Convert it into a `[]byte` slice like this: `data := []byte(secret)`.
3. Print `data`. What numbers do you see in the terminal? (Hint: These are the ASCII codes for A, B, and C).

---

### Goal
Build an Encryption Tool that scrambles animal data so hackers can't understand it.

### 1. The XOR Cipher (The "Secret Switch")
We will use a professional trick called **XOR** (Exclusive OR). 
- Imagine every letter is a series of `0`s and `1`s.
- XOR takes your letter and a "Key" and flips the `1`s and `0`s. 
- **The cool part:** If you XOR it again with the same key, it flips back to the original message!

### 2. Bytes and Strings
In Go, letters are really just numbers (bytes). To scramble them, we talk to the computer in bytes.

### 3. The Scrambler Code
Create `main.go` in your `code/` folder:

```go
package main

import (
	"fmt"
)

// The XOR function: This is the "Engine" of our vault.
func encryptDecrypt(input string, key byte) string {
	// 1. Convert the string into a slice of bytes (numbers)
	data := []byte(input)

	// 2. Loop through every single character
	for i := 0; i < len(data); i++ {
		// 3. Flip the bits using the XOR operator: ^
		data[i] = data[i] ^ key
	}

	// 4. Convert it back into a readable string
	return string(data)
}

func main() {
	secretMessage := "Tiger located at 45.2N, 12.3E"
	var secretKey byte = 42 // This is our "Locker Key" (any number 1-255)

	fmt.Println("🔒 Original Message:", secretMessage)

	// ENCRYPT it!
	scrambled := encryptDecrypt(secretMessage, secretKey)
	fmt.Println("🕵️ Scrambled Data:", scrambled)

	// DECRYPT it! (Running it again with the same key fixes it)
	unscrambled := encryptDecrypt(scrambled, secretKey)
	fmt.Println("🔓 Unscrambled:", unscrambled)
}
```

---

### 🚀 Research Task
1. Run the code. You should see the message turn into weird symbols.
2. **The "Secret Key" Test:** Change the `secretKey` to a different number (like `99`) but try to unscramble it with the old key. Does it work?
3. **Question:** Why is it safer to store location data this way instead of plain text?

### 🎓 Task: The "Double Lock" Challenge
- Can you change the code so it asks the user to type in the message using `fmt.Scanln`?
- **Pro-Tip:** If you lose your `secretKey`, the message is lost forever! That's why professional researchers always keep backups.

---

### Summary
1. **Encryption** turns data into gibberish to keep it safe.
2. **XOR (^)** is a mathematical way to flip data back and forth.
3. In Go, we use `[]byte` to manipulate individual characters of a string.
