/*
Package implement_queue_using_stacks
https://leetcode-cn.com/problems/implement-queue-using-stacks/

232. 用栈实现队列

请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：
	void push(int x) 将元素 x 推到队列的末尾
	int pop() 从队列的开头移除并返回元素
	int peek() 返回队列开头的元素
	boolean empty() 如果队列为空，返回 true ；否则，返回 false

说明：
	你只能使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
	你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。


进阶：
	你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。

*/
package implement_queue_using_stacks

//type MyQueue struct {}
//func Constructor() MyQueue {}
//func (this *MyQueue) Push(x int)  {}
//func (this *MyQueue) Pop() int {}
//func (this *MyQueue) Peek() int {}
//func (this *MyQueue) Empty() bool {}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type StackList struct {
	head *ListNode
}

func NewStackList() *StackList {
	return &StackList{
		head: &ListNode{},
	}
}

func (stack *StackList) IsEmpty() bool {
	return stack.head.Next == nil
}

func (stack *StackList) Pop() int {
	if stack.head.Next == nil {
		return -1
	}

	val := stack.head.Next.Val
	stack.head.Next = stack.head.Next.Next
	return val
}

func (stack *StackList) Push(val int) {
	node := &ListNode{
		Val: val,
	}

	node.Next = stack.head.Next
	stack.head.Next = node
}

func (stack *StackList) Peek() int {
	if stack.head.Next == nil {
		return -1
	}

	return stack.head.Next.Val
}

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/

type MyQueue struct {
	in  *StackList
	out *StackList
}

func Constructor() MyQueue {
	return MyQueue{
		in:  NewStackList(),
		out: NewStackList(),
	}
}

func (q *MyQueue) Push(x int) {
	q.in.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.out.IsEmpty() {
		for !q.in.IsEmpty() {
			q.out.Push(q.in.Pop())
		}
	}

	return q.out.Pop()
}

func (q *MyQueue) Peek() int {
	if q.out.IsEmpty() {
		for !q.in.IsEmpty() {
			q.out.Push(q.in.Pop())
		}
	}

	return q.out.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.in.IsEmpty() && q.out.IsEmpty()
}
