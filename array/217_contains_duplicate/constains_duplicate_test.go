package array

import (
	"testing"
)

// Given an integer array nums, return true if any value appears "at least twice" in the array,
// and return false if every element is distinct.
func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"TC-1", []int{1, 2, 3, 1}, true},
		{"TC-2", []int{1, 2, 3, 4}, false},
		{"TC-3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{"Empty array", []int{}, false},           // Edge case
		{"Single elem", []int{1}, false},          // Edge case
		{"Duplicate Elements", []int{1, 1}, true}, // Edge case
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := containsDuplicate(tc.nums)
			if got != tc.want {
				t.Errorf("got:%v, want:%v", got, tc.want)
			}
		})
	}
}
