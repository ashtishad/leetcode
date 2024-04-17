package array

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int // num of unique elements
	}{
		{"TC-1", []int{0, 0, 1, 1, 1, 2, 2, 3, 4}, 5},
		{"TC-2", []int{1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 2},
		{"One Number", []int{1}, 1},
		{"All duplicates", []int{2, 2, 2, 2, 2, 2, 2, 2, 2}, 1},
		{"Nothing duplicate", []int{1, 2, 3, 4, 5, 6}, 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := removeDuplicates(tc.nums)
			if res != tc.k {
				t.Errorf("got:%d, want:%d", res, tc.k)
			}
		})
	}
}
