package matrix

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "Basic Matrix",
			input:    [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			expected: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name:     "Matrix with Multiple Zeros",
			input:    [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			expected: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		{
			name:     "Single Row Matrix",
			input:    [][]int{{0, 1, 2}},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "Single Column Matrix",
			input:    [][]int{{1}, {0}, {2}},
			expected: [][]int{{0}, {0}, {0}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			setZeroes(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected: %v, got: %v", tc.expected, tc.input)
			}
		})
	}
}
