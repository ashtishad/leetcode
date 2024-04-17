package string

import "testing"

func TestStringToInteger(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"LC-1", "   -37", -37},
		{"LC-2", "4193 with words", 4193},
		{"LC-3", "42", 42},
		{"Edge-1", "-91283472332", -2147483648},
		{"Edge-2", "21474836460", 2147483647},
		{"Edge-3", "9223372036854775808", 2147483647},
		{"Edge-4", "-91283472332", -2147483648},
		{"Edge-5", "  ", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := myAtoi(tc.s)

			if got != tc.want {
				t.Errorf("got:%d, want:%d", got, tc.want)
			}
		})
	}
}
