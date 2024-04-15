package string

// Hash Table
// Time: O(n) and Space: O(1)
// when a smaller value appears before a value, it represents Subtraction.
// when a larger/equal value appears before a  value, it represents Addition.
func romanToInt(s string) int {
	romans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res, n := 0, len(s)

	for i := 0; i < n; i++ {
		curVal := romans[s[i]]

		// when a smaller value appears before a value, it represents Subtraction
		if i+1 < n && curVal < romans[s[i+1]] {
			res -= curVal
			continue
		}

		// when a larger/equal value appears before a  value, it represents Addition
		res += romans[s[i]]

	}

	return res
}
