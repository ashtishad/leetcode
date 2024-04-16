package string

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Simple Palindrome",
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			name:  "Not a Palindrome",
			input: "race a car",
			want:  false,
		},
		{
			name:  "Single Character",
			input: "a",
			want:  true,
		},
		{
			name:  "Empty String",
			input: "",
			want:  true,
		},
		{
			name:  "Numbers and Symbols",
			input: "0P;!?:",
			want:  false,
		},
		{
			name:  "Numbers and Symbols",
			input: "P1p:",
			want:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isPalindrome(tc.input)
			if got != tc.want {
				t.Errorf("Expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
