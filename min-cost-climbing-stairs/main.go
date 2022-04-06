/*
Package min_cost_climbing_stairs
https://leetcode-cn.com/problems/min-cost-climbing-stairs/

746. 使用最小花费爬楼梯

给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。

提示：
	2 <= cost.length <= 1000
	0 <= cost[i] <= 999

*/
package min_cost_climbing_stairs

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = cost[0]
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i-1]
	}
	ans := min(dp[len(cost)], dp[len(cost)-1])

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
