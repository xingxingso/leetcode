/*
Package counting_bits
https://leetcode-cn.com/problems/counting-bits/

338. 比特位计数

给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。

提示：
	0 <= n <= 10^5

进阶：
	很容易就能实现时间复杂度为 O(n log n) 的解决方案，你可以在线性时间复杂度 O(n) 内用一趟扫描解决此问题吗？
	你能不使用任何内置函数解决此问题吗？（如，C++ 中的 __builtin_popcount ）

*/
package counting_bits

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func countBits(n int) []int {
	ans := make([]int, n+1)
	ans[0] = 0
	for i := 1; i <= n; i++ {
		if (i-1)&1 == 0 {
			ans[i] = ans[i-1] + 1
		} else {
			x := ^(i - 1)
			p := x & (-x)
			c := 0
			for p > 1 {
				c++
				p >>= 1
			}
			ans[i] = ans[i-1] - c + 1
		}
	}
	return ans
}

// --- 他人

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func countBitsP1(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i&1 == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i>>1]
		}
	}
	return dp
}
