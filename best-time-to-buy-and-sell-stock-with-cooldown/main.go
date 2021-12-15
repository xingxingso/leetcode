/*
Package best_time_to_buy_and_sell_stock_with_cooldown
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

309. 最佳买卖股票时机含冷冻期

给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
	- 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
	- 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

*/
package best_time_to_buy_and_sell_stock_with_cooldown

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func maxProfit(prices []int) int {
	//0: 持有，1: 不持有，冷冻，2：不持有，非冷冻
	dp := make([][3]int, len(prices))
	dp[0] = [3]int{-prices[0], 0, 0}
	for i := 1; i < len(prices); i++ {
		dp[i] = [3]int{}
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
