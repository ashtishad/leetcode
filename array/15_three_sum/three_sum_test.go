package array

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Basic Case",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "No Triplets",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "All Zeros",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "With Duplicates",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "Large Input",
			nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			want: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := threeSum(tc.nums)

			for _, res := range got {
				slices.Sort(res)
			}
			sort.Slice(got, func(i, j int) bool {
				return reflect.DeepEqual(got[i], got[j])
			})

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
