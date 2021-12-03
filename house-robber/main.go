/*
Package house_robber
https://leetcode-cn.com/problems/house-robber/

198. 打家劫舍

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

提示：
	0 <= nums.length <= 100
	0 <= nums[i] <= 400

*/
package house_robber

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([][2]int, len(nums))
	dp[0][0] = 0       // 第1间 不偷
	dp[0][1] = nums[0] // 第1间 偷

	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]) // 上一间偷了 或者 没偷
		dp[i][1] = dp[i-1][0] + nums[i]        // 上一间没偷
	}
	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法一: 动态规划优化

时间复杂度：
空间复杂度：
*/
func robS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	first, second := 0, 0
	maxProfit := 0

	for i := len(nums) - 1; i >= 0; i-- {
		maxProfit = max(first, nums[i]+second)
		second = first
		first = maxProfit
	}
	return maxProfit
}
