package array

// Hash Table
// Time O(n) and Space O(n)
func majorityElement(nums []int) int {
	n := len(nums)

	seen := make(map[int]int, n) // k: num, v: frequency

	for _, v := range nums {
		seen[v]++

		if seen[v] > n/2 {
			return v
		}
	}

	return -1
}
