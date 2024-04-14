package array

// Two Pointers.
// time: O(n) space: O(1)
// out order matters, sorted ASC, idx starts from 1, exactly one solution
// Input: nums = [2,3,4,5,7,9], target = 6, Output: [1,3]
func twoSumII(nums []int, target int) []int {
	n := len(nums)
	l, r := 0, n-1

	for r < n {
		sum := nums[l] + nums[r]

		switch {
		case sum < target:
			l++
		case sum > target:
			r--
		case target == sum:
			return []int{l + 1, r + 1}
		}
	}

	return []int{}
}
