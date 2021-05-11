/*
https://leetcode-cn.com/problems/next-greater-element-ii/

503. 下一个更大元素 II

给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，
这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

注意:
	输入数组的长度不会超过 10000。

*/
package next_greater_element_ii

// --- 自己

/*
方法一: 单调栈

时间复杂度：
空间复杂度：
*/
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	stack := NewStackSlice()
	res := make([]int, len(nums))
	for i := n - 1; i >= 0; i-- {
		res[i] = -1

		for !stack.IsEmpty() && stack.Peek() <= nums[i] {
			stack.Pop()
		}

		if stack.IsEmpty() {
			for j := i + 1; j%n != i; j++ {
				if nums[j%n] > nums[i] {
					stack.Push(nums[j%n])
					break
				}
			}
		}

		if !stack.IsEmpty() {
			res[i] = stack.Peek()
		}
		stack.Push(nums[i])
	}

	return res
}

type StackSlice []int

func NewStackSlice() *StackSlice {
	return &StackSlice{}
}

func (s *StackSlice) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StackSlice) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return val
}

func (s *StackSlice) Push(val int) {
	*s = append(*s, val)
}

func (s *StackSlice) Peek() int {
	if s.IsEmpty() {
		return 0
	}

	return (*s)[len(*s)-1]
}

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247487704&idx=1&sn=eb9ac24c644aa0950638c9b20384e982&chksm=9bd7eed0aca067c6b4424c40b7f234c815f83edfbb5efc9f51581335f110e9577114a528f3ec&scene=21#wechat_redirect

/*
方法一: 单调栈

时间复杂度：
空间复杂度：
*/
func nextGreaterElementsP1(nums []int) []int {
	n := len(nums)
	stack := NewStackSlice()
	res := make([]int, len(nums))
	for i := 2*n - 1; i >= 0; i-- {
		res[i%n] = -1
		for !stack.IsEmpty() && stack.Peek() <= nums[i%n] {
			stack.Pop()
		}
		if !stack.IsEmpty() {
			res[i%n] = stack.Peek()
		}
		stack.Push(nums[i%n])
	}
	return res
}
