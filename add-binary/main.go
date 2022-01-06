/*
Package add_binary
https://leetcode-cn.com/problems/add-binary/

67. 二进制求和

给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。

提示：
	每个字符串仅由字符 '0' 或 '1' 组成。
	1 <= a.length, b.length <= 10^4
	字符串如果不是 "0" ，就都不含前导零。

*/
package add_binary

// --- 自己

/*
方法一:
	利用 当前位: b1 ^ b2, 进位： b1 && b2

时间复杂度：
空间复杂度：
*/
func addBinary(a string, b string) string {
	ans := make([]byte, max(len(a), len(b))+1)
	pre := uint8(0b0) // 这里直接用0也可以
	for i := 1; i <= len(a) || i <= len(b); i++ {
		b1, b2 := uint8(0b0), uint8(0b0) // 这里直接用0也可以
		if i <= len(a) {
			b1 = a[len(a)-i] - '0'
		}
		if i <= len(b) {
			b2 = b[len(b)-i] - '0'
		}

		// 计算 b1 + b2
		s := b1 ^ b2
		c := b1 & b2 // 计算进位

		// 加上进位
		ans[len(ans)-i] = s ^ pre + '0'
		//fmt.Printf("%c\n", ans[len(ans)-i])
		pre = (s & pre) | c // 重新计算进位
	}

	ans[0] = pre + '0'
	if ans[0] == '0' {
		ans = ans[1:]
	}

	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
