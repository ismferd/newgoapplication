package hasher

import (
	"bufio"
	"io"
	"log"
	"strings"

	"github.com/ismferd/newGoApplication/pkg/sanitizerwords"
	"github.com/ismferd/newGoApplication/pkg/sorter"
)

// Hasher is the manager of the whole program it recieve a io reader and return a sortedlist
func Hasher(r io.Reader) sorter.SortedList {
	i := 0
	hasher := []string{}
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
			hash = HashMakerAndScorer(joiner, hash)
			// Moving the hasher elements in order to simulate a FIFO
			hasher[0] = hasher[1]
			hasher[1] = hasher[2]
			hasher = RemoveIndex(hasher, 2)
		}
		if word != "" {
			hasher = append(hasher, word)
		}
		i++
	}
	return sorter.Sorter(hash)
}

// RemoveIndex will remove the last position of the passed array
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// HashMakerAndScorer will create a new hashmap and also it will increase the score of them
func HashMakerAndScorer(joiner string, hash map[string]int) map[string]int {
	value, isMapContainsKey := hash[joiner]
	if isMapContainsKey {
		hash[joiner] = value + 1
	} else {
		hash[joiner] = 1
	}
	return hash
}
