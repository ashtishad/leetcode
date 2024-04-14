package array

// brute force (Accepted, n <= 3 * 10^4)
// time: O(n^2) space: O(1)
// sorted ASC, idx starts from 1, exactly one solution
// Input: nums = [2,3,4], target = 6, Output: [1,3]
func twoSumIIBF(nums []int, target int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i + 1, j + 1} // idx starts from 1 in ps
			}
		}
	}

	return []int{}
}
