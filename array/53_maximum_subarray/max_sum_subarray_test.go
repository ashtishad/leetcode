package array

import "testing"

// Test cases for the MaxSubArray function
func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"LC-1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"LC-2", []int{5, 4, -1, 7, 8}, 23},
		{"One Positive, rest negatives", []int{-1, -2, -3, -4, 6}, 6},
		{"Single element", []int{5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayBF(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
