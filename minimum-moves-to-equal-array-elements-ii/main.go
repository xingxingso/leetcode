/*
Package minimum_moves_to_equal_array_elements_ii
https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/

462. 最少移动次数使数组元素相等 II

给定一个非空整数数组，找到使所有数组元素相等所需的最小移动数，其中每次移动可将选定的一个元素加1或减1。 您可以假设数组的长度最多为10000。

*/
package minimum_moves_to_equal_array_elements_ii

import "sort"

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func minMoves2(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	ans := 0
	for l < r {
		ans += nums[r] - nums[l]
		l++
		r--
	}
	return ans
}
