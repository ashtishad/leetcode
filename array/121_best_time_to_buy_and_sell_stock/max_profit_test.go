package array

import (
	"testing"
)

func Test_MaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"LC-1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"LC-2", []int{7, 6, 4, 3, 1}, 0},
		{"Single Element", []int{5}, 0},
		{"Increasing prices", []int{1, 2, 3, 4, 5}, 4},
		{"No profit, decreasing prices", []int{8, 7, 6, 5, 4, 3, 2, 1}, 0},
		{"No profit, constant prices", []int{2, 2, 2, 2, 2}, 0},
		{"Buy on day 2, sell on day 6", []int{5, 1, 8, 6, 2, 10}, 9},
		{"Buy on day 4, sell on day 7", []int{3, 3, 5, 0, 0, 3, 1, 4}, 4},
		{"Buy on day 0, sell on day 7", []int{1, 2, 3, 4, 5, 6, 7, 8}, 7},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxProfit(tc.prices)
			if got != tc.want {
				t.Errorf("got:%d, want:%d", got, tc.want)
			}
		})
	}
}
