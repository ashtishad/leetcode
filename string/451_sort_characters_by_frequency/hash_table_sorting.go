package string

import (
	"sort"
	"strings"
)

// Approach 1: Hash Table + Sort by frequency
// Time: O(nlogn) and Space: O(n)
func frequencySortHS(s string) string {
	count := make(map[string]int)
	for _, ch := range s {
		count[string(ch)]++
	}

	// sorting a map by freq(value) requires, create a key-value slice from it, then sort that slice
	type keyVal struct {
		char string
		freq int
	}

	kvSlice := make([]keyVal, 0, len(count))
	for ch, cnt := range count {
		kvSlice = append(kvSlice, keyVal{ch, cnt})
	}

	// kvSlice: [{c 3} {a 1} {b 2} {A 1}]

	sort.Slice(kvSlice, func(i, j int) bool {
		return kvSlice[i].freq > kvSlice[j].freq
	})

	// kvSlice: [{c 3} {b 2} {a 1} {A 1}]

	res := ""
	for _, kv := range kvSlice {
		res += strings.Repeat(kv.char, kv.freq)
	}

	return res
}

/*
// Last Portion can be done by this:
var sb strings.Builder
	for _, kv := range kvSlice {
		sb.WriteString(strings.Repeat(kv.char, kv.freq))
	}

	return sb.String()
*/
