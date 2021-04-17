/*
https://leetcode-cn.com/problems/stone-game-iii/

1406. 石子游戏 III

Alice 和 Bob 用几堆石子在做游戏。几堆石子排成一行，每堆石子都对应一个得分，由数组 stoneValue 给出。
Alice 和 Bob 轮流取石子，Alice 总是先开始。在每个玩家的回合中，该玩家可以拿走剩下石子中的的前 1、2 或 3 堆石子 。比赛一直持续到所有石头都被拿走。
每个玩家的最终得分为他所拿到的每堆石子的对应得分之和。每个玩家的初始分数都是 0 。比赛的目标是决出最高分，得分最高的选手将会赢得比赛，比赛也可能会出现平局。
假设 Alice 和 Bob 都采取 最优策略 。如果 Alice 赢了就返回 "Alice" ，Bob 赢了就返回 "Bob"，平局（分数相同）返回 "Tie" 。

提示：
	1 <= values.length <= 50000
	-1000 <= values[i] <= 1000

*/
package stone_game_iii

import (
	"math"
)

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func stoneGameIII(stoneValue []int) string {
	// 剩余 [i:] 取 前j个时候 Alice 和 Bob 的最大差值
	// dp[i][j]
	length := len(stoneValue)
	dp := make([][3]int, length)
	n := 0
	for i := length - 1; i >= 0; i-- {
		n += stoneValue[i]
		for j := 2; j >= 0; j-- {
			dp[i][j] = math.MinInt64
			if i+j < length-1 {
				dp[i][j] = maxN(dp[i][j], sum(stoneValue[i:i+j+1])-maxN(dp[i+j+1][0], dp[i+j+1][1], dp[i+j+1][2]))
			} else {
				dp[i][j] = n
			}
		}
	}
	res := maxN(dp[0][0], dp[0][1], dp[0][2])
	if res > 0 {
		return "Alice"
	} else if res == 0 {
		return "Tie"
	} else {
		return "Bob"
	}
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

func sum(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans += v
	}

	return ans
}

// --- 官方

/*
方法一: 动态规划

时间复杂度：O(N)，其中 N 是数组 values 的长度。
空间复杂度：O(N)。
*/
func stoneGameIIIO1(stoneValue []int) string {
	// dp[i] 表示还剩下第 i,i+1,⋯,n−1 堆石子时，当前玩家最多能从剩下的石子中拿到的石子数目
	length := len(stoneValue)
	dp := make([]int, length+1)
	n := 0
	dp[length] = 0
	for i := length - 1; i >= 0; i-- {
		n += stoneValue[i]
		dp[i] = n - dp[i+1]
		for j := i + 2; j <= i+3 && j <= length; j++ {
			dp[i] = max(dp[i], n-dp[j])
		}
	}
	res := 2*dp[0] - n
	if res > 0 {
		return "Alice"
	} else if res == 0 {
		return "Tie"
	} else {
		return "Bob"
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法二: 动态规划

时间复杂度：O(N)，其中 N 是数组 values 的长度。
空间复杂度：O(N)。
*/
func stoneGameIIIO2(stoneValue []int) string {
	n := len(stoneValue)
	f := make([]int, n+1)
	// 边界情况，当没有石子时，分数为 0
	f[n] = 0
	for i := n - 1; i >= 0; i-- {
		f[i] = math.MinInt64
		pre := 0
		for j := i + 1; j <= i+3 && j <= n; j++ {
			pre += stoneValue[j-1]
			f[i] = max(f[i], pre-f[j])
		}
	}
	if f[0] == 0 {
		return "Tie"
	} else if f[0] > 0 {
		return "Alice"
	} else {
		return "Bob"
	}
}
