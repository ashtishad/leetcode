package array

// Kadane's Algo
// Time:O(n) and Space: O(1)
// prices = [7,1,5,3,6,4], want:5
func maxProfitKadane(prices []int) int {
	var profit int

	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		profit = max(profit, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}

	return profit
}

// // Short Syntax
// func maxProfitKadaneSS(prices []int) int {
// 	var profit int
//
// 	minPrice := prices[0]
// 	for _, currPrice := range prices[1:] {
// 		profit = max(profit, currPrice-minPrice)
// 		minPrice = min(minPrice, currPrice)
// 	}
//
// 	return profit
// }
