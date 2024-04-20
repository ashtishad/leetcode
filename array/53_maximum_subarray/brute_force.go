package array

import (
	"math"
)

// brute force
// time: O(n^2) and space: O(1)
func maxSubArrayBF(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	sum, maxSum := 0, math.MinInt

	for i := range nums {
		sum = nums[i]
		maxSum = max(maxSum, sum)

		for j := i + 1; j < n; j++ {
			sum += nums[j]
			maxSum = max(maxSum, sum)
		}
	}

	return maxSum
}
