/*
Package minimum_moves_to_equal_array_elements
https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/

453. 最小操作次数使数组元素相等

给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。

提示：
	n == nums.length
	1 <= nums.length <= 10^5
	-10^9 <= nums[i] <= 10^9
	答案保证符合 32-bit 整数

*/
package minimum_moves_to_equal_array_elements

// --- 自己

/*
方法一:
	// 找到最小值，所有元素与它的差值和即是问题解

时间复杂度：
空间复杂度：
*/
func minMoves(nums []int) int {
	ans := 0
	//sort.Ints(nums) // 可以不排序
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		ans += nums[i] - min
	}

	return ans
}

/*
方法一: 一轮循环
	// 找到最小值，所有元素与它的差值和即是问题解

时间复杂度：
空间复杂度：
*/
func minMovesS2(nums []int) int {
	ans := 0
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			ans += (min - nums[i]) * i
			min = nums[i]
		} else {
			ans += nums[i] - min
		}
	}
	return ans
}
