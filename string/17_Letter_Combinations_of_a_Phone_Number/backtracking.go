package string

// Backtracking
// Time: O(4^n) and Space: O(4^n)
// Assuming each digit corresponds to the maximum of 4 possible letters for simplicity.
// letterCombinations returns all possible letter combinations that the number could represent.
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	phone := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var res []string

	var backtrack func(string, string)

	backtrack = func(combination, nextDigits string) {
		if nextDigits == "" {
			res = append(res, combination)
			return
		}

		letters := phone[nextDigits[0]-'2']
		for _, l := range letters {
			backtrack(combination+string(l), nextDigits[1:])
		}
	}

	backtrack("", digits)
	return res
}
