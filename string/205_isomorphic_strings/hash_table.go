package string

// Hash Table
// Time O(n) and Space:O(1), just ascii characters(upper bound 128)
// Imagine two secret codes, ciphers, just characters are replaced. we are just checking is this string is a valid cipher of another string.
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapST := make(map[byte]byte) // k: char of s, v:char of t
	mapTS := make(map[byte]byte) // k: char of t, v:char of s

	for i := range len(s) {
		cs, ct := s[i], t[i]

		if v, ok := mapST[cs]; ok && v != ct {
			return false
		}

		mapST[cs] = ct

		if v, ok := mapTS[ct]; ok && v != cs {
			return false
		}

		mapTS[ct] = cs
	}

	return true
}

/*
Two strings are isomorphic if a one-to-one mapping is possible for every character of the first
string ‘str1’ to every character of the second string ‘str2’ while preserving the order of the characters.
*/
