package array

import "slices"

// Time: o(nlogn) and Space: O(1)
func maxOperations(nums []int, k int) int {
	slices.Sort(nums)

	var ops int

	l, r := 0, len(nums)-1

	for l < r {
		sum := nums[l] + nums[r]

		switch {
		case sum == k:
			ops++
			l++
			r--
		case sum < k:
			l++
		case sum > k:
			r--
		}
	}

	return ops
}
