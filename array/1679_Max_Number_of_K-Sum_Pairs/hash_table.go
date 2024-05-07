package array

// Time: O(n) and Space: O(n)
func maxOperations(nums []int, k int) int {
	var ops int

	count := make(map[int]int)

	for _, v := range nums {
		if val, ok := count[k-v]; ok && val > 0 {
			count[k-v]--
			ops++
			continue
		}

		count[v]++

	}

	return ops
}
