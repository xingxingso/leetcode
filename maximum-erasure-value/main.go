/*
Package maximum_erasure_value
https://leetcode-cn.com/problems/maximum-erasure-value/

1695. 删除子数组的最大得分

给你一个正整数数组 nums ，请你从中删除一个含有 若干不同元素 的子数组。删除子数组的 得分 就是子数组各元素之 和 。

返回 只删除一个 子数组可获得的 最大得分 。

如果数组 b 是数组 a 的一个连续子序列，即如果它等于 a[l],a[l+1],...,a[r] ，那么它就是 a 的一个子数组。

提示：
	1 <= nums.length <= 10^5
	1 <= nums[i] <= 10^4

*/
package maximum_erasure_value

// --- 自己

/*
方法一: 滑动窗口

时间复杂度：
空间复杂度：
*/
func maximumUniqueSubarray(nums []int) int {
	window := make(map[int]int)
	ans := 0
	sum := 0
	for l, r := -1, 0; r < len(nums); r++ {
		if pos, ok := window[nums[r]]; ok && l < pos {
			for l < pos {
				l++
				sum -= nums[l]
			}
		}
		window[nums[r]] = r
		sum += nums[r]
		if sum > ans {
			ans = sum
		}
	}

	return ans
}
