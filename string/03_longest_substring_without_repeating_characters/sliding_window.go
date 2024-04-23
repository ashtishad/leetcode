package main

func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int) // k:ch v:idx
	var l, res int

	for r, ch := range s {
		if idx, ok := seen[ch]; ok && idx >= l {
			l = idx + 1 // start new window
		}

		seen[ch] = r
		res = max(res, r-l+1)
	}

	return res
}
