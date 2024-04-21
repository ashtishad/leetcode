package array

import "slices"

// Find Pivot and Then choose which half to search
// time: O(logn) and space: O(1)
func search(nums []int, target int) int {
	n := len(nums)

	isRotated := nums[0] > nums[n-1]
	if !isRotated {
		if idx, ok := slices.BinarySearch(nums, target); ok {
			return idx
		}

		return -1
	}

	// rotated
	pivot := findPivotIdx(nums)
	idx, ok := -1, false

	switch {
	case nums[pivot] == target: // target is at pivot
		return pivot

	case nums[pivot] <= target && target <= nums[n-1]: // Target in right half
		if idx, ok = slices.BinarySearch(nums[pivot:], target); ok {
			return idx + pivot
		}

	case nums[0] <= target: // Target in left half
		if idx, ok = slices.BinarySearch(nums[:pivot], target); ok {
			return idx
		}
	}

	return -1
}

func findPivotIdx(nums []int) int {
	l, h := 0, len(nums)-1

	for l < h {
		mid := l + ((h - l) >> 1)

		if nums[mid] > nums[h] {
			l = mid + 1
		} else {
			h = mid
		}
	}

	return l
}
