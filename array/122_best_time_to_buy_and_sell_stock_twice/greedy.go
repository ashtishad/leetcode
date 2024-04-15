package array

// Greedy
// Time O(n) and Space: O(1)
func maxProfitII(prices []int) int {
	var cp, mp int // current and max profits

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			cp = prices[i] - prices[i-1]
			mp += cp
		}
	}

	return mp
}
