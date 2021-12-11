/*
Package ones_and_zeroes
https://leetcode-cn.com/problems/ones-and-zeroes/

474. 一和零

给你一个二进制字符串数组 strs 和两个整数 m 和 n 。

请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。

如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

提示：
	1 <= strs.length <= 600
	1 <= strs[i].length <= 100
	strs[i] 仅由 '0' 和 '1' 组成
	1 <= m, n <= 100

*/
package ones_and_zeroes

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs))
	for i := 0; i < len(strs); i++ {
		dp[i] = make([][]int, m+1)
		p, q := countOnesZeros(strs[i])
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
			for k := 0; k <= n; k++ {
				if i == 0 {
					if j >= p && k >= q {
						dp[i][j][k] = 1
					}
					continue
				}
				if j >= p && k >= q {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-p][k-q]+1)
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}
	return dp[len(strs)-1][m][n]
}

/*
方法一: 动态规划
	空间压缩

时间复杂度：
空间复杂度：
*/
func findMaxFormS2(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		p, q := countOnesZeros(strs[i])
		for j := m; j >= p; j-- {
			for k := n; k >= q; k-- {
				dp[j][k] = max(dp[j][k], dp[j-p][k-q]+1)
			}
		}
	}
	return dp[m][n]
}

func countOnesZeros(s string) (p, q int) {
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			p++
		} else {
			q++
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
