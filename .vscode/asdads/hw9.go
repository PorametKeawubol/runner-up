package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func rot13(text string) string {
	var result strings.Builder
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			result.WriteRune((char-'A'+13)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			result.WriteRune((char-'a'+13)%26 + 'a')
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func maidn() {
	var option int
	fmt.Println("Select 2 options")
	fmt.Println(" - 1 encrypt with ROT 13")
	fmt.Println(" - 2 decrypt with ROT 13")
	fmt.Println()
	fmt.Print("Choose option: ")
	fmt.Scanln(&option)

	fmt.Print("Enter text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()

	if option == 1 {
		ciphertext := rot13(message)
		fmt.Printf("Ciphertext is \"%s\"\n", ciphertext)
	} else if option == 2 {
		plaintext := rot13(message)
		fmt.Printf("Plaintext is \"%s\"\n", plaintext)
	}

}
