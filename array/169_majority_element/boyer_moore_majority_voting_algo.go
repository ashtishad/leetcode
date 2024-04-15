package array

// Boyer-Moore Majority Voting Algorithm
// Time: O(n) and Space: O(1)
func majorityElementX(nums []int) int {
	cnt, candidate := 0, nums[0]

	for _, v := range nums[1:] {
		switch {
		case v == candidate:
			cnt++
		case cnt == 0:
			candidate = v
		default:
			cnt--
		}
	}

	return candidate
}
