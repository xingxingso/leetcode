/*
Package integer_break
https://leetcode-cn.com/problems/integer-break/

343. 整数拆分

给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

说明:
	你可以假设 n 不小于 2 且不大于 58。

*/
package integer_break

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	dp[1] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i-2; j++ {
			dp[i] = max(dp[i], max(j*(i-j), dp[j]*(i-j)))
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
