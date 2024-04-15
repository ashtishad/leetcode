package string

import "testing"

func Test_isIsomorphicBF(t *testing.T) {
	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"LC-1", args{s: "egg", t: "add"}, true},
		{"LC-2", args{s: "foo", t: "bar"}, false},
		{"LC-3", args{s: "paper", t: "title"}, true},
		{"LC-3", args{s: "s", t: "t"}, true},
		{"LC-4", args{s: "badc", t: "baba"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphicBF() = %v, want %v", got, tt.want)
			}
		})
	}
}
