/*
Package generate_parentheses
https://leetcode-cn.com/problems/generate-parentheses/

22. 括号生成

数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

提示：
	1 <= n <= 8

*/
package generate_parentheses

// --- 自己

/*
方法一: 回溯

时间复杂度：
空间复杂度：
*/
func generateParenthesis(n int) []string {
	stack := make([]byte, 0)
	ans := make([]string, 0)
	buf := make([]byte, 0)
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == 2*n {
			if len(stack) == 0 && len(buf) == 2*n {
				temp := make([]byte, len(buf))
				copy(temp, buf)
				ans = append(ans, string(temp))
			}
			return
		}
		for ; i < 2*n; i++ {
			for _, v := range []byte{'(', ')'} {
				if len(stack) > n {
					continue
				}
				if len(stack) == 0 {
					if v == ')' {
						continue
					}
				}
				if v == '(' {
					stack = append(stack, v)
				} else {
					stack = stack[:len(stack)-1]
				}
				buf = append(buf, v)
				backtrack(i + 1)
				buf = buf[:len(buf)-1]

				if v == '(' {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, ')')
				}

			}
		}
	}

	backtrack(0)
	return ans
}

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485115&idx=1&sn=dd64bfedb1da22f308228a0933583adf&chksm=9bd7f8b3aca071a5b96e7cb9464c01c045997d36d677b14163b6b009df2aa9b1b613ace3bc5a&scene=21#wechat_redirect

/*
方法一: 回溯
	1、一个「合法」括号组合的左括号数量一定等于右括号数量，这个显而易见。
	2、对于一个「合法」的括号字符串组合p，必然对于任何0 <= i < len(p)都有：子串p[0..i]中左括号的数量都大于或等于右括号的数量。

时间复杂度：
空间复杂度：
*/
func generateParenthesisP1(n int) []string {
	track := make([]byte, 0)
	res := make([]string, 0)
	var backTrack func(left, right int)
	backTrack = func(left, right int) {
		if left < 0 || right < 0 {
			return
		}
		// 若左括号剩下的多，说明不合法
		if left > right {
			return
		}
		if left == 0 && right == 0 {
			temp := make([]byte, len(track))
			copy(temp, track)
			res = append(res, string(temp))
			return
		}

		track = append(track, '(')   // 选择
		backTrack(left-1, right)     // 回溯
		track = track[:len(track)-1] // 撤消选择

		track = append(track, ')')   // 选择
		backTrack(left, right-1)     // 回溯
		track = track[:len(track)-1] // 撤消选择
	}

	backTrack(n, n)
	return res
}
