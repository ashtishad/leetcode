package array

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"LC-1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"Increasing Prices", []int{7, 6, 4, 3, 1}, 0},
		{"Single Input", []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitBF(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
