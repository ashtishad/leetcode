package array

import (
	"reflect"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"LC-1", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"LC-2", []int{0, 0, 3, 4}, 0, []int{1, 2}},
		{"LC-3", []int{1, 3, 4, 4}, 8, []int{3, 4}},
		{"LC-4", []int{2, 3, 4}, 6, []int{1, 3}},
		{"LC-5", []int{-1, 0}, -1, []int{1, 2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := twoSumII(tc.nums, tc.target)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got:%d want:%d", got, tc.want)
			}
		})
	}
}
