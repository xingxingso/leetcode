/*
Package binary_number_with_alternating_bits
https://leetcode-cn.com/problems/binary-number-with-alternating-bits/

693. 交替位二进制数

给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：换句话说，就是二进制表示中相邻两位的数字永不相同。

提示：
	1 <= n <= 2^31 - 1

*/
package binary_number_with_alternating_bits

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func hasAlternatingBits(n int) bool {
	p := n & 1
	n >>= 1

	for ; n > 0; n >>= 1 {
		q := n & 1
		if p == q {
			return false
		}
		p = q
	}

	return true
}
