/*
Package palindromic_substrings
https://leetcode-cn.com/problems/palindromic-substrings/

647. 回文子串

给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

提示：
	1 <= s.length <= 1000
	s 由小写英文字母组成

*/
package palindromic_substrings

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func countSubstrings(s string) int {
	count := 0
	palindrome := func(i, j int) {
		for i >= 0 && j < len(s) && s[i] == s[j] {
			i--
			j++
			count++
		}
	}

	for i := 0; i < len(s); i++ {
		palindrome(i, i)   // 以 i 为中心
		palindrome(i, i+1) // 以 i, i+1 为中心
	}
	return count
}
