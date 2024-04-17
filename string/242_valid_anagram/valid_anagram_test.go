package string

import (
	"testing"
)

func TestValidAnagram(t *testing.T) {
	var Tests = []struct {
		name string
		s    string // lower-case english letters
		t    string // lower-case english letters
		want bool
	}{
		{"TC-1", "anagram", "nagaram", true},
		{"TC-2", "rat", "car", false},
		{"TC-3", "palm", "lamp", true},
		{"TC-4", "ab", "a", false},
		{"Single char true", "s", "s", true},
		// {"Unicode true", "täglich", "täglich", true},
		// {"Unicode false", "hola", "hélo", false},
	}

	for _, tc := range Tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isAnagram(tc.s, tc.t); got != tc.want {
				t.Errorf("got:%v, want:%v", got, tc.want)
			}
		})
	}
}

// func BenchmarkValidAnagram(b *testing.B) {
// 	test := struct {
// 		s    string
// 		t    string
// 		want bool
// 	}{

// 		s:    strings.Repeat("a", 50) + strings.Repeat("b", 50),
// 		t:    strings.Repeat("b", 20) + strings.Repeat("a", 30) + strings.Repeat("b", 30) + strings.Repeat("a", 20),
// 		want: true,
// 	}

// 	for i := 0; i < b.N; i++ {
// 		isAnagram(test.s, test.t)
// 	}
// }
