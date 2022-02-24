/*
Package merge_sorted_array
https://leetcode-cn.com/problems/merge-sorted-array/

88. 合并两个有序数组

给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：
	最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。
	为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

提示：
	nums1.length == m + n
	nums2.length == n
	0 <= m, n <= 200
	1 <= m + n <= 200
	-10^9 <= nums1[i], nums2[j] <= 10^9

进阶：
	你可以设计实现一个时间复杂度为 O(m + n) 的算法解决此问题吗？

*/
package merge_sorted_array

// --- 自己

/*
方法一: 双指针

时间复杂度：
空间复杂度：
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, k := m-1, n-1, m+n-1; k >= 0; k-- {
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				nums1[k] = nums1[i]
				i--
			} else {
				nums1[k] = nums2[j]
				j--
			}
		} else if i >= 0 {
			nums1[k] = nums1[i]
			i--
		} else if j >= 0 {
			nums1[k] = nums2[j]
			j--
		}
	}
	return
}

/*
方法二: 双指针优化

时间复杂度：
空间复杂度：
*/
func mergeS2(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1
	m--
	n--
	for ; m >= 0 && n >= 0; pos-- {
		if nums1[m] > nums2[n] {
			nums1[pos] = nums1[m]
			m--
		} else {
			nums1[pos] = nums2[n]
			n--
		}
	}
	for ; n >= 0; pos-- {
		nums1[pos] = nums2[n]
		n--
	}
	return
}

/*
方法一: 双指针
	倒序调整

时间复杂度：
空间复杂度：
*/
func mergeS3(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	for i := len(nums1) - 1; i >= 0; i-- {
		if m >= 0 && n >= 0 {
			if nums1[m] > nums2[n] {
				nums1[i] = nums1[m]
				m--
			} else {
				nums1[i] = nums2[n]
				n--
			}
		} else if m >= 0 {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
	}
}
