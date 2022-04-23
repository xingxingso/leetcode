/*
Package peak_index_in_a_mountain_array
https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/

852. 山脉数组的峰顶索引

符合下列属性的数组 arr 称为 山脉数组 ：
	arr.length >= 3
	存在 i（0 < i < arr.length - 1）使得：
		arr[0] < arr[1] < ... arr[i-1] < arr[i]
		arr[i] > arr[i+1] > ... > arr[arr.length - 1]

给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i 。

提示：
	3 <= arr.length <= 10^4
	0 <= arr[i] <= 10^6
	题目数据保证 arr 是一个山脉数组

进阶：
	很容易想到时间复杂度 O(n) 的解决方案，你可以设计一个 O(log(n)) 的解决方案吗？

*/
package peak_index_in_a_mountain_array

import "sort"

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2
		if arr[m-1] < arr[m] {
			if arr[m+1] < arr[m] {
				return m
			}
			l = m
		} else {
			r = m
		}
	}
	return -1
}

// --- 官方

/*
方法一:

时间复杂度：O(logn)，其中 n 是数组 arr 的长度。我们需要进行二分查找的次数为 O(logn)。
空间复杂度：O(1)。
*/
func peakIndexInMountainArrayO1(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool {
		return arr[i] > arr[i+1]
	})
}
