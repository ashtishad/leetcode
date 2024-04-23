package array

import "math"

func findMaxAveragePFS(nums []int, k int) float64 {
	n := len(nums)
	if n == 1 {
		return float64(nums[0])
	}

	res := -math.MaxFloat64

	pfs := make([]int, n+1)
	for i := range nums {
		pfs[i+1] = pfs[i] + nums[i]
	}

	// pfs: 0, 1, 13, 8, 2, 52, 55
	// pfs: 0,0,1,2,5,8

	var l int
	for r := k; r <= n; r++ {
		sum := pfs[r] - pfs[l]
		avg := float64(sum) / float64(k)
		res = max(res, avg)
		l++
	}

	return res
}
