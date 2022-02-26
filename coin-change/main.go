/*
Package coin_change
https://leetcode-cn.com/problems/coin-change/

322. 零钱兑换

给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
如果没有任何一种硬币组合能组成总金额，返回 -1。
你可以认为每种硬币的数量是无限的。

提示：
	1 <= coins.length <= 12
	1 <= coins[i] <= 2^31 - 1
	0 <= amount <= 104

*/
package coin_change

import (
	"math"
)

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func coinChange(coins []int, amount int) int {
	// 状态转移方程
	// i种硬币，第i种硬币选j个, 凑出k的金额
	// 初始状态
	// dp[0][j] = -1, dp[i][0] = 0
	// 1. 单个面值大于金额
	// dp[i][j] = dp[i-1][j]
	// 2. 单个面值小于等于金额, k 为 1 到 j/coins[i] 之间
	// dp[i][j] = dp[i-1][j-k*coins[i]]

	dp := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			dp[i][j] = -1
			t := j / coins[i]
			if i == 0 {
				if j%coins[i] == 0 {
					dp[i][j] = t
				}
				continue
			}

			for k := t; k >= 0; k-- {
				n := dp[i-1][j-k*coins[i]]
				if n != -1 {
					if dp[i][j] == -1 {
						dp[i][j] = n + k
					} else {
						dp[i][j] = min(dp[i][j], n+k)
					}
				}
			}
		}
	}
	return dp[len(coins)-1][amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func coinChangeS2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := 0; j < len(coins); j++ {
			if i == 0 {
				dp[i] = 0
				continue
			}
			if coins[j] > i {
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// --- 官方

/*
方法一: 记忆化搜索

时间复杂度：O(Sn)，其中 S 是金额，n 是面额数。我们一共需要计算 S 个状态的答案，且每个状态 F(S) 由于上面的记忆化的措施只计算了一次，
	而计算一个状态的答案需要枚举 n 个面额值，所以一共需要 O(Sn) 的时间复杂度。
空间复杂度：O(S)，我们需要额外开一个长为 S 的数组来存储计算出来的答案 F(S)。
*/
func coinChangeO1(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	return coinsChangeO1(coins, amount, make([]int, amount))
}
func coinsChangeO1(coins []int, rem int, count []int) int {
	if rem < 0 {
		return -1
	}
	if rem == 0 {
		return 0
	}
	if count[rem-1] != 0 {
		return count[rem-1]
	}
	min := math.MaxInt64
	for _, coin := range coins {
		res := coinsChangeO1(coins, rem-coin, count)
		if res >= 0 && res < min {
			min = 1 + res
		}
	}
	if min == math.MaxInt64 {
		count[rem-1] = -1
	} else {
		count[rem-1] = min
	}
	return count[rem-1]
}

/*
方法二: 动态规划

时间复杂度：O(Sn)，其中 S 是金额，n 是面额数。我们一共需要计算 O(S) 个状态，S 为题目所给的总金额。
	对于每个状态，每次需要枚举 n 个面额来转移状态，所以一共需要 O(Sn) 的时间复杂度。
空间复杂度：O(S)。数组 dp 需要开长度为总金额 S 的空间。
*/
func coinChangeO2(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = max
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		dp[amount] = -1
	}

	return dp[amount]
}
