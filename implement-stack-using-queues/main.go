/*
Package implement_stack_using_queues
https://leetcode-cn.com/problems/implement-stack-using-queues/

225. 用队列实现栈

请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。

实现 MyStack 类：
	void push(int x) 将元素 x 压入栈顶。
	int pop() 移除并返回栈顶元素。
	int top() 返回栈顶元素。
	boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。

注意：
	你只能使用队列的基本操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
	你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。

提示：
	1 <= x <= 9
	最多调用100 次 push、pop、top 和 empty
	每次调用 pop 和 top 都保证栈不为空

进阶：
	你能否实现每种操作的均摊时间复杂度为 O(1) 的栈？
	换句话说，执行 n 个操作的总时间复杂度 O(n) ，尽管其中某个操作可能需要比其他操作更长的时间。
	你可以使用两个以上的队列。

*/
package implement_stack_using_queues

//type MyStack struct {}
//func Constructor() MyStack {}
//func (this *MyStack) Push(x int)  {}
//func (this *MyStack) Pop() int {}
//func (this *MyStack) Top() int {}
//func (this *MyStack) Empty() bool {}

/**
 * Your MyStackS1 object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/

type simpleQueue struct {
	data []int
}

func (q *simpleQueue) Push(val int) {
	q.data = append(q.data, val)
}

func (q *simpleQueue) Pop() int {
	n := q.data[0]
	q.data = q.data[1:]
	return n
}

func (q *simpleQueue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *simpleQueue) size() int {
	return len(q.data)
}

func (q *simpleQueue) Peek() int {
	return q.data[0]
}

type MyStackS1 struct {
	store simpleQueue
	cur   simpleQueue
	size  int
}

func ConstructorS1() MyStackS1 {
	return MyStackS1{}
}

func (s *MyStackS1) Push(x int) {
	s.cur.Push(x)
	s.size++
	return
}

func (s *MyStackS1) Pop() int {
	for tempSize := s.size; tempSize > 1; tempSize-- {
		s.store.Push(s.cur.Pop())
	}
	s.size--
	n := s.cur.Pop()
	s.cur, s.store = s.store, s.cur
	return n
}

func (s *MyStackS1) Top() int {
	for tempSize := s.size; tempSize > 1; tempSize-- {
		s.store.Push(s.cur.Pop())
	}
	n := s.cur.Pop()
	s.store.Push(n)
	s.cur, s.store = s.store, s.cur
	return n
}

func (s *MyStackS1) Empty() bool {
	return s.size == 0
}
