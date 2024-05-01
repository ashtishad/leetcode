package string

// Two Pointers, Same Direction(Slow Fast)
// Time: O(n) and Space: O(1)
func isSubsequenceTP(s string, t string) bool {
	if s == "" {
		return true
	}

	var l, r int

	for r < len(t) {
		if s[l] == t[r] {
			l++

			if l == len(s) {
				return true
			}
		}

		r++
	}

	return false
}
