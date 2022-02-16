package hasher

import (
	"os"
	"reflect"
	"testing"

	"github.com/ismferd/newGoApplication/pkg/sorter"
)

func TestScanGroupWords(t *testing.T) {
	var (
		testOL = sorter.OrganizedList{
			sorter.Organized{Key: "bar bar bar", Value: 3},
			sorter.Organized{Key: "bar bar foo", Value: 1},
		}
		fileName = "test.txt"
	)
	r, _ := os.Open(fileName)

	m := Hasher(r)
	eq := reflect.DeepEqual(m, testOL)
	if !eq {
		t.Errorf("Maps hasn't the expected values")
	}
}
