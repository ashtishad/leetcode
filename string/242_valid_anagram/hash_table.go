package string

// Hash Table
// time: O(n) and space: O(1)
func isAnagramHT(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	seen := make(map[rune]int, len(s))

	for _, c := range s {
		seen[c]++
	}

	for _, c := range t {
		seen[c]--

		f, ok := seen[c]
		if !ok || f < 0 {
			return false
		}

	}

	return true
}
