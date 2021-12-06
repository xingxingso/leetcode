/*
Package minimum_path_sum
https://leetcode-cn.com/problems/minimum-path-sum/

64. 最小路径和

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

提示：
	m == grid.length
	n == grid[i].length
	1 <= m, n <= 200
	0 <= grid[i][j] <= 100

*/
package minimum_path_sum

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				if j == 0 {
					dp[i][j] = grid[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				}
			} else {
				if j == 0 {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
方法一: 动态规划
	空间压缩

时间复杂度：
空间复杂度：
*/
func minPathSumS2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j == 0 {
					dp[j] = grid[i][j]
				} else {
					dp[j] = dp[j-1] + grid[i][j]
				}
			} else {
				if j == 0 {
					dp[j] = dp[j] + grid[i][j]
				} else {
					dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
				}
			}
		}
	}
	return dp[n-1]
}
