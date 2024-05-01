package string

import "testing"

func TestIsSubsequence(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"Both empty", "", "", true},
		{"Empty subsequence", "", "abc", true},
		{"Empty string", "abc", "", false},
		{"Subsequence at start", "abc", "abcdef", true},
		{"Subsequence at end", "def", "abcdef", true},
		{"Subsequence in middle", "bce", "abcdef", true},
		{"Not a subsequence", "ace", "abcde", true},
		{"Single character true", "c", "abcde", true},
		{"Single character false", "z", "abcde", false},
		{"Long subsequence", "ace", "aabbccddee", true},
		{"Complete match", "abcde", "abcde", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isSubsequence(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("isSubsequence(%q, %q) = %t, want %t", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
