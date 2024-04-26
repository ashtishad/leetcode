package array

import "slices"

// time: O(n^2) space: O(1)
func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	res := make([][]int, 0)

	for i, v := range nums {
		if i > 0 && nums[i-1] == v {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := v + nums[l] + nums[r]

			switch {
			case sum < 0:
				l++
			case sum > 0:
				r--
			default:
				res = append(res, []int{v, nums[l], nums[r]})
				l++

				// skipping duplicates
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return res
}
