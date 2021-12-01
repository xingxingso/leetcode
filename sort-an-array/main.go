/*
Package sort_an_array
https://leetcode-cn.com/problems/sort-an-array/

912. 排序数组

给你一个整数数组 nums，请你将该数组升序排列。

提示：
	1. 1 <= nums.length <= 50000
	2. -50000 <= nums[i] <= 50000

*/
package sort_an_array

// --- 自己

/*
方法一: 快速排序

时间复杂度：
空间复杂度：
*/
func sortArray(nums []int) []int {
	var partition func(lo, hi int) int
	partition = func(lo, hi int) int {
		v := nums[lo]
		i, j := lo, hi+1
		for {
			for i++; nums[i] < v && i < hi; i++ {
			}
			for j--; nums[j] > v && j > lo; j-- {
			}
			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[lo], nums[j] = nums[j], nums[lo]
		return j
	}

	var sort func(lo, hi int)
	sort = func(lo, hi int) {
		if lo >= hi {
			return
		}
		p := partition(lo, hi)
		sort(lo, p-1)
		sort(p+1, hi)
	}
	sort(0, len(nums)-1)
	return nums
}
