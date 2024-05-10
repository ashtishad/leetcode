package array

import "container/heap"

// Min-Heap of fractions
// Time: O(n^2*logn + klogn) and Space: O(n^2)
func kthSmallestPrimeFraction(nums []int, k int) []int {
	if len(nums) == 2 {
		return nums
	}

	h := make(minHeap, 0)
	for i := 0; i+1 < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			h = append(h, &Fraction{
				i: nums[i], j: nums[j],
				val: float64(nums[i]) / float64(nums[j]),
			})
		}
	}

	heap.Init(&h)

	var f *Fraction
	for k > 0 {
		f = heap.Pop(&h).(*Fraction)
		k--
	}

	return []int{f.i, f.j}
}

type Fraction struct {
	i, j int     // nums
	val  float64 // fractions
}

type minHeap []*Fraction

func (h minHeap) Len() int { return len(h) }

func (h minHeap) Less(i, j int) bool { return h[i].val < h[j].val }

func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) { *h = append(*h, x.(*Fraction)) }

func (h *minHeap) Pop() any {
	n := h.Len()
	old := *h
	fraction := old[n-1]
	old[n-1] = nil
	*h = old[:n-1]
	return fraction
}
