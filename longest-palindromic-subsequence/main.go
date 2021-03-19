/*
https://leetcode-cn.com/problems/longest-palindromic-subsequence/

516. 最长回文子序列

给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。

提示：
	1 <= s.length <= 1000
	s 只包含小写英文字母

*/
package longest_palindromic_subsequence

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func longestPalindromeSubseq(s string) int {
	// dp := make([][]int, len(s)+1)
	d := make([]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		// dp[i] = make([]int, len(s)+1)
		// dp[i][i] = 1
		d[i] = 1
		pre := 0
		for j := i + 1; j < len(s); j++ {
			temp := d[j]
			if s[i] == s[j] {
				// dp[i][j] = dp[i+1][j-1] + 2
				d[j] = pre + 2
			} else {
				// dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				d[j] = max(d[j], d[j-1])
			}
			pre = temp
		}
	}
	// return dp[0][len(s)-1]
	return d[len(s)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
