package string

// Character Count Array
// Time: O(n) and Space: O(1)
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	cnt := make([]rune, 26)

	for _, ch := range magazine {
		cnt[ch-'a']++
	}

	for _, ch := range ransomNote {
		if cnt[ch-'a'] == 0 {
			return false
		}

		cnt[ch-'a']--
	}

	return true
}
