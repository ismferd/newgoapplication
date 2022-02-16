package sanitizerwords

import (
	"strings"
	"unicode"
)

func SanitizerWords(s string) string {
	r := []rune(s)
	var sanitizeWord string
	for i := 0; i < len(r); i++ {
		if unicode.IsLetter(r[i]) {
			sanitizeWord = sanitizeWord + string(r[i])
		}
	}
	sanitizeWord = strings.ReplaceAll(sanitizeWord, " ", "")
	return strings.ToLower(sanitizeWord)
}
