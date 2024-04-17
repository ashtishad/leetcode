package string

import (
	"math"
	"strings"
)

// Time: O(n) and Space: O(n)
func myAtoi(s string) int {
	// trim spaces
	s = strings.Trim(s, " ")

	if s == "" {
		return 0
	}

	// check sign
	sign := 1
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -sign
		}

		s = s[1:]
	}

	// read digits(skip letters)
	var res int
	for _, ch := range []byte(s) {
		ch -= '0' // ascii to digit

		if ch < 0 || ch > 9 {
			break
		}

		res = res*10 + int(ch)

		// there might be a scenario res would gone beyond int64, and we would find a garbage value with - as sign bit.
		if res > math.MaxInt32 {
			break
		}
	}

	if sign == -1 {
		res = -res
	}

	// int32 bound check
	if res < math.MinInt32 {
		res = math.MinInt32
	}

	if res > math.MaxInt32 {
		res = math.MaxInt32
	}

	return res
}
