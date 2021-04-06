/*
https://leetcode-cn.com/problems/regular-expression-matching/

10. 正则表达式匹配

给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

提示：
	0 <= s.length <= 20
	0 <= p.length <= 30
	s 可能为空，且只包含从 a-z 的小写字母。
	p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
	保证每次出现字符 * 时，前面都匹配到有效的字符

*/
package regular_expression_matching

import "strconv"

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func isMatch(s string, p string) bool {
	var dp func(i, j int) bool
	memo := make(map[string]bool)
	dp = func(i, j int) bool {
		if j == len(p) {
			return i == len(s)
		}

		if res, ok := memo[strconv.Itoa(i)+strconv.Itoa(j)]; ok {
			return res
		}
		ans := false
		first := i < len(s) && (p[j] == s[i] || p[j] == '.')
		if j < len(p)-1 && p[j+1] == '*' {
			ans = dp(i, j+2) || (first && dp(i+1, j))
		} else {
			ans = first && dp(i+1, j+1)
		}
		memo[strconv.Itoa(i)+strconv.Itoa(j)] = ans
		return ans
	}
	return dp(0, 0)
}
