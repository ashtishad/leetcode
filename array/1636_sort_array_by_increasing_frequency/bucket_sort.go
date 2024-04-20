package array

import "slices"

// Approach: Bucket Sort
// Time: O(n) and Space: O(n)
func frequencySort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	count := make(map[int]int)

	for _, v := range nums {
		count[v]++
	}

	buckets := make([][]int, len(nums)+1)
	for v, f := range count {
		buckets[f] = append(buckets[f], v)
	}

	// buckets: [[] [3] [1] [2] [] [] []]
	// freq:     0   1   2   3

	var res []int
	for f := 1; f < len(buckets); f++ {
		vals := buckets[f]
		slices.SortFunc(vals, descendingCmp)

		for _, v := range vals {
			for i := 0; i < f; i++ {
				res = append(res, v)
			}
		}
	}

	return res
}

func descendingCmp(a, b int) int {
	if a > b {
		return -1
	}

	if a < b {
		return 1
	}

	return 0
}
