/*
Package factorial_trailing_zeroes
https://leetcode-cn.com/problems/factorial-trailing-zeroes/

172. 阶乘后的零

给定一个整数 n，返回 n! 结果尾数中零的数量。

说明: 你算法的时间复杂度应为 O(log n) 。

*/
package factorial_trailing_zeroes

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func trailingZeroes(n int) int {
	res := 0
	for n >= 5 {
		res += n / 5
		n /= 5
	}

	return res
}
