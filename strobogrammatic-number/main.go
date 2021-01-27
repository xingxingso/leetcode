/*
https://leetcode-cn.com/problems/strobogrammatic-number/

246. 中心对称数

中心对称数是指一个数字在旋转了 180 度之后看起来依旧相同的数字（或者上下颠倒地看）。
请写一个函数来判断该数字是否是中心对称数，其输入将会以一个字符串的形式来表达数字。

*/
package strobogrammatic_number

// --- 自己

/*
方法一:
	字符	  0  1  8  | 6  9
	ASCII 48 49 56 | 54 57

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isStrobogrammatic(num string) bool {
	l := 0
	r := len(num) - 1
	for l <= r {
		if num[l] == '0' && num[r] == '0' ||
			num[l] == '1' && num[r] == '1' ||
			num[l] == '8' && num[r] == '8' ||
			num[l] == '6' && num[r] == '9' ||
			num[l] == '9' && num[r] == '6' {
			l++
			r--
		} else {
			return false
		}
	}

	return true
}
