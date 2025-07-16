package ceaser

import (
	"fmt"
	"strings"
)

func Encrypt(text string, shift int) string {
	var result strings.Builder
	shift = shift % 26
	fmt.Printf("Encrypting with shift: %d\n", shift)

	for _, c := range text {
		switch {
		case c >= 'A' && c <= 'Z':
			result.WriteByte(byte((int(c-'A')+shift)%26 + 'A'))
		case c >= 'a' && c <= 'z':
			result.WriteByte(byte((int(c-'a')+shift)%26 + 'a'))
		default:
			result.WriteByte(byte(c))
		}
	}

	return result.String()
}

func Decrypt(text string, shift int) string {
	return Encrypt(text, 26 - (shift % 26))
}
