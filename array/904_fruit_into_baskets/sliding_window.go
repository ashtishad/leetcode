package array

// Sliding Window
// Time: O(n) and Space: O(1)
// longest contigous subarray holding 2 type of fruits
func totalFruit(fruits []int) int {
	var res, l int
	baskets := make(map[int]int)

	for r, f := range fruits {
		baskets[f]++

		for len(baskets) > 2 {
			baskets[fruits[l]]--

			if baskets[fruits[l]] == 0 {
				delete(baskets, fruits[l])
			}

			l++
		}

		res = max(res, r-l+1)
	}

	return res
}
