package string

import "testing"

func TestCompareVersion(t *testing.T) {
	testCases := []struct {
		name     string
		version1 string
		version2 string
		want     int
	}{
		{"Same version", "1.0", "1.0", 0},
		{"First greater", "1.2", "1.1", 1},
		{"Second greater", "1.0.1", "1.1", -1},
		{"Leading zeros", "1.01", "1.1", 0},
		{"More segments", "1.0.0", "1", 0},
		{"Second longer", "1.0", "1.0.1", -1},
		{"First longer positive", "1.0.1", "1.0", 1},
		{"Empty strings", "", "", 0},
		{"First empty", "", "1.1", -1},
		{"Second empty", "1.1", "", 1},
		{"Version with large numbers", "1.10.100", "1.10.99", 1},
		{"Version with zeros", "0.1", "0.0.1", 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := compareVersion(tc.version1, tc.version2)
			if got != tc.want {
				t.Errorf("compareVersion(%q, %q) = %d, want %d", tc.version1, tc.version2, got, tc.want)
			}
		})
	}
}
