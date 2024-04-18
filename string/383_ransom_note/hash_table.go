package string

// Approach: Hash Table
// Time: O(n) and Space: O(n)
func canConstructHashTable(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	cnt := make(map[rune]int)

	for _, ch := range magazine {
		cnt[ch]++
	}

	for _, ch := range ransomNote {
		if v, ok := cnt[ch]; !ok || v == 0 {
			return false
		}

		cnt[ch]--
	}

	return true
}
