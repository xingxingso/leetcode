/*
Package longest_harmonious_subsequence
https://leetcode-cn.com/problems/longest-harmonious-subsequence/

594. 最长和谐子序列

和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。

现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。

数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。

提示：
	1 <= nums.length <= 2 * 10^4
	-10^9 <= nums[i] <= 10^9

*/
package longest_harmonious_subsequence

import "sort"

// --- 自己

/*
方法一: 排序

时间复杂度：
空间复杂度：
*/
func findLHS(nums []int) int {
	sort.Ints(nums)
	ans := 0
	pre := 0
	for i := 1; i < len(nums); {
		if nums[i]-nums[pre] == 0 {
			i++
		} else if nums[i]-nums[pre] == 1 {
			ans = max(ans, i-pre+1)
			i++
		} else {
			pre++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
