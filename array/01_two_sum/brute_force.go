package array

// Brute Force
// time O(n^2) and space O(1)
// Input: nums = [2,7,11,15], target = 9, Output: [0,1]
func twoSumBF(nums []int, target int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j} // only one answer exists
			}
		}
	}

	return []int{}
}
