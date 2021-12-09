/*
Package maximal_square
https://leetcode-cn.com/problems/maximal-square/

221. 最大正方形

在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

提示：
	m == matrix.length
	n == matrix[i].length
	1 <= m, n <= 300
	matrix[i][j] 为 '0' 或 '1'

*/
package maximal_square

// --- 他人

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func maximalSquareP1(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	maxL := 0
	dp[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
			if maxL < dp[i][j] {
				maxL = dp[i][j]
			}
		}
	}
	return maxL * maxL
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
