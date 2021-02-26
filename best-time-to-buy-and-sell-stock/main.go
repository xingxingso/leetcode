/*
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

121. 买卖股票的最佳时机

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

提示：
	1 <= prices.length <= 105
	0 <= prices[i] <= 104

*/
package best_time_to_buy_and_sell_stock

import "math"

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	dp[0] = make([]int, 2)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i] = make([]int, 2)
		// 状态转移方程: 0 未持有 1 持有
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 第i天卖出的时候最大收益
		dp[i][1] = max(dp[i-1][1], -prices[i])           // 第i天买入的时候最大收益 只有1次交易机会 所以是 -price[i]
	}

	return dp[len(prices)-1][0] // 返回最后卖出的最大收益
}

/*
方法一: 动态规划优化

时间复杂度：
空间复杂度：
*/
func maxProfitS1(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	sell := 0
	buy := -prices[0]
	for i := 1; i < len(prices); i++ {
		sell = max(sell, buy+prices[i]) // 第i天卖出的时候最大收益
		buy = max(buy, -prices[i])      // 第i天买入的时候最大收益
	}
	return sell // 返回最后卖出的最大收益
}

/*
方法0: 暴力解法
	超出时限

时间复杂度：
空间复杂度：
*/
func maxProfitS0(prices []int) int {
	ans := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			ans = max(ans, prices[j]-prices[i])
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// --- 官方

/*
方法二: 一次遍历

时间复杂度：O(n)，只需要遍历一次。
空间复杂度：O(1)，只使用了常数个变量。
*/
func maxProfitO1(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice { // 股票在跌, 不用买
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
