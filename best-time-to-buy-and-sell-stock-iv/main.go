/*
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/

188. 买卖股票的最佳时机 IV

给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

提示：
	0 <= k <= 100
	0 <= prices.length <= 1000
	0 <= prices[i] <= 1000

*/
package best_time_to_buy_and_sell_stock_iv

import "math"

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func maxProfit(k int, prices []int) int {
	if k > len(prices)/2 { // k 失效
		return maxProfitS1(prices)
	}

	dp := make([][][2]int, len(prices))

	for i := 0; i < len(prices); i++ {
		dp[i] = make([][2]int, k+1)
		for j := k; j >= 1; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			if j == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}

			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[len(prices)-1][k][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfitS1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	free := 0
	hold := -prices[0]

	for i := 1; i < len(prices); i++ {
		free = max(free, hold+prices[i])
		hold = max(hold, free-prices[i])
	}

	return free
}

// --- 官方

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func maxProfitO1(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k = min(k, n/2)
	buy := make([][]int, n)
	sell := make([][]int, n)
	for i := range buy {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64 / 2
		sell[0][i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[i][0] = maxN(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = maxN(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = maxN(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}
	return maxN(sell[n-1]...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxN(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

/*
方法一: 动态规划优化

时间复杂度：O(n min(n,k))，其中 n 是数组 prices 的大小，即我们使用二重循环进行动态规划需要的时间。
空间复杂度：O(n min(n,k)) 或 O(min(n,k))，取决于我们使用二维数组还是一维数组进行动态规划。
*/
func maxProfitO2(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k = min(k, n/2)
	buy := make([]int, k+1)
	sell := make([]int, k+1)
	buy[0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[i] = math.MinInt64 / 2
		sell[i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[0] = maxN(buy[0], sell[0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[j] = maxN(buy[j], sell[j]-prices[i])
			sell[j] = maxN(sell[j], buy[j-1]+prices[i])
		}
	}
	return maxN(sell...)
}
