package string

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"0100", 1},
		{"10", 0},
		{"1111", 2},
		{"", 0},
	}

	for _, tc := range tests {
		got := minOperations(tc.input)
		if got != tc.want {
			t.Errorf("Expected %d, got %d for input %s", tc.want, got, tc.input)
		}
	}
}
