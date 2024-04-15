package array

import "testing"

// appears more than ⌊n / 2⌋ times
func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"TC-1", []int{3, 3, 2, 3, 4, 3, 2, 4, 3, 4, 4, 4, 4, 7, 3, 4, 4, 4, 4, 4}, 4},
		{"TC-2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"TC-3", []int{7}, 7},
		{"TC-4", []int{2, 8, 7, 2, 2, 5, 2, 3, 1, 2, 2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			if got != tt.want {
				t.Errorf("got:%d, want:%d", got, tt.want)
			}
		})
	}
}
