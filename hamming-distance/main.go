/*
Package hamming_distance
https://leetcode-cn.com/problems/hamming-distance/

461. 汉明距离

两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。

给你两个整数 x 和 y，计算并返回它们之间的汉明距离。

提示：
	0 <= x, y <= 2^31 - 1

*/
package hamming_distance

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func hammingDistance(x int, y int) int {
	p := x ^ y // 统计 p 的 1 的数量

	ans := 0
	for p > 0 {
		p = p & (p - 1)
		ans++
	}

	return ans
}
