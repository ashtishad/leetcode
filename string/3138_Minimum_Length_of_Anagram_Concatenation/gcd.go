package string

func minAnagramLength(s string) int {
	m := make([]rune, 26)

	for _, ch := range s {
		m[ch-'a']++
	}

	min_freq := m[s[0]-'a']

	for i := 1; i < 26; i++ {
		min_freq = gcd(min_freq, m[i])
	}

	return len(s) / int(min_freq)
}

func gcd(a, b rune) rune {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
