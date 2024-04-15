package array

// counting sort, counting 0's, 1's and 2's'
// time: O(n) and space: O(1)
func sortColorsCS(nums []int) {
	count := make([]int, 3)

	for _, v := range nums {
		switch v {
		case 0:
			count[0]++
		case 1:
			count[1]++
		case 2:
			count[2]++
		}
	}

	var idx int // to track indexes of nums where we are about to swap values.
	for i := range count {
		for range count[i] {
			nums[idx] = i
			idx++
		}
	}
}
