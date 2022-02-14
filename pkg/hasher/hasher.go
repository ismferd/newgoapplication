package hasher

import (
	"fmt"
	"strings"

	"github.com/ismferd/newGoApplication/pkg/sanitizerwords"
	"github.com/ismferd/newGoApplication/pkg/sorter"
)

func Hasher(s string) {
	words := strings.Fields(s)
	hasher := []string{}
	i := 0
	hash := map[string]int{}
	for _, word := range words {
		word := sanitizerwords.SanitizerWords(word)
		if len(hasher) == 3 {
			joiner := strings.Join(hasher, " ")
			value, isMapContainsKey := hash[joiner]
			if isMapContainsKey {
				hash[joiner] = value + 1
			} else {
				hash[joiner] = 1
			}
			hasher[0] = hasher[1]
			hasher[1] = hasher[2]
			hasher = RemoveIndex(hasher, 2)
		}
		hasher = append(hasher, word)
		i++
	}
	ns := sorter.Organizer(hash)

	for i := 0; i < len(ns) && i < 100; i++ {
		fmt.Println(ns[i])
	}
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
