/*
Package ugly_number_ii
https://leetcode-cn.com/problems/ugly-number-ii/

264. 丑数 II

给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是只包含质因数 2、3 和/或 5 的正整数。

提示：
	1 <= n <= 1690

*/
package ugly_number_ii

// --- 官方

/*
方法一:

时间复杂度：
空间复杂度：
*/
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[p2]*2, dp[p3]*3, dp[p5]*5)
		if dp[i] == dp[p2]*2 {
			p2++
		}
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}
	}
	//fmt.Println(dp)

	return dp[n]
}

func min(a, b, c int) int {
	if a > b {
		if b > c {
			return c
		}
		return b
	}
	if a > c {
		return c
	}
	return a
}
