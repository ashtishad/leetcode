package array

import (
	"slices"
	"testing"
)

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{"LC-1", []int{2, 0, 2, 1, 1, 0}},
		{"LC-2", []int{2, 0, 1}},
		{"LC-3", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)

			if !slices.IsSorted(tt.nums) {
				t.Errorf("got:%d", tt.nums)
			}
		})
	}
}
