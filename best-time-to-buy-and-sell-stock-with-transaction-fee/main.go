/*
Package best_time_to_buy_and_sell_stock_with_transaction_fee
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

714. 买卖股票的最佳时机含手续费

给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

注意：
	这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

提示：
	1 <= prices.length <= 5 * 10^4
	1 <= prices[i] < 5 * 10^4
	0 <= fee < 5 * 10^4

*/
package best_time_to_buy_and_sell_stock_with_transaction_fee

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func maxProfit(prices []int, fee int) int {
	// 0:有股票,1:无股票, 无卖出,2:无股票, 有卖出
	dp := make([][3]int, len(prices))
	dp[0] = [3]int{-prices[0], 0, 0}
	for i := 1; i < len(prices); i++ {
		dp[i] = [3]int{}
		dp[i][0] = maxN(dp[i-1][0], dp[i-1][1]-prices[i], dp[i-1][2]-prices[i])
		dp[i][1] = maxN(dp[i-1][1], dp[i-1][2])
		dp[i][2] = dp[i-1][0] + prices[i] - fee
	}
	//fmt.Println(dp)
	return maxN(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

func maxN(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}
