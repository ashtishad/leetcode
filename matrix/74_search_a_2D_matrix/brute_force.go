package matrix

// Time: O(log(m * n)) Space: O(1)
func searchMatrixBF(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + ((right - left) >> 1)
		midVal := matrix[mid/n][mid%n]

		switch {
		case midVal == target:
			return true
		case midVal < target:
			left = mid + 1
		case midVal > target:
			right = mid - 1
		}
	}

	return false
}
