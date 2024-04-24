package main

import (
	"container/heap"
	"fmt"
)

// Time: O(N + M log M)
// 1. Creating count map O(n)
// 2. Create and Initialize PQ O(M), where M is no of unique elems.
// 3. Pushing M elems to PQ; O(M*logM)
// 4. Popping the top K elements: O(K log M)
// Space: O(N)
func topKFrequentPQ(nums []int, k int) []int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}

	pq := make(PriorityQueue, 0, len(count))
	heap.Init(&pq)

	for v, f := range count {
		item := &Item{
			val:      v,
			priority: f,
		}

		heap.Push(&pq, item)
	}

	res := make([]int, 0)
	for k > 0 {
		item := heap.Pop(&pq).(*Item)
		res = append(res, item.val)
		k--
	}

	return res
}

type Item struct {
	val      int
	priority int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Pop will give us the highest
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func main() {
	fmt.Println(topKFrequentPQ([]int{1, 1, 1, 2, 2, 3}, 2))
}
