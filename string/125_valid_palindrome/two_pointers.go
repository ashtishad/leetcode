package string

import "unicode"

// Two Pointers
// Time: O(n) and Space: O(1)
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		switch {
		case !isAlphanumeric(s[l]):
			l++
		case !isAlphanumeric(s[r]):
			r--
		case unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])):
			return false
		default:
			l++
			r--
		}
	}

	return true
}

func isAlphanumeric(c byte) bool {
	return unicode.IsDigit(rune(c)) || unicode.IsLetter(rune(c))
}
