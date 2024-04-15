package array

// Dutch National Flag Algorithm/Three way partition
// Time: O(n) and Space: O(1)
//
// //          0         1        yet to be sorted       2
// -----------------------------------------------------------
// nums: |          |          |                  |           |
// -----------------------------------------------------------
// //          ^          ^             ^
// //         low         mid                    high
//
// traverse the array with mid pointer.
func sortColors(nums []int) {
	if len(nums) == 1 {
		return
	}

	var low, mid int
	high := len(nums) - 1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}
