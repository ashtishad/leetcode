package string

// Char Count Array
// time: O(n) and space: O(1)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make([]rune, 26)

	for _, ch := range s {
		count[ch-'a']++
	}

	for _, ch := range t {
		count[ch-'a']--

		if count[ch-'a'] < 0 {
			return false
		}
	}

	return true
}
