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

import "fmt"

//type MedianFinder struct {}
///** initialize your data structure here. */
//func Constructor() MedianFinder {}
//func (this *MedianFinder) AddNum(num int)  {}
//func (this *MedianFinder) FindMedian() float64 {}
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/

// 大顶堆
type bigHeap []int

func (h bigHeap) Len() int           { return len(h) }
func (h bigHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h bigHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *bigHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	//*h = append(*h, x.(*element))
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

// 小顶堆
type littleHeap []int

func (h littleHeap) Len() int           { return len(h) }
func (h littleHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h littleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *littleHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	//*h = append(*h, x.(*element))
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
	bigHeap    bigHeap
	littleHeap littleHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	fmt.Println("bigHeap:", this.bigHeap)
	fmt.Println("littleHeap:", this.littleHeap)
	if this.bigHeap.Len() == 0 {
		this.bigHeap.Push(num)
		return
	}
	big := this.bigHeap.Peek().(int)
	if this.littleHeap.Len() == 0 {
		if big > num {
			this.littleHeap.Push(num)
			return
		}
		this.bigHeap.Pop()
		this.littleHeap.Push(big)
		this.bigHeap.Push(num)
		return
	}
	little := this.littleHeap.Peek().(int)
	if num > big {
		this.bigHeap.Push(num)
	} else if num < little {
		this.littleHeap.Push(num)
	} else {
		this.bigHeap.Push(num)
	}
	fmt.Println("bigHeap2:", this.bigHeap)
	for this.bigHeap.Len() > this.littleHeap.Len()+1 {
		big := this.bigHeap.Pop()
		fmt.Println("big:", big)
		this.littleHeap.Push(big)
	}

	//if this.littleHeap.Len() >= this.bigHeap.Len() {
	//	this.bigHeap.Push(num)
	//} else {
	//	this.littleHeap.Push(num)
	//}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.bigHeap.Len() > this.littleHeap.Len() {
		value := this.bigHeap.Peek()
		return float64(value.(int))
	}
	left, right := this.bigHeap.Peek(), this.littleHeap.Peek()
	return float64(left.(int) + right.(int))
}
