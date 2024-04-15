package array

import "math"

// Problem statement from EPI Book
// Time:O(n) and Space: O(n)
func buyAndSellStockTwice(prices []int) int {
	profit := 0
	firstProfit := make([]int, 0)
	minPrice := math.MaxInt

	// Forward phase
	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		profit = max(profit, prices[i]-minPrice)
		firstProfit = append(firstProfit, profit)
	}

	// Backward phase
	maxPriceSoFar := -math.MaxInt
	for i := len(prices) - 1; i > 0; i-- {
		maxPriceSoFar = max(maxPriceSoFar, prices[i])
		profit = max(profit, maxPriceSoFar-prices[i]+firstProfit[i-1])
	}

	return profit
}
