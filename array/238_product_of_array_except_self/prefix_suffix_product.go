package lc_238_product_of_array_except_self

// Approach 3: Prefix And Suffix Product Array
// Time: O(n) and Space: O(n)
func productExceptSelfPrefixSuffixProduct(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// Prefix Product[i] = pre[i - 1] * a[i - 1]
	pre := make([]int, n)
	pre[0] = 1
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}

	// Suffix/Postfix Product[i] = suff[i + 1] * a[i + 1]
	suff := make([]int, n)
	suff[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suff[i] = suff[i+1] * nums[i+1]
	}

	// res[i] = pre[i] * suff[i]
	for i := 0; i < n; i++ {
		res[i] = pre[i] * suff[i]
	}

	return res
}
