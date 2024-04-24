package dsa

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

// example usage
// func topKFrequentPQ(nums []int, k int) []int {
// 	count := make(map[int]int)
// 	for _, v := range nums {
// 		count[v]++
// 	}

// 	pq := make(PriorityQueue, 0, len(count))
// 	heap.Init(&pq)

// 	for v, f := range count {
// 		item := &Item{
// 			val:      v,
// 			priority: f,
// 		}

// 		heap.Push(&pq, item)
// 	}

// 	res := make([]int, 0)
// 	for k > 0 {
// 		item := heap.Pop(&pq).(*Item)
// 		res = append(res, item.val)
// 		k--
// 	}

// 	return res
// }
