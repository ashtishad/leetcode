package array

// brute force
// time: O(n^2) and space: O(1)
func maxProfitBF(prices []int) int {
	n := len(prices)
	var cp, mp int

	for i, buy := range prices[:n-1] { // up to 2nd last day
		for _, sell := range prices[i+1:] {
			// if profitable, calc profit and max profit
			if sell > buy {
				cp = sell - buy
				mp = max(mp, cp)
			}
		}
	}

	return mp
}
