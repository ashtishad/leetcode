package array

// Time: O(n) and Space: O(1)
func rotate(nums []int, k int) {
	n := len(nums)

	k %= n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, s, e int) {
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
}
