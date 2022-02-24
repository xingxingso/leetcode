/*
Package zui_xiao_de_kge_shu_lcof
https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

剑指 Offer 40. 最小的k个数

输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

限制：
	0 <= k <= arr.length <= 10000
	0 <= arr[i] <= 10000

*/
package zui_xiao_de_kge_shu_lcof

import (
	"container/heap"
)

// --- 自己

/*
方法一: 堆
	// 大顶堆，小于大顶堆的堆顶，才可以加入

时间复杂度：O(n)
空间复杂度：O(1)
*/
func getLeastNumbers(arr []int, k int) []int {
	h := &IntHeap{}
	for i := 0; i < len(arr); i++ {
		heap.Push(h, arr[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	//fmt.Println(h)

	return *h
}

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
