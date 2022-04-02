/*
Package search_insert_position
https://leetcode-cn.com/problems/search-insert-position/

35. 搜索插入位置

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

提示:
	1 <= nums.length <= 10^4
	-10^4 <= nums[i] <= 10^4
	nums 为 无重复元素 的 升序 排列数组
	-10^4 <= target <= 10^4

*/
package search_insert_position

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
