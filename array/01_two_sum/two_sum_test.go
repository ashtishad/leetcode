package array

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"LC-1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"LC-2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"LC-3", []int{3, 3}, 6, []int{0, 1}},
		{"Zero values.", []int{0, 4, 3, 0}, 0, []int{0, 3}},
		{"Negative values.", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{"Positive values.", []int{1, 2, 3, 4, 5}, 9, []int{3, 4}},
		{"Repeated values.", []int{2, 2, 2, 2}, 4, []int{0, 1}},
		{"Ascending values.", []int{0, 1, 2, 3, 4, 5}, 9, []int{4, 5}},
		{"Descending values.", []int{5, 4, 3, 2, 1, 0}, 9, []int{0, 1}},
		{"Large numbers.", []int{100, 200, 300, 400, 500}, 800, []int{2, 4}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got:%d want:%d", got, tc.want)
			}
		})

	}
}
