/*
Package find_minimum_in_rotated_sorted_array_ii
https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

154. 寻找旋转排序数组中的最小值 II

已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。
例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
	- 若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
	- 若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]

注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素

提示：
	n == nums.length
	1 <= n <= 5000
	-5000 <= nums[i] <= 5000
	nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转

进阶：
	这道题是 寻找旋转排序数组中的最小值 的延伸题目。
	允许重复会影响算法的时间复杂度吗？会如何影响，为什么？

*/
package find_minimum_in_rotated_sorted_array_ii

// --- 官方

/*
方法一: 二分查找

时间复杂度：
空间复杂度：
*/
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] > nums[high] {
			low = mid + 1
		} else if nums[mid] < nums[high] {
			high = mid
		} else {
			high--
		}
	}
	return nums[low]
}
