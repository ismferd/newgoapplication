package sorter

import (
	"sort"
)

//Organized Structure to add the organized values from a map
type Sorted struct {
	Key   string
	Value int
}

//OrganizedList slice of Organized structure
type SortedList []Sorted

/*
//TotalElements number of elements that the Organized structure has
func (o *SortedList) TotalElements() int {
	return len(*o)
}
*/
//Organizer receive a map and return a OrganizedList with elements in the correct order
func Organizer(m map[string]int) SortedList {
	ns := make([]Sorted, 0)
	for k, v := range m {
		ns = append(ns, Sorted{k, v})
	}

	// Then sorting the slice by value, higher first.
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].Value > ns[j].Value
	})

	return ns
}
