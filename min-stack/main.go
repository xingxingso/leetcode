/*
Package min_stack
https://leetcode-cn.com/problems/min-stack/

155. 最小栈

设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。

提示：
	pop、top 和 getMin 操作总是在 非空栈 上调用。

*/
package min_stack

//type MinStack struct {}
//func Constructor() MinStack {}
//func (this *MinStack) Push(val int)  {}
//func (this *MinStack) Pop()  {}
//func (this *MinStack) Top() int {}
//func (this *MinStack) GetMin() int {}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// --- 自己

/*
方法一: 辅助栈

时间复杂度：
空间复杂度：
*/

type MinStack struct {
	origin []int
	min    []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.origin = append(s.origin, val)
	if len(s.min) == 0 {
		s.min = append(s.min, val)
		return
	}
	if val <= s.min[len(s.min)-1] {
		s.min = append(s.min, val)
	}
	return
}

func (s *MinStack) Pop() {
	top := s.origin[len(s.origin)-1]
	if top == s.min[len(s.min)-1] {
		s.min = s.min[:len(s.min)-1]
	}
	s.origin = s.origin[:len(s.origin)-1]
}

func (s *MinStack) Top() int {
	return s.origin[len(s.origin)-1]
}

func (s *MinStack) GetMin() int {
	return s.min[len(s.min)-1]
}

// --- 他人

/*
方法一:
	不使用额外空间的做法

时间复杂度：
空间复杂度：
*/

type MinStackP1 struct {
	origin []int
	min    int
}

func ConstructorP1() MinStackP1 {
	return MinStackP1{}
}

func (s *MinStackP1) Push(val int) {
	if len(s.origin) == 0 {
		s.origin = append(s.origin, 0)
		s.min = val
		return
	}
	diff := val - s.min
	s.origin = append(s.origin, diff)
	if diff <= 0 {
		s.min = val
	}
}

func (s *MinStackP1) Pop() {
	diff := s.origin[len(s.origin)-1]
	if diff < 0 {
		s.min = s.min - diff
	}
	s.origin = s.origin[:len(s.origin)-1]
}

func (s *MinStackP1) Top() int {
	diff := s.origin[len(s.origin)-1]
	if diff <= 0 {
		return s.min
	}

	return s.origin[len(s.origin)-1] + s.min
}

func (s *MinStackP1) GetMin() int {
	return s.min
}
