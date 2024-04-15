package lc_238_product_of_array_except_self

// Approach 4: Prefix and Suffix Product in-place calculation
// I don't need separate array to store prefix and suffix products, storing directly into result array instead.
// Time: O(n) and Space: O(1)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// calculate prefix products in-place (going forward)
	currProd := 1
	for i := 0; i < n; i++ {
		res[i] = currProd
		currProd *= nums[i]
	}

	// calculate suffix products in-place (going backward)
	currProd = 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= currProd
		currProd *= nums[i]
	}

	return res
}
