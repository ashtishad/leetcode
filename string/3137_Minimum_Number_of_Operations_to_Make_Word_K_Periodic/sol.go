package string

// Time: O(n) and Space: O(n)
func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)

	seq := make(map[string]int)

	var maxFreq, cnt int

	for i := 0; i < n; i = i + k {
		t := word[i : i+k]
		seq[t]++

		maxFreq = max(maxFreq, seq[t])
		cnt++
	}

	return cnt - maxFreq
}
