package hasher

import (
	"bufio"
	"io"
	"log"
	"strings"

	"github.com/ismferd/newGoApplication/pkg/sanitizerwords"
	"github.com/ismferd/newGoApplication/pkg/sorter"
)

func Hasher(r io.Reader) sorter.OrganizedList {
	hasher := []string{}
	i := 0
	hash := map[string]int{}
	scanner := bufio.NewScanner(r)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		word = sanitizerwords.SanitizerWords(word)
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
	return sorter.Organizer(hash)
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
