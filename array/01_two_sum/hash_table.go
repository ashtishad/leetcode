package array

// Hash Table, finding compliment of the target.
// Time: O(n) and Space: O(n)
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) //k:val, v:idx

	for i, v := range nums {
		if idx, ok := seen[target-v]; ok {
			return []int{idx, i}
		}

		seen[v] = i
	}

	return []int{}
}
