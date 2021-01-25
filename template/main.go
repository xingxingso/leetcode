/*
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

4. 寻找两个正序数组的中位数
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

提示：
nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/
package name

// --- 自己

/*
方法一：二分查找

时间复杂度：O(log(m+n))，其中 m 和 n 分别是数组 nums1 和 nums2 的长度。
	初始时有 k=(m+n)/2 或 k=(m+n)/2+1，每一轮循环可以将查找范围减少一半，因此时间复杂度是 O(log(m+n))。
空间复杂度：O(1)。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}

// --- 官方

/*
方法一：二分查找

时间复杂度：O(log(m+n))，其中 m 和 n 分别是数组 nums1 和 nums2 的长度。
	初始时有 k=(m+n)/2 或 k=(m+n)/2+1，每一轮循环可以将查找范围减少一半，因此时间复杂度是 O(log(m+n))。
空间复杂度：O(1)。
*/
func findMedianSortedArraysO1(nums1 []int, nums2 []int) float64 {
	return 0
}
