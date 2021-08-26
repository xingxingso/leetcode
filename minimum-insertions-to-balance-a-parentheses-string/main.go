/*
Package minimum_insertions_to_balance_a_parentheses_string
https://leetcode-cn.com/problems/minimum-insertions-to-balance-a-parentheses-string/

1541. 平衡括号字符串的最少插入次数

给你一个括号字符串s，它只包含字符'(' 和')'。一个括号字符串被称为平衡的当它满足：
	- 任何左括号'('必须对应两个连续的右括号'))'。
	- 左括号'('必须在对应的连续两个右括号'))'之前。
比方说"())"，"())(())))" 和"(())())))"都是平衡的，")()"，"()))" 和"(()))"都是不平衡的。
你可以在任意位置插入字符 '(' 和 ')' 使字符串平衡。
请你返回让s平衡的最少插入次数。

提示：
	1 <= s.length <= 10^5
	s 只包含 '(' 和 ')' 。

*/
package minimum_insertions_to_balance_a_parentheses_string

// --- 自己

/*
方法一：

时间复杂度：
空间复杂度：
*/
func minInsertions(s string) int {
	ans := 0
	stack := make([]byte, 0)
	count := 0
	for i := 0; i <= len(s); i++ {
		if i < len(s) && s[i] == ')' {
			count++
			continue
		}

		if count == 0 {
			if i < len(s) {
				stack = append(stack, '(')
			}
			continue
		}

		pickupNumber := (count + 1) / 2
		if pickupNumber*2 > count {
			ans++
		}
		if len(stack) >= pickupNumber {
			stack = stack[:len(stack)-pickupNumber]
		} else {
			ans += pickupNumber - len(stack)
			stack = stack[:0]
		}
		count = 0
		if i < len(s) {
			stack = append(stack, '(')
		}
	}
	return ans + 2*len(stack)
}

// ---他人
//https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247487246&idx=1&sn=4a514020ce9dc8777e2d1d503188b62b&scene=21#wechat_redirect

/*
方法一：

时间复杂度：
空间复杂度：
*/
func minInsertionsP1(s string) int {
	res, need := 0, 0
	for _, v := range s {
		if v == '(' {
			need += 2
			if need%2 == 1 {
				res++
				need--
			}
			continue
		}
		need--
		if need == -1 {
			res++
			need = 1
		}
	}
	return res + need
}
