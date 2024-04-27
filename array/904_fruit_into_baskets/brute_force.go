package array

// Brute Force (TLE)
// Time: O(n^2) and Space: O(1)
func totalFruitBF(fruits []int) int {
	maxLen := 0

	for i := 0; i < len(fruits); i++ {
		baskets := make(map[int]bool)
		total := 0

		for j := i; j < len(fruits); j++ {
			baskets[fruits[j]] = true

			// If we need a third basket, break
			if len(baskets) > 2 {
				break
			}

			total++
			maxLen = max(maxLen, total)
		}
	}

	return maxLen
}
