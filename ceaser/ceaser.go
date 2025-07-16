package ceaser

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encrypt(text string, shift int) string {
	for i := 0; i < len(text); i++ {
		if text[i] >= 'A' && text[i] <= 'Z' {
			text = text[:i] + string((text[i]-'A'-byte(shift)+26)%26+'A') + text[i+1:]
		} else if text[i] >= 'a' && text[i] <= 'z' {
			text = text[:i] + string((text[i]-'a'-byte(shift)+26)%26+'a') + text[i+1:]
		}
	}
	return text
}

func Decrypt(text string, shift int) string {
	for i := 0; i < len(text); i++ {
		if text[i] >= 'A' && text[i] <= 'Z' {
			text = text[:i] + string((text[i]-'A'-byte(shift)+26)%26+'A') + text[i+1:]
		} else if text[i] >= 'a' && text[i] <= 'z' {
			text = text[:i] + string((text[i]-'a'-byte(shift)+26)%26+'a') + text[i+1:]
		}
	}
	return text
}