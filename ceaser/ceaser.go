package ceaser

import "strings"

func Encrypt(text string, shift int) string {
	var result strings.Builder
	shift = shift % 26

	for _, c := range text {
		switch {
		case c >= 'A' && c <= 'Z':
			result.WriteByte(byte((int(c-'A')+shift)%26 + 'A'))
		case c >= 'a' && c <= 'z':
			result.WriteByte(byte((int(c-'a')+shift)%26 + 'a'))
		default:
			result.WriteByte(byte(c)) // leave unchanged
		}
	}

	return result.String()
}

func Decrypt(text string, shift int) string {
	// Decryption is simply encryption with negative shift
	return Encrypt(text, 26 - (shift % 26))
}
