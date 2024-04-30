package string

import (
	"strconv"
	"strings"
)

// Recursion
// Time: O(n*2^n) and space: O(2^n)
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay(n - 1)
	return countDigits(prev)
}

func countDigits(s string) string {
	var sb strings.Builder
	var i, j int

	for i = 0; i < len(s); {
		for j = i; j < len(s); j++ {
			if s[j] != s[i] {
				break
			}
		}

		sb.WriteString(strconv.Itoa(j - i))
		sb.WriteByte(s[i])

		i = j
	}

	return sb.String()
}
