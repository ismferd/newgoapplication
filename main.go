package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/ismferd/newGoApplication/pkg/sanitizerwords"
)

func main() {
	if len(os.Args) >= 1 {
		for i := 1; i < len(os.Args); i++ {
			r := os.Args[1]
			content, err := ioutil.ReadFile(r) // the file is inside the local directory
			if err != nil {
				fmt.Println("Err")
			}

			Hasher(string(content))
		}
	}
}

func Hasher(s string) {
	words := strings.Fields(s)
	hasher := []string{}
	i := 0
	hash2 := map[string]int{}
	for _, word := range words {
		word := sanitizerwords.SanitizerWords(word)
		if len(hasher) == 3 {
			joiner := strings.Join(hasher, " ")
			value, isMapContainsKey := hash2[joiner]
			if isMapContainsKey {
				hash2[joiner] = value + 1
			} else {
				hash2[joiner] = 1
			}

			hasher[0] = hasher[1]
			hasher[1] = hasher[2]
			hasher = RemoveIndex(hasher, 2)
		}
		hasher = append(hasher, word)
		i++
	}
	type Organized struct {
		Key   string
		Value int
	}
	ns := make([]Organized, 0)
	for k, v := range hash2 {
		ns = append(ns, Organized{k, v})
	}

	// Then sorting the slice by value, higher first.
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].Value > ns[j].Value
	})
	for i := 0; i < len(ns) && i < 100; i++ {
		fmt.Println(ns[i])
	}
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
