/*
Package remove_k_digits
https://leetcode-cn.com/problems/remove-k-digits/

402. 移掉 K 位数字

给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。

提示：
	1 <= k <= num.length <= 10^5
	num 仅由若干位数字（0 - 9）组成
	除了 0 本身之外，num 不含任何前导零

*/
package remove_k_digits

// --- 自己

/*
方法一: 单调栈

时间复杂度：
空间复杂度：
*/
func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	stack := make([]byte, 0)
	for i := 0; i < len(num); i++ {
		if len(stack) == 0 || k == 0 {
			stack = append(stack, num[i])
			continue
		}

		// 栈顶元素大于 num[i] 则弹出, 注意 k <= 0 的时候无需再弹出
		j := len(stack) - 1
		for ; j >= 0 && stack[j] > num[i] && k > 0; j-- {
			k--
		}
		stack = stack[:j+1]

		stack = append(stack, num[i])
	}

	// 剩余的 k 直接弹出
	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	// 处理前缀0
	i := 0
	for ; i < len(stack) && stack[i] == '0'; i++ {
	}
	if i >= len(stack) {
		return "0"
	}

	return string(stack[i:])
}
