package array

import (
	"reflect"
	"testing"
)

func TestSortArrayByIncreasingFrequency(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"LC-1", []int{1, 1, 2, 2, 2, 3}, []int{3, 1, 1, 2, 2, 2}},
		{"LC-2", []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
		{"LC-3", []int{2, 3, 1, 3, 2}, []int{1, 3, 3, 2, 2}},
		{"LC-4", []int{-6, -6, -6}, []int{-6, -6, -6}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := frequencySort(tc.nums)

			if len(got) != len(tc.want) {
				t.Errorf("lenghts are not equal, got:%d, want:%d", got, tc.nums)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got:%d, want:%d", got, tc.want)
			}

			// // when output order is not mandatory
			// seen := make(map[int]int)
			//
			// for _, v := range got {
			// 	seen[v]++
			// }
			//
			// for _, v := range tc.want {
			// 	if f, ok := seen[v]; !ok || f < 0 {
			// 		t.Errorf("got:%d, want:%d", got, tc.want)
			// 	}
			//
			// 	seen[v]--
			// }
		})
	}
}
