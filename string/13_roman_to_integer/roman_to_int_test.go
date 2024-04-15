package string

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"LC-1", "IX", 9},
		{"LC-1", "XI", 11},
		{"LC-1", "LVIII", 58},
		{"LC-1", "MCMXCIV", 1994},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
