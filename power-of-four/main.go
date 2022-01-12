/*
Package power_of_four
https://leetcode-cn.com/problems/power-of-four/

342. 4的幂

给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4x

提示：
	-2^31 <= n <= 2^31 - 1

进阶：
	你能不使用循环或者递归来完成本题吗？

*/
package power_of_four

// --- 自己

/*
方法一:
	4^n = 2^(2n)

时间复杂度：
空间复杂度：
*/
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}

	if n&(n-1) > 0 { // 最低位1去除后大于0 表示 不只一个 1
		return false
	}

	// 只有1个1 且位于倒数的奇数位
	// 0101010101010101010101010101010101010101010101010101010101010101, 64bit
	return 0x5555555555555555&n > 0
	//return 0x55555555&n > 0
}
