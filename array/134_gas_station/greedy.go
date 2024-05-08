package array

// Greedy
// Time: O(n) and Space: O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	var totalTank, currTank, start int

	for i := range gas {
		totalTank += gas[i] - cost[i]
		currTank += gas[i] - cost[i]

		if currTank < 0 {
			// means we cannot start from this station, so start from next
			start = i + 1
			currTank = 0
		}
	}

	if totalTank >= 0 {
		return start
	}

	return -1
}
