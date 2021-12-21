/*
Package add_strings
https://leetcode-cn.com/problems/add-strings/

415. 字符串相加

给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。

你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。

提示：
	1 <= num1.length, num2.length <= 10^4
	num1 和num2 都只包含数字 0-9
	num1 和num2 都不包含任何前导零

*/
package add_strings

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func addStrings(num1 string, num2 string) string {
	ans := make([]byte, 0, max(len(num1), len(num2)))
	rem := uint8(0)
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; {
		a, b := uint8(0), uint8(0)
		if i >= 0 {
			a = num1[i] - '0'
			i--
		}
		if j >= 0 {
			b = num2[j] - '0'
			j--
		}
		p := (a + b + rem) % 10
		rem = (a + b + rem) / 10
		ans = append(ans, p+'0')
	}
	if rem > 0 {
		ans = append(ans, rem+'0')
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
