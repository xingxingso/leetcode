/*
https://leetcode-cn.com/problems/stone-game-ii/

1140. 石子游戏 II

亚历克斯和李继续他们的石子游戏。许多堆石子 排成一行，每堆都有正整数颗石子 piles[i]。游戏以谁手中的石子最多来决出胜负。
亚历克斯和李轮流进行，亚历克斯先开始。最初，M = 1。
在每个玩家的回合中，该玩家可以拿走剩下的 前 X 堆的所有石子，其中 1 <= X <= 2M。然后，令 M = max(M, X)。
游戏一直持续到所有石子都被拿走。
假设亚历克斯和李都发挥出最佳水平，返回亚历克斯可以得到的最大数量的石头。

提示：
	1 <= piles.length <= 100
	1 <= piles[i] <= 10^4

*/
package stone_game_ii

// --- 他人

/*
方法一: 动态规划
	https://leetcode-cn.com/problems/stone-game-ii/solution/java-dong-tai-gui-hua-qing-xi-yi-dong-17xing-by-lg/

时间复杂度：
空间复杂度：
*/
func stoneGameIIP1(piles []int) int {
	// [i, length-1] m 为 j  先取的人能拿到的最大
	// dp[i][j]
	length := len(piles)
	dp := make([][]int, length)
	sum := 0
	for i := length - 1; i >= 0; i-- {
		sum += piles[i]
		dp[i] = make([]int, length+1)
		for j := 1; j <= length; j++ {
			if i+2*j >= length {
				dp[i][j] = sum
			} else {
				for k := 1; k <= 2*j; k++ {
					dp[i][j] = max(dp[i][j], sum-dp[i+k][max(j, k)])
				}
			}
		}
	}
	return dp[0][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
