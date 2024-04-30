package string

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{"Empty input", "", []string{}},
		{"Single digit", "2", []string{"a", "b", "c"}},
		{"Two digits", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"Digits with 4 letters", "7", []string{"p", "q", "r", "s"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.digits)
			sort.Strings(got) // Sorting because the order of combinations is not guaranteed
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations(%s) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}
