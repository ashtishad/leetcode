package array

// Binary Search
// Time: O(w*logn) and space; O(1)
// w is number of weights and n is number of capacities
func shipWithinDays(weights []int, days int) int {
	sumW, maxW := 0, weights[0]
	for _, w := range weights {
		sumW += w
		maxW = max(maxW, w)
	}

	l, h := maxW, sumW

	res := h
	for l <= h {
		mid := l + ((h - l) >> 1)

		if canFit(mid, weights, days) {
			res = min(res, mid)
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return res
}

func canFit(capacity int, weights []int, days int) bool {
	cw, d := 0, 1 // curr weight, days taken

	for _, w := range weights {
		if cw+w <= capacity {
			cw += w
		} else {
			d++
			cw = w
		}
	}

	return d <= days
}
