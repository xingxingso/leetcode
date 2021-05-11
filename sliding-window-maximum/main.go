/*
https://leetcode-cn.com/problems/sliding-window-maximum/

239. 滑动窗口最大值

给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

提示：
	1 <= nums.length <= 10^5
	-10^4 <= nums[i] <= 10^4
	1 <= k <= nums.length

*/
package sliding_window_maximum

import "fmt"

// --- 自己

/*
方法一: 单调队列
	这个有点像队列和栈的结合, 出队 入栈

时间复杂度：
空间复杂度：
*/
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	queue := NewQueueInt()
	for i := 0; i < n; i++ {
		for !queue.IsEmpty() {
			if _, val := queue.tailPeek(); nums[i] < val {
				break
			}
			queue.tailPop()
		}
		queue.Push(i, nums[i])

		for key, _ := queue.Peek(); !queue.IsEmpty() && key <= i-k; key, _ = queue.Peek() {
			queue.Pop()
		}

		if !queue.IsEmpty() && i+1 >= k {
			_, ans[i+1-k] = queue.Peek()
		}
	}
	return ans
}

type DoubleList struct {
	Key, Val  int
	Pre, Next *DoubleList
}

type QueueInt struct {
	head *DoubleList
	tail *DoubleList
}

func NewQueueInt() *QueueInt {
	head, tail := &DoubleList{}, &DoubleList{}
	head.Next = tail
	tail.Pre = head
	return &QueueInt{
		head: head,
		tail: tail,
	}
}

func (q *QueueInt) IsEmpty() bool { return q.head.Next == q.tail }

func (q *QueueInt) Pop() int {
	if q.IsEmpty() {
		return -1
	}
	node := q.head.Next
	q.head.Next = node.Next
	node.Next.Pre = q.head
	return node.Val
}

func (q *QueueInt) tailPop() int {
	if q.IsEmpty() {
		return -1
	}
	node := q.tail.Pre
	q.tail.Pre = node.Pre
	node.Pre.Next = q.tail
	return node.Val
}

func (q *QueueInt) Push(key, val int) {
	node := &DoubleList{
		Key: key,
		Val: val,
	}
	node.Next = q.tail
	node.Pre = q.tail.Pre
	node.Pre.Next = node
	q.tail.Pre = node
}

func (q *QueueInt) Peek() (int, int) {
	if q.IsEmpty() {
		return -1, -1
	}

	return q.head.Next.Key, q.head.Next.Val
}

func (q *QueueInt) tailPeek() (int, int) {
	if q.IsEmpty() {
		return -1, -1
	}
	return q.tail.Pre.Key, q.tail.Pre.Val
}

// --- 官方

/*
方法二: 单调队列

时间复杂度：O(n)，其中 n 是数组 nums 的长度。每一个下标恰好被放入队列一次，并且最多被弹出队列一次，因此时间复杂度为 O(n)。
空间复杂度：O(k)。与方法一不同的是，在方法二中我们使用的数据结构是双向的，
	因此「不断从队首弹出元素」保证了队列中最多不会有超过 k+1 个元素，因此队列使用的空间为 O(k)。
*/
func maxSlidingWindowO1(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

func printList(node *DoubleList) {
	for node != nil {
		fmt.Printf("key:%d,val:%d,", node.Key, node.Val)
		node = node.Next
	}
	fmt.Println()
}
