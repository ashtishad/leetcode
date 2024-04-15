package array

import "slices"

// Sorting.
// Time: O(n log n) and Space: O(1)
func containsDuplicateSorting(nums []int) bool {
	n := len(nums)

	if n < 1 {
		return false
	}

	slices.Sort(nums)

	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}
