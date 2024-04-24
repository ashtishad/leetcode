package array

import (
	"reflect"
	"testing"
)

func TestTopKFrequentPQ(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Basic test",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2}, // Order of elements in the result can vary
		},
		{
			name:     "More elements than k",
			nums:     []int{1, 1, 1, 2, 2, 3, 3, 4},
			k:        3,
			expected: []int{1, 2, 3}, // Order of elements in the result can vary
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := topKFrequent(tc.nums, tc.k)

			// Check if the elements are the same (order doesn't matter)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
