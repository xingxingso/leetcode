/*
https://leetcode-cn.com/problems/find-median-from-data-stream/

295. 数据流的中位数

中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
例如，
	[2,3,4] 的中位数是 3
	[2,3] 的中位数是 (2 + 3) / 2 = 2.5
设计一个支持以下两种操作的数据结构：
	void addNum(int num) - 从数据流中添加一个整数到数据结构中。
	double findMedian() - 返回目前所有元素的中位数。


进阶:
	1. 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
	2. 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？

*/
package find_median_from_data_stream

import "container/heap"

// type MedianFinder struct {}
// /** initialize your data structure here. */
// func Constructor() MedianFinder {}
// func (this *MedianFinder) AddNum(num int)  {}
// func (this *MedianFinder) FindMedian() float64 {}
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

// --- 自己

/*
方法一: 双顶堆

时间复杂度：
空间复杂度：
*/

// 小顶堆 -中位数右边的
type bigHeap []int

func (h bigHeap) Len() int           { return len(h) }
func (h bigHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h bigHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *bigHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *bigHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *bigHeap) Peek() interface{} { return (*h)[0] }

// 大顶堆 -中位数左边的
type littleHeap []int

func (h littleHeap) Len() int           { return len(h) }
func (h littleHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h littleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *littleHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *littleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *littleHeap) Peek() interface{} { return (*h)[0] }

type MedianFinder struct {
	bigHeap    *bigHeap
	littleHeap *littleHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		bigHeap:    &bigHeap{},
		littleHeap: &littleHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 直接加入大顶堆
	heap.Push(this.bigHeap, num)
	// 保证两个顶堆的 一大 一小
	heap.Push(this.littleHeap, heap.Pop(this.bigHeap))

	// 保持两个堆差值不能超过1
	for this.bigHeap.Len() < this.littleHeap.Len() {
		heap.Push(this.bigHeap, heap.Pop(this.littleHeap))
	}

	// fmt.Println("after", this.bigHeap.Len(), this.littleHeap.Len())
	// fmt.Println(this.bigHeap)
	// fmt.Println(this.littleHeap)
}

// 早期的复杂做法，不过好像效率高一些
// func (this *MedianFinder) AddNum(num int) {
// 	// fmt.Println(this.bigHeap)
// 	// fmt.Println(this.littleHeap)
// 	if this.bigHeap.Len() == 0 {
// 		heap.Push(this.bigHeap, num)
// 		return
// 	}
// 	big := this.bigHeap.Peek().(int)
// 	if this.littleHeap.Len() == 0 {
// 		if big > num {
// 			heap.Push(this.littleHeap, num)
// 			return
// 		}
// 		heap.Pop(this.bigHeap)
// 		heap.Push(this.littleHeap, big)
// 		heap.Push(this.bigHeap, num)
// 		return
// 	}
// 	little := this.littleHeap.Peek().(int)
// 	if num > big {
// 		heap.Push(this.bigHeap, num)
// 	} else if num < little {
// 		heap.Push(this.littleHeap, num)
// 	} else {
// 		heap.Push(this.bigHeap, num)
// 	}
// 	for this.bigHeap.Len() > this.littleHeap.Len()+1 {
// 		big := heap.Pop(this.bigHeap)
// 		heap.Push(this.littleHeap, big)
// 	}
// 	for this.littleHeap.Len() > this.bigHeap.Len() {
// 		little := heap.Pop(this.littleHeap)
// 		heap.Push(this.bigHeap, little)
// 	}
// 	// fmt.Println("after", this.bigHeap.Len(), this.littleHeap.Len())
// 	// fmt.Println(this.bigHeap)
// 	// fmt.Println(this.littleHeap)
// }

func (this *MedianFinder) FindMedian() float64 {
	if this.bigHeap.Len() > this.littleHeap.Len() {
		value := this.bigHeap.Peek()
		return float64(value.(int))
	}

	left, right := this.bigHeap.Peek(), this.littleHeap.Peek()
	return float64(left.(int)+right.(int)) / 2
}
