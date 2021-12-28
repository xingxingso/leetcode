/*
Package unique_paths
https://leetcode-cn.com/problems/unique-paths/

62. 不同路径

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

提示：
	1 <= m, n <= 100
	题目数据保证答案小于等于 2 * 10^9

*/
package unique_paths

// --- 自己

/*
方法一: 回溯
	超出时限

时间复杂度：
空间复杂度：
*/
func uniquePaths(m int, n int) int {
	var backtrack func(m, n int) int
	backtrack = func(m, n int) int {
		if m == 1 || n == 1 {
			return 1
		}
		var a, b int
		if m > 0 {
			a = backtrack(m-1, n)
		}
		if n > 0 {
			b = backtrack(m, n-1)
		}
		return a + b
	}

	return backtrack(m, n)
}

/*
方法二: 动态规划

时间复杂度：
空间复杂度：
*/
func uniquePathsS2(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = 1
				continue
			}
			if j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

/*
方法二: 动态规划
	状态压缩

时间复杂度：
空间复杂度：
*/
func uniquePathsS3(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[j] = 1
				continue
			}
			if j == 0 {
				dp[j] = 1
				continue
			}
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}
