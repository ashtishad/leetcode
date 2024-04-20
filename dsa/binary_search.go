package dsa

// Approach 1: Find the Exact Value
func binarySearch(x []int, target int) int {
	l, r := 0, len(x)-1

	for l <= r {
		m := l + ((r - l) >> 1) // floor value

		switch {
		case x[m] == target:
			return m
		case x[m] < target:
			l = m + 1
		case x[m] > target:
			r = m - 1
		}
	}

	return -1
}

// Approach 2: Find Insert Position, check slices.BinarySearch()
func binarySearchX(nums []int, target int) (int, bool) {
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
