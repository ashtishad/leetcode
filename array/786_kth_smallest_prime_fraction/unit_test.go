package array

import (
	"reflect"
	"testing"
)

func TestKthSmallestPrimeFraction(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"Basic Case", []int{1, 2, 3, 5}, 3, []int{2, 5}},
		{"Smallest Fraction", []int{1, 2}, 1, []int{1, 2}},
		{"Largest Prime", []int{1, 2, 3, 5, 7}, 5, []int{2, 5}},
		{"Repetition", []int{1, 2, 2, 3, 5}, 4, []int{2, 5}},
		{"Larger Input", []int{1, 3, 5, 7, 11, 13, 17, 19}, 8, []int{1, 5}},
		{"All Primes", []int{2, 3, 5, 7, 11}, 9, []int{2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := kthSmallestPrimeFraction(tc.nums, tc.k)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("kthSmallestPrimeFraction(%v, %d) = %v; want %v", tc.nums, tc.k, got, tc.want)
			}
		})
	}
}
