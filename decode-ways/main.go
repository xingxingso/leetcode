/*
Package decode_ways
https://leetcode-cn.com/problems/decode-ways/

91. 解码方法

一条包含字母 A-Z 的消息通过以下映射进行了 编码 ：
	'A' -> 1
	'B' -> 2
	...
	'Z' -> 26
要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：
	- "AAJF" ，将消息分组为 (1 1 10 6)
	- "KJF" ，将消息分组为 (11 10 6)
注意，消息不能分组为  (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。

给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。

题目数据保证答案肯定是一个 32 位 的整数。

提示：
	1 <= s.length <= 100
	s 只包含数字，并且可能包含前导零。

*/
package decode_ways

// --- 自己

/*
方法一: 动态规划
	将 '10', '20' 视作一个字符

时间复杂度：
空间复杂度：
*/
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre := 0
	last := 0
	res := 0
	for i := 0; i < len(s); i++ {
		temp := res
		cur := int(s[i]) - '0'
		if cur == 0 {
			return 0
		}
		if i < len(s)-1 && s[i+1] == '0' {
			cur = 10*cur + int(s[i+1]) - '0'
			i++
		}
		if cur > 26 {
			return 0
		}
		if pre == 0 {
			res = 1
		} else if pre == 1 {
			if cur < 10 {
				if last == 0 { // 相当于进行到第2个字符，此时前面的前面没有字符
					res = 2
				} else {
					res = last + res
				}
			}
		} else if pre == 2 {
			if cur < 7 {
				if last == 0 {
					res = 2
				} else {
					res = last + res
				}
			}
		}
		pre = cur
		last = temp
		//fmt.Println(i, res, last)
	}

	return res
}
