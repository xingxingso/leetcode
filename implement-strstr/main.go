/*
Package implement_strstr
https://leetcode-cn.com/problems/implement-strstr/

28. 实现 strStr()

实现 strStr() 函数。
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。
如果不存在，则返回  -1 。

说明：
	当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
	对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

提示：
	0 <= haystack.length, needle.length <= 5 * 10^4
	haystack 和 needle 仅由小写英文字符组成

*/
package implement_strstr

// --- 自己

/*
方法一: KMP算法简化版
	参考 labuladong

时间复杂度：
空间复杂度：
*/
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}

	// needle 构建状态机
	m := len(needle)
	dfa := make([][256]int, m)
	dfa[0][needle[0]] = 1
	x := 0 // 影子状态
	for i := 1; i < m; i++ {
		for c := 0; c < 256; c++ {
			dfa[i][c] = dfa[x][c]
		}
		dfa[i][needle[i]] = i + 1
		// !!! 核心: 更新影子状态
		x = dfa[x][needle[i]]
	}

	// 使用状态机匹配 haystack
	n := len(haystack)
	j := 0
	for i := 0; i < n; i++ {
		j = dfa[j][haystack[i]]
		if j == m {
			return i - m + 1
		}
	}

	return -1
}

// --- 官方

/*
方法一: Knuth-Morris-Pratt 算法

时间复杂度：O(n+m)，其中 n 是字符串 haystack 的长度，m 是字符串 needle 的长度。我们至多需要遍历两字符串一次。
空间复杂度：O(m)，其中 m 是字符串 needle 的长度。我们只需要保存字符串 needle 的前缀函数。
*/
func strStrO1(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
