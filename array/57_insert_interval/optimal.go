package array

// Time: O(n) and Space: O(1)
// sorted ASC, non-overlapping intervals
// [[1,3], [6,9]]
func insertX(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}

	for i := range intervals {
		switch {
		// new interval appears before, newInterval = {0,1}
		case newInterval[1] < intervals[i][0]:
			res = append(res, newInterval)
			return append(res, intervals[i:]...)

		// after all intervals, newInterval = {10,12}
		case newInterval[0] > intervals[i][1]:
			res = append(res, intervals[i])
		// otherwise, we continue modifying newInterval={min, max}
		default:
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
		}
	}

	res = append(res, newInterval)
	return res
}

/*
Let's say the existing intervals are: [[1,3], [6,9]]

Empty list edge case: If the list were empty, and newInterval = [2,5], the result would be [[2,5]].
Before all intervals: If newInterval = [0,1], the result would be [[0,1], [1,3], [6,9]].
After all intervals: If newInterval = [10,12], the result would be [[1,3], [6,9], [10,12]].
Multiple overlaps: If newInterval = [4,8], the result would be [[1,3], [4,9]].
Complete containment: If newInterval = [2,7], the result would be [[1,7]].

*/
