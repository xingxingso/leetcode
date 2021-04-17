/*
https://leetcode-cn.com/problems/stone-game/

877. 石子游戏

亚历克斯和李用几堆石子在做游戏。偶数堆石子排成一行，每堆都有正整数颗石子 piles[i] 。
游戏以谁手中的石子最多来决出胜负。石子的总数是奇数，所以没有平局。
亚历克斯和李轮流进行，亚历克斯先开始。 每回合，玩家从行的开始或结束处取走整堆石头。 这种情况一直持续到没有更多的石子堆为止，此时手中石子最多的玩家获胜。
假设亚历克斯和李都发挥出最佳水平，当亚历克斯赢得比赛时返回 true ，当李赢得比赛时返回 false 。

提示：
	2 <= piles.length <= 500
	piles.length 是偶数。
	1 <= piles[i] <= 500
	sum(piles) 是奇数。

*/
package stone_game

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func stoneGame(piles []int) bool {
	dp := make([][][2]int, len(piles))
	for i := len(piles) - 1; i >= 0; i-- {
		dp[i] = make([][2]int, len(piles))
		for j := i; j < len(piles); j++ {
			if j == i {
				dp[i][j][0] = piles[i]
				dp[i][j][1] = 0
			} else {
				left, right := piles[i]+dp[i+1][j][1], piles[j]+dp[i][j-1][1]
				if left > right {
					dp[i][j][0] = left
					dp[i][j][1] = dp[i+1][j][0]
				} else {
					dp[i][j][0] = right
					dp[i][j][1] = dp[i][j-1][0]
				}
			}
		}
	}
	return dp[0][len(piles)-1][0] > dp[0][len(piles)-1][1]
}

// --- 官方

/*
方法一: 动态规划

时间复杂度：O(n^2)，其中 n 是数组的长度。需要计算每个子数组对应的 dp 的值，共有 n(n+1)/2 个子数组。
空间复杂度：O(n^2)，其中 n 是数组的长度。
*/
func stoneGameO1(piles []int) bool {
	length := len(piles)
	// dp[i][j] 表示当剩下的石子堆为下标 ii 到下标 jj 时，当前玩家与另一个玩家的石子数量之差的最大值，注意当前玩家不一定是先手 Alex。
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = piles[i]
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][length-1] > 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法一: 动态规划优化

时间复杂度：O(n^2)，其中 n 是数组的长度。需要计算每个子数组对应的 dp 的值，共有 n(n+1)/2 个子数组。
空间复杂度：O(n)，其中 n 是数组的长度。空间复杂度取决于额外创建的数组 dp，如果不优化空间，则空间复杂度是 O(n^2)，
	使用一维数组优化之后空间复杂度可以降至 O(n)。
*/
func stoneGameO2(piles []int) bool {
	length := len(piles)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = piles[i]
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[j] = max(piles[i]-dp[j], piles[j]-dp[j-1])
		}
	}
	return dp[length-1] > 0
}

/*
方法二: 数学
	作为先手的 Alex 可以在第一次取走石子时就决定取走哪一组的石子，并全程取走同一组的石子。

时间复杂度：O(1)。
空间复杂度：O(1)。
*/
func stoneGameO3(piles []int) bool {
	return true
}
