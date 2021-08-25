/*
Package multiply_strings
https://leetcode-cn.com/problems/multiply-strings/

43. 字符串相乘

给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

说明：
	1. num1 和 num2 的长度小于110。
	2. num1 和 num2 只包含数字 0-9。
	3. num1 和 num2 均不以零开头，除非是数字 0 本身。
	4. 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

*/
package multiply_strings

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func multiply(num1 string, num2 string) string {
	ans := make([]byte, len(num2+num1))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			p := (num1[i]-'0')*(num2[j]-'0') + ans[i+j+1]
			ans[i+j+1] = p % 10
			ans[i+j] += p / 10
		}
	}
	k := 0
	for k < len(ans)-1 && ans[k] == 0 {
		k++
	}
	ans = ans[k:]
	for i := 0; i < len(ans); i++ {
		ans[i] += '0'
	}

	return string(ans)
}
