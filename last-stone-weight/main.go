/*
Package last_stone_weight
https://leetcode-cn.com/problems/last-stone-weight/

1046. 最后一块石头的重量

有一堆石头，每块石头的重量都是正整数。

每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。

提示：
	1 <= stones.length <= 30
	1 <= stones[i] <= 1000

*/
package last_stone_weight

import (
	"container/heap"
)

// --- 自己

/*
方法一: 优先队列/堆

时间复杂度：
空间复杂度：
*/
func lastStoneWeight(stones []int) int {
	h := InitHeap(stones)
	heap.Init(&h)
	for h.Len() > 1 {
		x, y := heap.Pop(&h).(int), heap.Pop(&h).(int)
		heap.Push(&h, x-y)
	}
	return heap.Pop(&h).(int)
}

type InitHeap []int

func (h InitHeap) Len() int           { return len(h) }
func (h InitHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h InitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *InitHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *InitHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
