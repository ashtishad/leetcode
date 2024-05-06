package array

// Two Pointers
// Time: o(n) and Space: O(1)
// area = min(left_height, right height) * width(distance between walls)
func maxArea(x []int) int {
	l, r := 0, len(x)-1

	var res int

	for l < r {
		area := min(x[l], x[r]) * (r - l)
		res = max(res, area)

		if x[l] <= x[r] {
			l++
		} else {
			r--
		}
	}

	return res
}
