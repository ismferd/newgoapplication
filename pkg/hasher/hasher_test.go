package hasher

import (
	"os"
	"reflect"
	"testing"

	"github.com/ismferd/newGoApplication/pkg/sorter"
)

func TestHasher(t *testing.T) {

	expected := sorter.SortedList{
		sorter.Sorted{Key: "bar bar bar", Value: 3},
		sorter.Sorted{Key: "bar bar foo", Value: 1},
	}
	fileName := "test.txt"

	r, _ := os.Open(fileName)

	m := Hasher(r)
	eq := reflect.DeepEqual(m, expected)
	if !eq {
		t.Errorf("Maps hasn't the expected values")
	}
}

func TestHashMakerAndScorer(t *testing.T) {
	hash := map[string]int{}
	test := []string{"foo foo foo", "foo foo foo", "bar bar bar"}
	for _, value := range test {
		hash = HashMakerAndScorer(value, hash)
	}
	expected := map[string]int{"foo foo foo": 2, "bar bar bar": 1}

	eq := reflect.DeepEqual(hash, expected)
	if !eq {
		t.Errorf("Array hasn't the expected values")
	}
}

func TestRemoveIndex(t *testing.T) {
	hasher := []string{"foo", "bar", "foobar"}
	hasher = RemoveIndex(hasher, 2)

	expected := []string{"foo", "bar"}
	eq := reflect.DeepEqual(hasher, expected)
	if !eq {
		t.Errorf("Array hasn't the expected values")
	}
}
