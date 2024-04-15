package array

// Brute force
// Time: O(n^2) and Space: O(1)
func containsDuplicateBF(nums []int) bool {
	n := len(nums)

	if n < 1 {
		return false
	}

	for i := range nums {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}
