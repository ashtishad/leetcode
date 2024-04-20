package array

import "math"

// Kadane
// time: O(n) and space: O(1)
func maxSubArrayKadane(nums []int) int {
	sum, maxSum := 0, math.MinInt32

	for _, v := range nums {
		sum = max(v, sum+v)
		maxSum = max(maxSum, sum)
	}

	return maxSum
}
