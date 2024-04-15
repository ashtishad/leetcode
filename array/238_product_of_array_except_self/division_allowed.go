package lc_238_product_of_array_except_self

// Division Operator Allowed (NOT Allowed In Leetcode)
// Time: O(n) and Space: O(n)
func productExceptSelfDivisionAllowed(nums []int) []int {
	res := make([]int, len(nums))
	prod, zeroCount := 1, 0

	for _, v := range nums {
		if v == 0 {
			zeroCount++
			continue
		}

		prod *= v
	}

	for i := 0; i < len(nums); i++ {
		switch {
		case zeroCount > 1:
			return res
		case zeroCount == 1:
			if nums[i] == 0 {
				res[i] = prod
			} else {
				res[i] = 0
			}
		default:
			res[i] = prod / nums[i]
		}
	}

	return res
}
