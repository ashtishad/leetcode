package array

import "slices"

// Sorting, then return element of index n/2
// Time: O(nlogn) Space: O(1)
func majorityElementSorting(nums []int) int {
	slices.Sort(nums)

	return nums[len(nums)/2]
}
