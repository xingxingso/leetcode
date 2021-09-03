/*
Package valid_parentheses
https://leetcode-cn.com/problems/valid-parentheses/

20. 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。

提示：
	1 <= s.length <= 104
	s 仅由括号 '()[]{}' 组成

*/
package valid_parentheses

// --- 自己

/*
方法一: 栈

时间复杂度：O(n)，其中 n 是字符串 s 的长度。
空间复杂度：O(n+∣Σ∣)，其中 Σ 表示字符集，本题中字符串只包含 6 种括号，∣Σ∣=6。栈中的字符数量为 O(n)，而哈希表使用的空间为 O(∣Σ∣)，相加即可得到总空间复杂度。
*/
func isValid(s string) bool {
	if len(s) < 1 {
		return true
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack) < 1 {
				return false
			}
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !isMatch(p, s[i]) {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func isMatch(a, b byte) bool {
	return a == '(' && b == ')' || a == '[' && b == ']' || a == '{' && b == '}'
}
