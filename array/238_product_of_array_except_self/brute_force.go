package lc_238_product_of_array_except_self

// Brute Force (TLE in Leetcode)
// Time: O(n^2) and Space: O(n)
func productExceptSelfBF(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		prod := 1
		for j := 0; j < n; j++ {
			if j != i {
				prod *= nums[j]
			}
		}

		res[i] = prod
	}

	return res
}
