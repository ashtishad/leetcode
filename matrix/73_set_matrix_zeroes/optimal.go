package matrix

// Time: O(m*n) and Space: O(1)
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	rowZero := false // track first row has zero or not

	for _, v := range matrix[0] {
		if v == 0 {
			rowZero = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// handle all rows, except the first row
	for i := 1; i < m; i++ {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// handle first row
	if rowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}
