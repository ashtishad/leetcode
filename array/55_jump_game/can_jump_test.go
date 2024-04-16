package array

import "testing"

func TestCanJump(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Basic Reachable",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "Unreachable",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "Single Jump",
			nums: []int{2, 0},
			want: true,
		},
		{
			name: "Zero at Start",
			nums: []int{0, 2, 3},
			want: false,
		},
		{
			name: "Only Zero",
			nums: []int{0},
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := canJump(tc.nums)
			if got != tc.want {
				t.Errorf("Expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
