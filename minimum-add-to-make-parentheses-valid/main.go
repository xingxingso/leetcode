/*
Package minimum_add_to_make_parentheses_valid
https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid/

921. 使括号有效的最少添加

给定一个由'('和')'括号组成的字符串 S，我们需要添加最少的括号（ '('或是')'，可以在任何位置），以使得到的括号字符串有效。

从形式上讲，只有满足下面几点之一，括号字符串才是有效的：
	- 它是一个空字符串，或者
	- 它可以被写成AB（A与B连接）, 其中A和B都是有效字符串，或者
	- 它可以被写作(A)，其中A是有效字符串。
给定一个括号字符串，返回为使结果字符串有效而必须添加的最少括号数。

提示：
	S.length <= 1000
	S 只包含 '(' 和 ')' 字符。

*/
package minimum_add_to_make_parentheses_valid

// --- 自己

/*
方法一：

时间复杂度：
空间复杂度：
*/
func minAddToMakeValid(s string) int {
	left := 0
	ans := 0
	for _, v := range s {
		if v == '(' {
			left++
		} else {
			left--
		}
		if left < 0 {
			ans++
			left = 0
		}
	}
	ans += left

	return ans
}
