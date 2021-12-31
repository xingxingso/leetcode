/*
Package regular_expression_matching
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

// --- 他人

/*
方法一: 动态规划

	// 东哥手写正则通配符算法，结构清晰，包教包会！
	// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247486742&idx=1&sn=73d38d4d8b51af81b782c6d11fa5e21e

时间复杂度：
空间复杂度：
*/
func isMatchP1(s string, p string) bool {
	memo := make(map[string]bool)
	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if j == len(p) {
			return i == len(s)
		}

		//if i == len(s) {
		//	return j == len(p) // 这样不对，如 s="a", p="ab*c*" 依然满足，只要p[j..]能够匹配空串，就可以算完成匹配
		//}

		// 这里的处理复杂了， 可以参考 isMatch 我的解法
		if i == len(s) {
			// 如果能匹配空串，一定是字符和 * 成对儿出现
			if (len(p)-j)%2 == 1 {
				return false
			}
			// 检查是否为 x*y*z* 这种形式
			for ; j+1 < len(p); j += 2 {
				if p[j+1] != '*' {
					return false
				}
			}
			return true
		}

		// 记录状态 (i, j)，消除重叠子问题
		if res, ok := memo[strconv.Itoa(i)+strconv.Itoa(j)]; ok {
			return res
		}
		res := false
		if s[i] == p[j] || p[j] == '.' {
			if j < len(p)-1 && p[j+1] == '*' {
				res = dp(i, j+2) || dp(i+1, j)
			} else {
				res = dp(i+1, j+1)
			}
		} else {
			if j < len(p)-1 && p[j+1] == '*' {
				res = dp(i, j+2) // 这时只能匹配0个， s="ab" p="b*a."
			} else {
				res = false
			}
		}
		memo[strconv.Itoa(i)+strconv.Itoa(j)] = res
		return res
	}
	return dp(0, 0)
}
