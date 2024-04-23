package array

import "math"

// Sliding Window
// Time: O(n) and Space: O(1)
func findMaxAverage(nums []int, k int) float64 {
	res := -math.MaxFloat64

	var l, sum int

	// [1,12,-5,-6,50,3], k=4
	for r, v := range nums {
		sum += v

		if r-l+1 == k {
			avg := float64(sum) / float64(k)
			res = max(res, avg)
			sum -= nums[l]
			l++
		}
	}

	return res
}
