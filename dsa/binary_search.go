package dsa

// binarySearch searches for target in a sorted slice and returns the position
// where target is found, or the position where target would appear in the
// sort order; it also returns a bool saying whether the target is really found
// in the slice. The slice must be sorted in increasing order.
// Input: nums = [-1,0,3,5,9,12], target = 13, out: 4
func binarySearch(nums []int, target int) (int, bool) {
	n := len(nums)

	l, r := 0, n

	for l < r {
		m := l + ((r - l) >> 1)
		// l <= m < r
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}

	return l, l < n && nums[l] == target
}
