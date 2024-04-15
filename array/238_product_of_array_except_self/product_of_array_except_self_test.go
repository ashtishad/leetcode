package lc_238_product_of_array_except_self

import (
	"slices"
	"testing"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"LC-1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"LC-2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{"LC-3", []int{2, 3, 4, 5}, []int{60, 40, 30, 24}},
		{"One Zero", []int{0, 3, 4, 5}, []int{60, 0, 0, 0}},
		{"Two Zeros", []int{0, 0, -1, -5, 4, 5}, []int{0, 0, 0, 0, 0, 0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := productExceptSelf(tc.nums)

			if !slices.Equal(got, tc.want) {
				t.Errorf("got:%d, want:%d", got, tc.want)
			}
		})
	}
}
