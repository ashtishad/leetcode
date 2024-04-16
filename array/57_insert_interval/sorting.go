package array

import "sort"

// Sorting
// Time: O(nlogn) and Space: O(n)
func insert(intervals [][]int, newInterval []int) [][]int {
	merged := append(intervals, newInterval)

	sort.Slice(merged, func(i, j int) bool {
		return merged[i][0] < merged[j][0] // asc
	})

	res := [][]int{}

	prev := merged[0]
	res = append(res, prev)
	for _, curr := range merged[1:] {
		if curr[0] <= prev[1] { // Overlap
			prev[1] = max(prev[1], curr[1]) // Merge
			continue
		}

		// not overlap, so continue merging and iterating
		prev = curr
		res = append(res, prev)
	}

	return res
}
