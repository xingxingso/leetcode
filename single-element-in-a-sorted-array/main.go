/*
Package single_element_in_a_sorted_array
https://leetcode-cn.com/problems/single-element-in-a-sorted-array/

540. 有序数组中的单一元素

给定一个只包含整数的有序数组，每个元素都会出现两次，唯有一个数只会出现一次，找出这个数。

提示:
	1 <= nums.length <= 10^5
	0 <= nums[i] <= 10^5

进阶:
	采用的方案可以在 O(log n) 时间复杂度和 O(1) 空间复杂度中运行吗？

*/
package single_element_in_a_sorted_array

// --- 自己

/*
方法一: 二分查找

时间复杂度：
空间复杂度：
*/
func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		if ((high-low)/2)%2 == 0 {
			if nums[mid-1] == nums[mid] {
				high = mid - 2
			} else if nums[mid+1] == nums[mid] {
				low = mid + 2
			} else {
				return nums[mid]
			}
		} else {
			if nums[mid-1] == nums[mid] {
				low = mid + 1
			} else if nums[mid+1] == nums[mid] {
				high = mid - 1
			} else {
				return nums[mid]
			}
		}
	}

	return nums[low]
}
