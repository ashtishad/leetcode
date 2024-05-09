package array

import (
	"testing"
)

func TestMaximizeHappinessSum(t *testing.T) {
	testCases := []struct {
		name      string
		happiness []int
		k         int
		want      int64
	}{
		{"Basic case", []int{1, 2, 3}, 2, 4},
		{"All children", []int{1, 1, 1, 1}, 4, 1},
		{"One child", []int{3, 7, 9, 5}, 1, 9},
		{"Zero children", []int{6, 5, 4, 3, 2, 1}, 0, 0},
		{"Large numbers", []int{1000, 2000, 3000, 4000, 5000}, 2, 8999},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maximumHappinessSum(tc.happiness, tc.k)
			if got != tc.want {
				t.Errorf("maximizeHappiness(%v, %d) = %d, want %d", tc.happiness, tc.k, got, tc.want)
			}
		})
	}
}
