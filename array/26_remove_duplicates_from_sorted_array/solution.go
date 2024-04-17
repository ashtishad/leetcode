package array

// Optimal, same can be achieved by slices.Compact()
// time: O(n) and space: O(1)
func removeDuplicates(nums []int) int {
	n := len(nums)

	if n < 2 {
		return n
	}

	k := 1 // number of unique elems
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
