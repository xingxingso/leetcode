/*
https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/

1312. 让字符串成为回文串的最少插入次数

给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。
请你返回让 s 成为回文串的 最少操作次数 。
「回文串」是正读和反读都相同的字符串。

提示：
	1 <= s.length <= 500
	s 中所有字符都是小写字母。

*/
package minimum_insertion_steps_to_make_a_string_palindrome

// --- 自己

/*
方法一: 动态规划
	个人理解为 求最大子回文串 如 s:mbadm->p:mbm->len(s)-len(p)->5-3=2

时间复杂度：
空间复杂度：
*/
func minInsertions(s string) int {
	// i,j 的最大子回文串
	l := len(s)
	dp := make([][]int, l)

	for i := l - 1; i >= 0; i-- {
		dp[i] = make([]int, l)
		dp[i][i] = 1
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
		// fmt.Println(i, dp[i])
	}

	return l - dp[0][l-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法一: 动态规划优化
	个人理解为 求最大子回文串 如 s:mbadm->p:mbm->len(s)-len(p)->5-3=2

时间复杂度：
空间复杂度：
*/
func minInsertionsS2(s string) int {
	// i,j 的最大子回文串
	l := len(s)
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 1
	}

	for i := l - 1; i >= 0; i-- {
		pre := 0
		for j := i + 1; j < l; j++ {
			temp := dp[j]
			if s[i] == s[j] {
				dp[j] = pre + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			pre = temp
		}
		// fmt.Println(i, dp)
	}

	return l - dp[l-1]
}
