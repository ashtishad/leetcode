package array

import (
	"slices"
)

// Sorting + Greedy
// Time: O(nlogn) and Space: O(1)
func maximumHappinessSum(x []int, k int) int64 {
	n := len(x)

	slices.Sort(x)

	var res, decr int

	for i := n - 1; i >= n-k; i-- {
		v := x[i] - decr
		if v > -1 {
			res += v
		}

		decr++
	}

	return int64(res)
}
