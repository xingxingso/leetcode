/*
Package maximum_product_subarray
https://leetcode-cn.com/problems/maximum-product-subarray/

152. 乘积最大子数组

给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个 32-位 整数。

子数组 是数组的连续子序列。

提示:
	1 <= nums.length <= 2 * 10^4
	-10 <= nums[i] <= 10
	nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数

*/
package maximum_product_subarray

import "math"

// --- 自己

/*
方法一: 动态规划
	// 超出内存限制
时间复杂度：
空间复杂度：
*/
func maxProductS1(nums []int) int {
	dp := make([][]int, len(nums))
	ans := math.MinInt64
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		for j := i; j < len(nums); j++ {
			if i == j {
				dp[i][j] = nums[i]
			} else {
				dp[i][j] = dp[i][j-1] * nums[j]
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func maxProductS2(nums []int) int {
	dp := make([]int, len(nums))
	ans := math.MinInt64
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if i == j {
				dp[j] = nums[i]
			} else {
				dp[j] = dp[j-1] * nums[j]
			}
			if dp[j] > ans {
				ans = dp[j]
			}
		}
	}
	return ans
}

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func maxProductS3(nums []int) int {
	ans := math.MinInt64
	var cur, pre int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if i == j {
				cur = nums[i]
			} else {
				cur = pre * nums[j]
			}
			if cur > ans {
				ans = cur
			}
			pre = cur
		}
	}
	return ans
}

// --- 官方

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func maxProductO1(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(ans, maxF)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
