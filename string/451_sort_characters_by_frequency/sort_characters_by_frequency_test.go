package string

import "testing"

func TestSortCharsByFreq(t *testing.T) {
	var Tests = []struct {
		name string
		s    string // s consists of uppercase and lowercase English letters and digits.
		want []string
	}{
		{"LC-1", "tree", []string{"eetr", "eert"}},
		{"LC-2", "cccaaa", []string{"cccaaa", "aaaccc"}},
		{"LC-3", "Acacbcb", []string{"cccbbAa", "cccbbaA"}},
	}

	for _, tc := range Tests {
		t.Run(tc.name, func(t *testing.T) {
			got := frequencySort(tc.s)

			var matched bool

			for _, want := range tc.want {
				if got == want {
					matched = true
				}
			}

			if !matched {
				t.Errorf("got:%s, want: any of %s", got, tc.want)
			}
		})
	}
}
