package array

// Divide and Conquer
// time: O(n) and space: O(1)
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	sum, maxSum := 0, nums[0]

	for _, v := range nums {
		sum += v

		if sum < v {
			sum = v
		}

		maxSum = max(maxSum, sum)
	}

	return maxSum
}
