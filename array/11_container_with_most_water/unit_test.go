package array

import "testing"

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		name   string
		height []int
		want   int
	}{
		{"Simple case", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"All the same height", []int{1, 1, 1, 1, 1}, 4},
		{"Two elements only", []int{1, 1}, 1},
		{"Increasing heights", []int{1, 2, 3, 4, 5}, 6},
		{"Decreasing heights", []int{5, 4, 3, 2, 1}, 6},
		{"Empty array", []int{}, 0},
		{"Single element", []int{5}, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxArea(tc.height)
			if got != tc.want {
				t.Errorf("maxArea(%v) = %d, want %d", tc.height, got, tc.want)
			}
		})
	}
}
