/*
Package find_first_and_last_position_of_element_in_sorted_array
https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

34. 在排序数组中查找元素的第一个和最后一个位置

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：
	你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？

提示：
	0 <= nums.length <= 10^5
	-10^9 <= nums[i] <= 10^9
	nums 是一个非递减数组
	-10^9 <= target <= 10^9

*/
package find_first_and_last_position_of_element_in_sorted_array

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func searchRange(nums []int, target int) []int {
	left := searchLeft(nums, 0, len(nums)-1, target)
	right := searchRight(nums, 0, len(nums)-1, target)
	return []int{left, right}
}

func searchLeft(nums []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)>>1
		if target <= nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target { // 不是找下标了
		return -1
	}
	return left
}

func searchRight(nums []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target { // 不是找下标了
		return -1
	}
	return right
}
