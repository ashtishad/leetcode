package main

func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	var l, res int

	for r, ch := range s {
		seen[ch]++

		for seen[ch] > 1 {
			seen[rune(s[l])]--
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}
