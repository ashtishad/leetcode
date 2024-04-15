package array

import "testing"

func Test_MaxProfitII(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"LC-1", []int{7, 1, 5, 3, 6, 4}, 7},
		{"LC-2", []int{1, 2, 3, 4, 5}, 4},
		{"LC-3", []int{7, 6, 4, 3, 1}, 0},
		{"LC-4", []int{3, 3, 5, 0, 0, 3, 1, 4}, 8},
		{"Single Element", []int{5}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxProfitII(tc.prices)
			if got != tc.want {
				t.Errorf("got:%d, want:%d", got, tc.want)
			}
		})
	}
}
