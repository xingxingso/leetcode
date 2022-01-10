/*
Package binary_search
https://leetcode-cn.com/problems/binary-search/

704. 二分查找

给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

提示：
	1. 你可以假设 nums 中的所有元素是不重复的。
	2. n 将在 [1, 10000]之间。
	3. nums 的每个元素都将在 [-9999, 9999]之间。
*/
package binary_search

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
