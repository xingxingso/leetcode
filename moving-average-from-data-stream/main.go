/*
https://leetcode-cn.com/problems/moving-average-from-data-stream/

346. 数据流中的移动平均值

给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算其所有整数的移动平均值。

*/
package moving_average_from_data_stream

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
type MovingAverage struct {
	head, tail    *ListNode
	len, cap, sum int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	node := &ListNode{}
	return MovingAverage{
		cap:  size,
		len:  0,
		sum:  0,
		head: node,
		tail: node,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	node := &ListNode{
		Val: val,
	}
	this.tail.Next = node
	this.tail = node
	this.sum += val
	if this.len < this.cap {
		this.len++
	} else {
		this.sum -= this.head.Next.Val
		this.head.Next = this.head.Next.Next
	}

	return float64(this.sum) / float64(this.len)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
