/*
Package best_time_to_buy_and_sell_stock_iii
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

123. 买卖股票的最佳时机 III

给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

注意：
	你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

提示：
	1 <= prices.length <= 10^5
	0 <= prices[i] <= 10^5

*/
package best_time_to_buy_and_sell_stock_iii

import "math"

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxK := 2
	dp := make([][3][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		for k := 1; k <= maxK; k++ {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}

			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][maxK][0]
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
func maxProfitS1(prices []int) int {
	sell1, sell2 := 0, 0
	buy1, buy2 := math.MinInt64, math.MinInt64
	for _, price := range prices {
		sell2 = max(sell2, buy2+price)
		buy2 = max(buy2, sell1-price)
		sell1 = max(sell1, buy1+price)
		buy1 = max(buy1, -price)
	}
	return sell2
}
