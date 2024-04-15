package array

// Hash Table
// Time: O(n) and Space: O(n)
func containsDuplicate(nums []int) bool {
	n := len(nums)

	if n < 1 {
		return false
	}

	seen := make(map[int]bool, n)

	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}

		// mark as seen
		seen[num] = true
	}

	return false
}
