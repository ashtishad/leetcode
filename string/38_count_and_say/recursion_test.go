package string

import (
	"strconv"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	testCases := []struct {
		n    int
		want string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
		{7, "13112221"},
		{8, "1113213211"},
		{9, "31131211131221"},
		{10, "13211311123113112211"},
	}

	for _, tc := range testCases {
		t.Run("n="+strconv.Itoa(tc.n), func(t *testing.T) {
			got := countAndSay(tc.n)
			if got != tc.want {
				t.Errorf("countAndSay(%d) = %s, want %s", tc.n, got, tc.want)
			}
		})
	}
}
