package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ismferd/newGoApplication/pkg/hasher"
	"github.com/ismferd/newGoApplication/pkg/sorter"
)

func main() {
	r := os.Stdin
	var err error
	var ns sorter.OrganizedList
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			r, err = os.Open(os.Args[i])
			if err != nil {
				log.Fatal(err)
			}
			ns = hasher.Hasher(r)
			sorterOutput(ns)

		}
	} else {
		ns = hasher.Hasher(r)
		sorterOutput(ns)
	}

}

func sorterOutput(ns sorter.OrganizedList) {
	for i := 0; i < len(ns) && i < 100; i++ {
		fmt.Println(ns[i])
	}
}
