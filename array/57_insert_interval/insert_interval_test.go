package array

import (
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			name:        "Empty List",
			intervals:   [][]int{},
			newInterval: []int{2, 5},
			want:        [][]int{{2, 5}},
		},
		{
			name:        "Before All",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{0, 1},
			want:        [][]int{{0, 3}, {6, 9}},
		},
		{
			name:        "After All",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{10, 12},
			want:        [][]int{{1, 3}, {6, 9}, {10, 12}},
		},
		{
			name:        "Multiple Overlaps",
			intervals:   [][]int{{1, 3}, {6, 8}},
			newInterval: []int{4, 9},
			want:        [][]int{{1, 3}, {4, 9}},
		},
		{
			name:        "Complete Containment",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 3},
			want:        [][]int{{1, 5}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := insertX(tc.intervals, tc.newInterval)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
