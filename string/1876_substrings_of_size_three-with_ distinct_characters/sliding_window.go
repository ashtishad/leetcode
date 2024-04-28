package string

// Input: s = "aababcabc" , Output: 4
func countGoodSubstrings(s string) int {
	seen := make(map[byte]int)
	var l, res int

	for r := 0; r < len(s); r++ {
		seen[s[r]]++

		if r-l+1 > 3 {
			seen[s[l]] -= 1

			if seen[s[l]] == 0 {
				delete(seen, s[l])
			}

			l++
		}

		if len(seen) == 3 {
			res++
		}
	}

	return res
}
