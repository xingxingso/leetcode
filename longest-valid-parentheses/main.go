/*
Package longest_valid_parentheses
https://leetcode-cn.com/problems/longest-valid-parentheses/

32. 最长有效括号

给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

提示：
	0 <= s.length <= 3 * 10^4
	s[i] 为 '(' 或 ')'

*/
package longest_valid_parentheses

// --- 官方

/*
方法一: 动态规划
	dp[i]=dp[i−2]+2
	dp[i]=dp[i−1]+dp[i−dp[i−1]−2]+2

时间复杂度： O(n)，其中 n 为字符串的长度。我们只需遍历整个字符串一次，即可将 dp 数组求出来。
空间复杂度：O(n)。我们需要一个大小为 n 的 dp 数组。
*/
func longestValidParenthesesO1(s string) int {
	ans := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' && s[i-1] == '(' {
			if i < 2 {
				dp[i] = 2
			} else {
				dp[i] = dp[i-2] + 2
			}
		} else if s[i] == ')' && s[i-1] == ')' && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			k := i - dp[i-1] - 2
			if k < 0 {
				dp[i] = dp[i-1] + 2
			} else {
				dp[i] = dp[i-1] + dp[k] + 2
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
