package array

import "testing"

func TestSearchInRotatedSortedArray(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"Single element - found", []int{1}, 1, 0},
		{"Single element - not found", []int{2}, 1, -1},
		{"Array not rotated", []int{1, 2, 3, 4, 5}, 3, 2},
		{"Target at pivot", []int{4, 5, 1, 2, 3}, 1, 2},
		{"Target not present", []int{6, 7, 8, 1, 2, 3}, 9, -1},
		{"Edge 1", []int{3, 1}, 3, 0},
		{"Edge 2", []int{5, 1, 3}, 3, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := search(tc.nums, tc.target)
			if got != tc.want {
				t.Errorf("search(%v, %d) = %d; want %d", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}
