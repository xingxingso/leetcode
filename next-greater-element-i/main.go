/*
Package next_greater_element_i
https://leetcode-cn.com/problems/next-greater-element-i/

496. 下一个更大元素 I

给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

提示：
	1 <= nums1.length <= nums2.length <= 1000
	0 <= nums1[i], nums2[i] <= 10^4
	nums1和nums2中所有整数 互不相同
	nums1 中的所有整数同样出现在 nums2 中

进阶：
	你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？

*/
package next_greater_element_i

// --- 自己

/*
方法一: 单调栈

时间复杂度：
空间复杂度：
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := NewStackSlice()
	mark := map[int]int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		mark[nums2[i]] = -1
		for !stack.IsEmpty() && nums2[i] >= stack.Peek() {
			stack.Pop()
		}
		if !stack.IsEmpty() {
			mark[nums2[i]] = stack.Peek()
		}
		stack.Push(nums2[i])
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = mark[nums1[i]]
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
