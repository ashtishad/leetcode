package array

// Approach: Brute Force
// Time: O(n^2) Space: O(1)
func majorityElementBF(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	var cnt int
	for i := range nums {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				cnt++
			}

			if cnt > n/2 {
				return nums[i]
			}
		}
	}

	return -1
}
