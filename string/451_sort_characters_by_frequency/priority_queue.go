package string

import (
	"container/heap"
	"strings"
)

// Approach 2: Priority Queue (where frequency is the priority).
// Time: O(n+mlogm), since n>m, simplifies to O(nlogn).
// Where, n is the string len, and m is the number of unique chars in the string
// Space: O(n),  freqMap O(m)+ PQ O(m)+ Res String O(n)
func frequencySortPQ(s string) string {
	// counting frequencies, Time: O(n)
	count := make(map[string]int)
	for _, ch := range s {
		count[string(ch)]++
	}

	pq := make(PriorityQueue, len(count))
	var i int

	for ch, freq := range count {
		pq[i] = &CharFreq{char: ch, freq: freq}
		i++
	}

	// Time: O(m), where m is the number of unique chars in the count slice. worst case: m=n.
	heap.Init(&pq)

	res := ""

	// Time: single pop: O(logm), So For all pop operations is O(mlogm).
	for pq.Len() > 0 {
		cf := heap.Pop(&pq).(*CharFreq)
		res += strings.Repeat(cf.char, cf.freq)
	}

	return res
}

// CharFreq holds a character and its frequency.
type CharFreq struct {
	char string
	freq int
	idx  int // idx of the char in the heap, needed by the heap.Interface methods
}

// PriorityQueue implements heap.Interface and holds CharFreq items.
type PriorityQueue []*CharFreq

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].freq > pq[j].freq
} // highest freq first

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*CharFreq)
	item.idx = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.idx = -1  // for safety
	*pq = old[0 : n-1]
	return item
}
