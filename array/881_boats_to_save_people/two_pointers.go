package array

import "slices"

// Sorting + Two Pointer
func numRescueBoats(weights []int, capacity int) int {
	slices.Sort(weights)

	var boats int
	low, high := 0, len(weights)-1

	for low <= high {
		if weights[low]+weights[high] > capacity {
			high--
			boats++ // high will get into the boat alone
		} else {
			low++
			high--
			boats++
		}
	}

	return boats
}
