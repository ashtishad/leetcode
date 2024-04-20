package dsa

import "testing"

type test struct {
	name      string
	nums      []int
	target    int
	wantPos   int
	wantFound bool
}

func TestBinarySearch(t *testing.T) {
	tests := []test{
		{"Target Found", []int{-1, 0, 3, 5, 9, 12}, 9, 4, true},
		{"Target Not Found", []int{-1, 0, 3, 5, 9, 12}, 13, 6, false},
		{"Empty Slice", []int{}, 5, 0, false},
		{"Single Element Found", []int{5}, 5, 0, true},
		{"Single Element Not Found", []int{2}, 5, 1, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotPos, gotFound := binarySearch(tc.nums, tc.target)

			if gotPos != tc.wantPos || gotFound != tc.wantFound {
				t.Errorf("binarySearch(%v, %d) = (%d, %v), want (%d, %v)", tc.nums, tc.target, gotPos, gotFound, tc.wantPos, tc.wantFound)
			}
		})
	}
}
