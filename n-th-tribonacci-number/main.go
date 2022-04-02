/*
Package n_th_tribonacci_number
https://leetcode-cn.com/problems/n-th-tribonacci-number/

1137. 第 N 个泰波那契数

泰波那契序列 Tn 定义如下：

T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

给你整数 n，请返回第 n 个泰波那契数 Tn 的值。

提示：
	0 <= n <= 37
	答案保证是一个 32 位整数，即 answer <= 2^31 - 1。

*/
package n_th_tribonacci_number

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func tribonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if n == 2 {
		return 1
	}

	p0, p1, p2 := 0, 1, 1
	var p3 int
	for n >= 3 {
		p3 = p0 + p1 + p2
		p0, p1, p2 = p1, p2, p3
		n--
	}

	return p3
}
