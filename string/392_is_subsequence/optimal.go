package string

// Travserse and track matched characters
func isSubsequence(s string, t string) bool {
	var k int

	for i := 0; i < len(t); i++ {
		if k < len(s) && s[k] == t[i] {
			k++
		}

	}

	return len(s) == k
}
