package array

// Two Pointers, Same Direction
// Time:O(n) and Space: O(1)
// prices = [7,1,5,3,6,4], want:5
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}

	var cp, mp, l int // current, max profits and left pointer.

	for r := 0; r < n; r++ {
		// if profitable? calc profit, max profit and shift pointers.
		if prices[r] > prices[l] {
			cp = prices[r] - prices[l]
			mp = max(mp, cp)
			continue
		}

		// if we reach here, means non profitable, encountered a low price, so set left pointer to it.
		l = r
	}

	return mp
}
