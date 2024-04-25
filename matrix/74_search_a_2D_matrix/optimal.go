package matrix

// Time: O(log(m * n)) Space: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	if matrix[0][0] > target || matrix[m-1][n-1] < target {
		return false
	}

	l, r := 0, m
	for l < r {
		mid := l + ((r - l) >> 1)

		v := matrix[mid][0]
		switch {
		case v == target:
			return true
		case v < target:
			l = mid + 1
		case v > target:
			r = mid
		}
	}

	row := l

	l, r = 0, n
	for l < r {
		mid := l + ((r - l) >> 1)

		v := matrix[row-1][mid]

		switch {
		case v == target:
			return true
		case v < target:
			l = mid + 1
		case v > target:
			r = mid
		}
	}

	return false
}
