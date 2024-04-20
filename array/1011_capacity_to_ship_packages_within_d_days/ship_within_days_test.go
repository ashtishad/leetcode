package array

import (
	"testing"
)

func TestShipWithinDays(t *testing.T) {
	type test struct {
		name    string
		weights []int
		days    int
		want    int
	}

	tests := []test{
		{"Basic Functionality", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
		{"Single Day Shipping", []int{5, 2, 8, 3}, 1, 18},
		{"Heaviest Item Exceeds Capacity", []int{3, 2, 2, 4, 1, 4}, 3, 6},
		{"All Items in One Day", []int{1, 1, 1}, 1, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := shipWithinDays(tc.weights, tc.days)
			if got != tc.want {
				t.Errorf("shipWithinDays(%v, %d) = %d, expected %d", tc.weights, tc.days, got, tc.want)
			}
		})
	}
}
