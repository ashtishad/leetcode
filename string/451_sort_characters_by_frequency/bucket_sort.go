package string

import "strings"

// Bucket Sort
// Time: O(n) and Space: O(n)
func frequencySort(s string) string {
	count := make(map[string]int)
	for _, ch := range s {
		count[string(ch)]++
	}

	// Bucket Sort, Idea is to put frequency in index, and all chars with same frequency in inner array.
	// Creating buckets up to len s + 1, because of 0 based indexing. So, freq 1 will be idx 1
	buckets := make([][]string, len(s)+1)
	for ch, f := range count {
		buckets[f] = append(buckets[f], ch)
	}

	// Now iterate from end ofd the bucket to construct res string with higher frequencies comes first.
	res := ""
	for f := len(buckets) - 1; f > 0; f-- {
		for _, ch := range buckets[f] {
			res += strings.Repeat(ch, f)
		}
	}

	return res
}
