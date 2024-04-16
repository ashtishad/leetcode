package array

// Greedy
// time: O(n) and Space: O(1)
// [2,3,1,1,4] true, [3,2,1,0,4] false
// each idx represents maximum jump from that position.
func canJump(nums []int) bool {
	if len(nums) != 1 && nums[0] == 0 {
		return false
	}

	var reachable int
	for i, jumps := range nums {
		if i > reachable {
			return false
		}

		reachable = max(reachable, i+jumps)
	}

	return true
}
