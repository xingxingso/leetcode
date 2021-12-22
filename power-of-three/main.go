/*
Package power_of_three
https://leetcode-cn.com/problems/power-of-three/

326. 3 的幂

给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3^x

提示：
	-2^31 <= n <= 2^31 - 1

进阶：
	你能不使用循环或者递归来完成本题吗？

*/
package power_of_three

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}

/*
方法二:
	3^19=1162261467
	2^31=2147483648
	3^20=3486784401

时间复杂度：
空间复杂度：
*/
func isPowerOfThreeS2(n int) bool {
	return n > 0 && 1162261467 % n == 0
}
