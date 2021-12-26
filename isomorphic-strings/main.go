/*
Package isomorphic_strings
https://leetcode-cn.com/problems/isomorphic-strings/

205. 同构字符串

给定两个字符串 s 和 t，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

提示：
	可以假设 s 和 t 长度相同。

*/
package isomorphic_strings

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func isIsomorphic(s string, t string) bool {
	memo := make(map[uint8]uint8)
	reverse := make(map[uint8]uint8)

	for i := 0; i < len(s); i++ {
		if memo[s[i]] == 0 {
			if reverse[t[i]] != 0 {
				return false
			}
			memo[s[i]] = t[i]
			reverse[t[i]] = s[i]
		} else {
			if memo[s[i]] != t[i] {
				return false
			}
		}
	}

	return true
}

// --- 他人

/*
方法一:

时间复杂度：
空间复杂度：
*/
func isIsomorphicP1(s string, t string) bool {
	sFirstIndex, tFirstIndex := [256]int{}, [256]int{}

	for i := 0; i < len(s); i++ {
		if sFirstIndex[s[i]] != tFirstIndex[t[i]] {
			return false
		}
		sFirstIndex[s[i]], tFirstIndex[t[i]] = i+1, i+1 // 不能用 i 替代 i+1, 如 ab, aa
	}

	return true
}
