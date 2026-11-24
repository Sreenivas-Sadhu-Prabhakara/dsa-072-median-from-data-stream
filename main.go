package main

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

// MedianFinder keeps the running median of a stream: a max-heap holds the
// lower half, a min-heap the upper half, balanced to differ by at most one.
type MedianFinder struct {
	low  maxHeap // lower half, top = its largest
	high minHeap // upper half, top = its smallest
}

// Add inserts a value in O(log n).
func (m *MedianFinder) Add(v int) {
	if m.low.Len() == 0 || v <= m.low[0] {
		heap.Push(&m.low, v)
	} else {
		heap.Push(&m.high, v)
	}
	// Rebalance so low holds equal count or exactly one more.
	if m.low.Len() > m.high.Len()+1 {
		heap.Push(&m.high, heap.Pop(&m.low))
	} else if m.high.Len() > m.low.Len() {
		heap.Push(&m.low, heap.Pop(&m.high))
	}
}

// Median returns the current median; ok is false while the stream is empty.
func (m *MedianFinder) Median() (float64, bool) {
	switch {
	case m.low.Len() == 0:
		return 0, false
	case m.low.Len() > m.high.Len():
		return float64(m.low[0]), true
	default:
		return (float64(m.low[0]) + float64(m.high[0])) / 2, true
	}
}

func main() {
	var mf MedianFinder
	for _, v := range []int{5, 15, 1, 3} {
		mf.Add(v)
		med, _ := mf.Median()
		fmt.Println(med)
	}
}
