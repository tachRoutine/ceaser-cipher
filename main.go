package main

import (
	"flag"
	"fmt"
	"practice/ceaser-cipher/ceaser"
)

func main() {
	shift := flag.Int("shift", 1, "number of shifts of the Caesar cipher")
	content := flag.String("content", "Hello World!", "content to be encrypted")
	operationType := flag.String("type", "Encrypt", "type of operation: Encrypt or Decrypt")
	fmt.Println("Number of substitutions:", *shift)
	fmt.Println("Content to be encrypted:", *content)

	switch *operationType {
	case "Encrypt":
		fmt.Println("Encrypted content:", ceaser.Encrypt(*content, *shift))
	case "Decrypt":
		fmt.Println("Decrypted content:", ceaser.Decrypt(*content, *shift))
	default:
		fmt.Println("Invalid type. Use 'Encrypt' or 'Decrypt'.")
	}
	flag.Parse()
}
