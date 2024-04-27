package array

import (
	"testing"
)

func TestTotalFruit(t *testing.T) {
	testCases := []struct {
		name   string
		fruits []int
		want   int
	}{
		{
			name:   "Basic Case 1",
			fruits: []int{1, 2, 1},
			want:   3,
		},
		{
			name:   "Basic Case 2",
			fruits: []int{0, 1, 2, 2},
			want:   3,
		},
		{
			name:   "Single Fruit",
			fruits: []int{1, 1, 1, 1},
			want:   4,
		},
		{
			name:   "Only Two Fruits",
			fruits: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			want:   5,
		},
		{
			name:   "Empty Array",
			fruits: []int{},
			want:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := totalFruit(tc.fruits)
			if result != tc.want {
				t.Errorf("Expected %d, got %d", tc.want, result)
			}
		})
	}
}
