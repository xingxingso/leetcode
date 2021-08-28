/*
Package max_stack
https://leetcode-cn.com/problems/max-stack/

716. 最大栈

设计一个最大栈数据结构，既支持栈操作，又支持查找栈中最大元素。

实现 MaxStack 类：
	MaxStack() 初始化栈对象
	void push(int x) 将元素 x 压入栈中。
	int pop() 移除栈顶元素并返回这个元素。
	int top() 返回栈顶元素，无需移除。
	int peekMax() 检索并返回栈中最大元素，无需移除。
	int popMax() 检索并返回栈中最大元素，并将其移除。如果有多个最大元素，只要移除 最靠近栈顶 的那个。

提示：
	-107 <= x <= 107
	最多调用 104 次 push、pop、top、peekMax 和 popMax
	调用 pop、top、peekMax 或 popMax 时，栈中 至少存在一个元素

进阶：
	试着设计解决方案：调用 top 方法的时间复杂度为 O(1) ，调用其他方法的时间复杂度为 O(logn) 。

*/
package max_stack

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
type MaxStack struct {
	data []int
}

/** initialize your data structure here. */
func Constructor() MaxStack {
	return MaxStack{}
}

func (this *MaxStack) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MaxStack) Pop() int {
	if len(this.data) < 1 {
		return -1
	}

	val := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return val
}

func (this *MaxStack) Top() int {
	val := this.data[len(this.data)-1]
	return val
}

func (this *MaxStack) PeekMax() int {
	if len(this.data) < 1 {
		return -1
	}

	max := this.data[len(this.data)-1]
	for i := len(this.data) - 2; i >= 0; i-- {
		if this.data[i] > max {
			max = this.data[i]
		}
	}
	return max
}

func (this *MaxStack) PopMax() int {
	if len(this.data) < 1 {
		return -1
	}

	maxI := len(this.data) - 1
	max := this.data[maxI]

	for i := len(this.data) - 2; i >= 0; i-- {
		if this.data[i] > max {
			max = this.data[i]
			maxI = i
		}
	}

	this.data = append(this.data[:maxI], this.data[maxI+1:]...)

	return max
}

// --- 官方
// https://leetcode-cn.com/problems/max-stack/solution/max-stack-by-leetcode/

/*
方法一: 双栈

时间复杂度：O(n)。由于前四个操作的时间复杂度都是 O(1)，而 popMax() 操作在最坏情况下需要将栈中的所有元素全部出栈再入栈，
	时间复杂度为 O(n)。因此总的时间复杂度为 O(n)。
空间复杂度：O(n)。
*/
type MaxStackO1 struct {
	stack    []int
	maxStack []int
}

/** initialize your data structure here. */
func ConstructorO1() MaxStack {
	return MaxStack{}
}

func (this *MaxStackO1) Push(x int) {
	this.stack = append(this.stack, x)
	maxX := x
	if len(this.maxStack) != 0 {
		maxX = this.PeekMax()
	}
	this.maxStack = append(this.maxStack, max(maxX, x))
}

func (this *MaxStackO1) Pop() int {
	v := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.maxStack = this.maxStack[:len(this.maxStack)-1]
	return v
}

func (this *MaxStackO1) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MaxStackO1) PeekMax() int {
	return this.maxStack[len(this.maxStack)-1]
}

func (this *MaxStackO1) PopMax() int {
	max := this.PeekMax()
	var b []int
	for this.Top() != max {
		b = append(b, this.Pop())
	}
	this.Pop()

	if len(b) != 0 {
		this.Push(b[len(b)-1])
		b = b[:len(b)-1]
	}
	return max
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
todo: 待转化
方法二: 双向链表 + 平衡树

时间复杂度：O(logn)。top() 操作的时间复杂度为 O(1)，其余操作的时间复杂度为 O(logn)，因此总的时间复杂度为 O(logn)。
空间复杂度：O(n)。
*/
type MaxStackO2 struct {
	stack    []int
	maxStack []int
}

/** initialize your data structure here. */
func ConstructorO2() MaxStack {
	return MaxStack{}
}

func (this *MaxStackO2) Push(x int) {
	this.stack = append(this.stack, x)
	maxX := x
	if len(this.maxStack) != 0 {
		maxX = this.PeekMax()
	}
	this.maxStack = append(this.maxStack, max(maxX, x))
}

func (this *MaxStackO2) Pop() int {
	v := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.maxStack = this.maxStack[:len(this.maxStack)-1]
	return v
}

func (this *MaxStackO2) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MaxStackO2) PeekMax() int {
	return this.maxStack[len(this.maxStack)-1]
}

func (this *MaxStackO2) PopMax() int {
	max := this.PeekMax()
	var b []int
	for this.Top() != max {
		b = append(b, this.Pop())
	}
	this.Pop()

	if len(b) != 0 {
		this.Push(b[len(b)-1])
		b = b[:len(b)-1]
	}
	return max
}
