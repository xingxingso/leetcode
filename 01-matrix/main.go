/*
Package _01_matrix
https://leetcode-cn.com/problems/01-matrix/

542. 01 矩阵

给定一个由 0 和 1 组成的矩阵 mat ，请输出一个大小相同的矩阵，其中每一个格子是 mat 中对应位置元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。

提示：
	m == mat.length
	n == mat[i].length
	1 <= m, n <= 10^4
	1 <= m * n <= 10^4
	mat[i][j] is either 0 or 1.
	mat 中至少有一个 0

*/
package _01_matrix

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				dp[i][j] = 20000 // 1 <= m, n <= 10^4
			} else if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + 1
			} else if i > 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + 1
			} else {
				dp[i][j] = minN(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			if mat[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			if i == m-1 && j == n-1 {
				dp[i][j] = minN(dp[i][j], 20000)
			} else if i == m-1 && j < n-1 {
				dp[i][j] = minN(dp[i][j+1], dp[i][j]-1) + 1
			} else if i < m-1 && j == n-1 {
				dp[i][j] = minN(dp[i+1][j], dp[i][j]-1) + 1
			} else {
				dp[i][j] = minN(dp[i+1][j], dp[i][j+1], dp[i][j]-1) + 1
			}
		}
	}

	return dp
}

func minN(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}
