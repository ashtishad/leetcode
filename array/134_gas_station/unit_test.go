package array

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	testCases := []struct {
		name string
		gas  []int
		cost []int
		want int
	}{
		{"Basic case", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{"Single station", []int{2}, []int{2}, 0},
		{"Not possible", []int{2, 3, 4}, []int{3, 4, 5}, -1},
		{"Complete cycle possible from start", []int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}, 4},
		{"Multiple possible starts", []int{5, 5, 1, 3, 4}, []int{4, 4, 5, 1, 1}, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := canCompleteCircuit(tc.gas, tc.cost)
			if got != tc.want {
				t.Errorf("canCompleteCircuit(%v, %v) = %d, want %d", tc.gas, tc.cost, got, tc.want)
			}
		})
	}
}
