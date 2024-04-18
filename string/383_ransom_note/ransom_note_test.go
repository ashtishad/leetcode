package string

import "testing"

func TestRansomNote(t *testing.T) {
	var Tests = []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{"LC-1", "a", "b", false},
		{"LC-2", "aa", "ab", false},
		{"LC-3", "aab", "aabbaba", true},
		{"LC-4", "aa", "aab", true},
	}

	for _, tc := range Tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := canConstruct(tc.ransomNote, tc.magazine); got != tc.want {
				t.Errorf("got:%v, want:%v", got, tc.want)
			}
		})
	}
}
