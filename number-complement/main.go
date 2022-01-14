/*
Package number_complement
https://leetcode-cn.com/problems/number-complement/

476. 数字的补数

对整数的二进制表示取反（0 变 1 ，1 变 0）后，再转换为十进制表示，可以得到这个整数的补数。

例如，整数 5 的二进制表示是 "101" ，取反后得到 "010" ，再转回十进制表示得到补数 2 。
给你一个整数 num ，输出它的补数。

提示：
	1 <= num < 2^31

注意：
	本题与 1009 https://leetcode-cn.com/problems/complement-of-base-10-integer/ 相同

*/
package number_complement

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func findComplement(num int) int {
	ans := 0
	for k := 0; num > 0; k++ {
		p := num & 1
		ans += (p ^ 1) << k
		num >>= 1
	}
	return ans
}
