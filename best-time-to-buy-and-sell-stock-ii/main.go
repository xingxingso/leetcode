/*
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

122. 买卖股票的最佳时机 II

给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

提示：
	1 <= prices.length <= 3 * 10 ^ 4
	0 <= prices[i] <= 10 ^ 4

*/
package best_time_to_buy_and_sell_stock_ii

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		// dp[i][1] = max(dp[i-1][1], -prices[i]) // 如果只能交易一次
	}

	return dp[len(prices)-1][0]
}

/*
方法一: 动态规划优化

时间复杂度：
空间复杂度：
*/
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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// --- 官方

/*
方法一: 动态规划

时间复杂度：O(n)，其中 n 为数组的长度。一共有 2n 个状态，每次状态转移的时间复杂度为 O(1)，因此时间复杂度为 O(2n)=O(n)。
空间复杂度：O(n)。我们需要开辟 O(n) 空间存储动态规划中的所有状态。如果使用空间优化，空间复杂度可以优化至 O(1)。
*/
func maxProfitO1(prices []int) int {
	// n := len(prices)
	// dp := make([][2]int, n)
	// dp[0][1] = -prices[0]
	// for i := 1; i < n; i++ {
	// 	dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	// 	dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	// }
	// return dp[n-1][0]

	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	return dp0
}

/*
方法二: 贪心

时间复杂度：O(n)，其中 n 为数组的长度。我们只需要遍历一次数组即可。
空间复杂度：O(1)。只需要常数空间存放若干变量。
*/
func maxProfitO2(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}
