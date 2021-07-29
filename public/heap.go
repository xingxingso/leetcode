package public

import (
	"container/heap"
	"sort"
)

//IntHeap 小顶堆类
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//Peek 获取堆顶元素 不弹出
func (h *IntHeap) Peek() interface{} { return (*h)[0] }

// 使用案例
func minMeetingRoomsO1(intervals [][]int) int {
	h := &IntHeap{}
	heap.Init(h)
	res := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); i++ {
		for len(*h) != 0 && intervals[i][0] >= (*h)[0] {
			heap.Pop(h)
		}
		heap.Push(h, intervals[i][1])
		res = max(res, len(*h))
	}
	return res
}
