package sanitizerword

import (
	"strings"
	"unicode"
)

func SanitizerWords(s string) string {
	r := []rune(s)
	var cleanWord string
	for i := 0; i < len(r); i++ {
		if unicode.IsLetter(r[i]) {
			cleanWord = cleanWord + string(r[i])
		}
	}
	cleanWord = strings.ReplaceAll(cleanWord, " ", "")
	return strings.ToLower(cleanWord)
}
