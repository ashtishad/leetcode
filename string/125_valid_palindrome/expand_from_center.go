package string

import (
	"regexp"
	"strings"
)

// Expand from center
// Time: O(n) and Space:O(1)
func isPalindromeEFC(s string) bool {
	reg, _ := regexp.Compile("[^A-Za-z0-9]")
	s = strings.ToLower(reg.ReplaceAllString(s, ""))

	// calculate middle indexes
	l := len(s) / 2
	r := l

	// isEven, "abba", adjust left pointer
	if len(s)%2 == 0 {
		l--
	}

	return expandFromCenter(s, l, r)
}

// expandFromCenter function has a worst-case time complexity of O(n), but for typical inputs and within the context of the
// overall algorithm, it's often treated as having an "effectively constant" O(1) time complexity.
func expandFromCenter(s string, l, r int) bool {
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			return false
		}

		l--
		r++
	}

	return true
}
