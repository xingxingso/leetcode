/*
Package non_decreasing_array
https://leetcode-cn.com/problems/non-decreasing-array/

665. 非递减数列

给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。

我们是这样定义一个非递减数列的： 对于数组中任意的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

提示：
	1 <= n <= 10^4
	- 10^5 <= nums[i] <= 10^5

*/
package non_decreasing_array

import (
	"sort"
)

// --- 官方

/*
方法一: 数组

时间复杂度：
空间复杂度：
*/
func checkPossibilityO1(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			nums[i] = y
			if sort.IntsAreSorted(nums) {
				return true
			}
			nums[i] = x
			nums[i+1] = x
			return sort.IntsAreSorted(nums)
		}
	}
	return true
}

/*
方法一: 优化

时间复杂度：O(n)，其中 n 是数组 nums 的长度。
空间复杂度：O(1)
*/
func checkPossibilityO2(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			cnt++
			if cnt > 1 {
				return false
			}
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
}
