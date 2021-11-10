/*
Package valid_palindrome_ii
https://leetcode-cn.com/problems/valid-palindrome-ii/

680. 验证回文字符串 Ⅱ

给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

提示:
	1 <= s.length <= 10^5
	s 由小写英文字母组成

*/
package valid_palindrome_ii

// --- 自己

/*
方法一:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func validPalindrome(s string) bool {
	var isPalindrome func(l, r int) bool
	isPalindrome = func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return isPalindrome(left+1, right) || isPalindrome(left, right-1)
		}
	}
	return true
}
