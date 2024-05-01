package array

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"Simple rotation", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"Rotation longer than length", []int{1, 2, 3}, 4, []int{3, 1, 2}},
		{"No rotation", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"Rotate by array length", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{"Single element", []int{1}, 10, []int{1}},
		{"Negative rotation", []int{1, 2, 3}, -3, []int{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rotate(tc.nums, tc.k)
			if !reflect.DeepEqual(tc.nums, tc.want) {
				t.Errorf("rotate(%v, %d) got %v, want %v", tc.nums, tc.k, tc.nums, tc.want)
			}
		})
	}
}
