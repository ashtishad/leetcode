package array

import (
	"math"
	"slices"
)

// Sorting + Binary Search
// Time: O(nlogn) and Space: O(1)
func findMaxKBS(nums []int) int {
	slices.Sort(nums)

	res := math.MinInt
	for _, v := range nums {
		if v > 0 {
			break
		}

		if _, ok := slices.BinarySearch(nums, abs(v)); ok {
			res = max(res, abs(v))
		}
	}

	if res == math.MinInt {
		return -1
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
