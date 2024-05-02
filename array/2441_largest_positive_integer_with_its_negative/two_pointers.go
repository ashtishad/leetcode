package array

import "slices"

// Sorting + Two Pointers
// Time: O(nlogn) and Space: O(1)
func findMaxK(nums []int) int {
	slices.Sort(nums)

	l, r := 0, len(nums)-1

	for l < r {
		switch {
		case nums[l]+nums[r] > 0:
			r--
		case nums[l]+nums[r] < 0:
			l++
		default:
			return nums[r]
		}
	}

	return -1
}
