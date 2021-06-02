/*
Package move_zeroes
https://leetcode-cn.com/problems/move-zeroes/

283. 移动零

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

说明:
	1. 必须在原数组上操作，不能拷贝额外的数组。
	2. 尽量减少操作次数。

*/
package move_zeroes

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
}

// --- 官方

/*
方法一:

时间复杂度：
空间复杂度：
*/
func moveZeroesO1(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
