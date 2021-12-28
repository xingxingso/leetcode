/*
Package trapping_rain_water
https://leetcode-cn.com/problems/trapping-rain-water/

42. 接雨水

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

提示：

	n == height.length
	0 <= n <= 3 * 10^4
	0 <= height[i] <= 10^5

*/
package trapping_rain_water

// --- 官方

/*
方法一: 暴力

时间复杂度： O(n^2)。数组中的每个元素都需要向左向右扫描。
空间复杂度： O(1) 的额外空间。
*/
func trapO1(height []int) int {
	ans := 0

	for i := 1; i < len(height)-1; i++ {
		maxLeft, maxRight := 0, 0
		// Search the left part for max bar size
		for j := i; j >= 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}
		// Search the right part for max bar size
		for j := i; j < len(height); j++ {
			maxRight = max(maxRight, height[j])
		}
		ans += min(maxLeft, maxRight) - height[i]
	}

	return ans
}

/*
方法二: 动态规划

时间复杂度： 时间复杂度：O(n)。
	存储最大高度数组，需要两次遍历，每次 O(n)。
	最终使用存储的数据更新 ans ，O(n)。
空间复杂度：O(n) 额外空间。
	使用了额外的 O(n) 空间用来放置 left_max 和 right_max 数组。
*/
func trapO2(height []int) int {
	if len(height) == 0 {
		return 0
	}

	ans := 0
	size := len(height)
	leftMax, rightMax := make([]int, size), make([]int, size)
	leftMax[0] = height[0]
	for i := 1; i < size; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	for i := 1; i < size-1; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}

	return ans
}

/*
方法三: 栈的应用

时间复杂度：O(n)。
	单次遍历 O(n) ，每个条形块最多访问两次（由于栈的弹入和弹出），并且弹入和弹出栈都是 O(1) 的。
空间复杂度：O(n)。 栈最多在阶梯型或平坦型条形块结构中占用 O(n) 的空间。
*/
func trapO3(height []int) int {
	ans := 0
	stack := NewStack()
	for current := 0; current < len(height); current++ {
		for !stack.IsEmpty() && height[current] > height[stack.Peek()] {
			top := stack.Pop()
			if stack.IsEmpty() {
				break
			}
			distance := current - stack.Peek() - 1
			boundedHeight := min(height[current], height[stack.Peek()]) - height[top]
			ans += distance * boundedHeight
		}
		stack.Push(current)
	}
	return ans
}

/*
方法四：使用双指针

时间复杂度：O(n)。单次遍历的时间O(n)。
空间复杂度：O(1) 的额外空间。left, right, left_max 和 right_max 只需要常数的空间。
*/
func trapO4(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type StackSlice struct {
	stack []int
}

func NewStack() *StackSlice {
	return &StackSlice{}
}

func (s *StackSlice) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *StackSlice) Pop() int {
	if s.IsEmpty() {
		return 0
	}

	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]

	return val
}

func (s *StackSlice) Push(val int) {
	s.stack = append(s.stack, val)
}

func (s *StackSlice) Peek() int {
	if s.IsEmpty() {
		return 0
	}

	return s.stack[len(s.stack)-1]
}
