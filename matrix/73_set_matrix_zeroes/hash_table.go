package matrix

// Time: O(m*n) and Space: O(m+n)
func setZeroesHT(matrix [][]int) {
	rows := make(map[int]bool)
	cols := make(map[int]bool)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for r := range rows {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[r][j] = 0
		}
	}

	for c := range cols {
		for j := 0; j < len(matrix); j++ {
			matrix[j][c] = 0
		}
	}
}
