package sorter

import (
	"reflect"
	"testing"
)

func TestOrganizer(t *testing.T) {
	var expectedList = []Sorted{
		{
			Key:   " lorem ipsum is",
			Value: 4,
		}, {
			Key:   " of the same",
			Value: 3,
		}, {
			Key:   " lorem ipsum good",
			Value: 0,
		},
	}
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want SortedList
	}{
		{"OrganizerTest",
			args{map[string]int{
				" of the same":      3,
				" lorem ipsum is":   4,
				" lorem ipsum good": 0,
			}}, expectedList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sorter(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Organizer() = %v, want %v", got, tt.want)
			}
		})
	}
}
