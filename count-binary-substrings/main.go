/*
Package count_binary_substrings
https://leetcode-cn.com/problems/count-binary-substrings/

696. 计数二进制子串

给定一个字符串 s，统计并返回具有相同数量 0 和 1 的非空（连续）子字符串的数量，并且这些子字符串中的所有 0 和所有 1 都是成组连续的。

重复出现（不同位置）的子串也要统计它们出现的次数。

提示：
	1 <= s.length <= 105
	s[i] 为 '0' 或 '1'

*/
package count_binary_substrings

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func countBinarySubstrings(s string) int {
	ans := 0
	first, second := 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			first++
		} else {
			// 计算数量
			ans += min(first, second)
			second = first
			first = 1
		}
	}
	ans += min(first, second)

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
