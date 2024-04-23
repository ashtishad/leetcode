package main

// Bucket Sort
// Time: O(n) and Space: O(n)
// nums {1, 1, 1, 2, 2, 3} k=2 want: [1,2]
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}

	// map[1:3 2:2 3:1]
	buckets := make([][]int, len(nums)+1) // extra one for zero based indexing
	for v, f := range count {
		buckets[f] = append(buckets[f], v)
	}

	// buckets : [[] [3] [2] [1] [] [] []] // v=num, idx=freq
	// freq    :   0  1   2   3   4  5  6
	var res []int

	for i := len(buckets) - 1; i > 0; i-- {
		for _, v := range buckets[i] {
			res = append(res, v)

			if len(res) == k {
				return res // returning the actual answer
			}
		}
	}

	return []int{}
}
