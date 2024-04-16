package string

import "unicode"

// brute force (check equality with reversed)
// time: O(n) and space: O(n)
func isPalindromeBF(s string) bool {
	processed := ""

	for _, ch := range s {
		// if non-alpha, then lowercase, and add char to processed s
		if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
			ch = unicode.ToLower(ch)
			processed += string(ch)
		}
	}

	reversed := ""
	for i := len(processed) - 1; i >= 0; i-- {
		reversed += string(processed[i])
	}

	return reversed == processed
}
