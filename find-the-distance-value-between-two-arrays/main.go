/*
Package find_the_distance_value_between_two_arrays
https://leetcode-cn.com/problems/find-the-distance-value-between-two-arrays/

1385. 两个数组间的距离值

给你两个整数数组 arr1 ， arr2 和一个整数 d ，请你返回两个数组之间的 距离值 。

「距离值」 定义为符合此距离要求的元素数目：对于元素 arr1[i] ，不存在任何元素 arr2[j] 满足 |arr1[i]-arr2[j]| <= d 。

提示：
	1 <= arr1.length, arr2.length <= 500
	-10^3 <= arr1[i], arr2[j] <= 10^3
	0 <= d <= 100

*/
package find_the_distance_value_between_two_arrays

import (
	"sort"
)

// --- 自己

/*
方法一: 二分查找

时间复杂度：
空间复杂度：
*/
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	ans := 0
	for i := 0; i < len(arr1); i++ {
		d2 := findMinDistance(arr2, arr1[i])
		if d2 > d {
			ans++
		}
	}
	return ans
}

func findMinDistance(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2
		if arr[m] == target {
			return 0
		}
		if arr[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	// fmt.Println(l, r)
	if r < 0 {
		r = 0
	}
	if l >= len(arr) {
		l = len(arr) - 1
	}
	ans := min(abs(arr[l]-target), abs(arr[r]-target))
	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a // a > max
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
