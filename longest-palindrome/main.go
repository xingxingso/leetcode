/*
Package longest_palindrome
https://leetcode-cn.com/problems/longest-palindrome/

409. 最长回文串

给定一个包含大写字母和小写字母的字符串 s ，返回 通过这些字母构造成的 最长的回文串 。

在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。

提示:
	1 <= s.length <= 2000
	s 只能由小写和/或大写英文字母组成

*/
package longest_palindrome

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func longestPalindrome(s string) int {
	mark := [126]int{}
	for _, v := range s {
		mark[v]++
	}

	ans := 0
	hasOdd := false
	for _, n := range mark {
		if n == 0 {
			continue
		}
		if n%2 == 0 {
			ans += n
		} else {
			hasOdd = true
			ans += n - 1
		}
	}

	if hasOdd {
		ans++
	}

	return ans
}
