package matrix

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			name: "Example 1",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			target:   3,
			expected: true,
		},
		{
			name: "Example 2",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			target:   13,
			expected: false,
		},
		{
			name:     "Single Element Matrix - Match",
			matrix:   [][]int{{1}},
			target:   1,
			expected: true,
		},
		{
			name:     "Single Element Matrix - No Match",
			matrix:   [][]int{{2}},
			target:   1,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := searchMatrix(tc.matrix, tc.target)
			if result != tc.expected {
				t.Errorf("Expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
