/*
Package base_7
https://leetcode-cn.com/problems/base-7/

504. 七进制数

给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。

提示：
	-10^7 <= num <= 10^7

*/
package base_7

import (
	"strconv"
)

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func convertToBase7(num int) string {
	ans := ""
	flag := false
	if num < 0 {
		flag = true
		num = -num
	}
	for ; num >= 7; num = num / 7 {
		ans = strconv.Itoa(num%7) + ans
	}

	ans = strconv.Itoa(num) + ans
	if flag {
		ans = "-" + ans
	}
	return ans
}
