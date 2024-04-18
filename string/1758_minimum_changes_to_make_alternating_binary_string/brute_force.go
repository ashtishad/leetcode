package string

// Brute Force
// Time O(n) and Space; O(1)
// Either start with a '0',  '010101010..' or with a '1','10101010..'
// The number of changes needed to reach it from the given string
func opsNeededToMake(s string, curr rune) int {
	var ops int

	for _, ch := range s {
		if ch != curr {
			ops++
		}

		if curr == '1' {
			curr = '0'
		} else {
			curr = '1'
		}
	}

	return ops
}

func minOperations(s string) int {
	zero := opsNeededToMake(s, '0')
	one := opsNeededToMake(s, '1')

	return min(zero, one)
}
