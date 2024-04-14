package array

import "slices"

// Binary Search, find compliment after i
// time: O(n) space: O(1)
// out order matters, sorted ASC, idx starts from 1, exactly one solution
// Input: nums = [2,3,4,5,7,9], target = 6, Output: [1,3]
// Using Golang slices.BinarySearch
func twoSumIIBS(nums []int, target int) []int {
	n := len(nums)

	for i, v := range nums[:n-1] {
		// slices.BinarySearch returns first occurrence, that's why searching after i
		j, ok := slices.BinarySearch(nums[i+1:], target-v)

		if ok {
			// nums[i+1:] this return j as 0 based idx starting after i, so, second idx = i+1+j+1 = i+j+2
			return []int{i + 1, i + j + 2}
		}
	}

	return []int{}
}

// // Easier Syntax:Using Golang slices.BinarySearch
// func twoSumIIBS(nums []int, target int) []int {
// 	for i, v := range nums {
// 		j, ok := slices.BinarySearch(nums, target-v)
//
// 		if ok && i != j {
// 			if i > j {
// 				return []int{j + 1, i + 1}
// 			}
//
// 			return []int{i + 1, j + 1}
// 		}
// 	}
//
// 	return []int{}
// }
