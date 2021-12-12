/*
Package _2_keys_keyboard
https://leetcode-cn.com/problems/2-keys-keyboard/

650. 只有两个键的键盘

最初记事本上只有一个字符 'A' 。你每次可以对这个记事本进行两种操作：

Copy All（复制全部）：复制这个记事本中的所有字符（不允许仅复制部分字符）。
Paste（粘贴）：粘贴 上一次 复制的字符。
给你一个数字 n ，你需要使用最少的操作次数，在记事本上输出 恰好 n 个 'A' 。返回能够打印出 n 个 'A' 的最少操作次数。

提示：
	1 <= n <= 1000

*/
package _2_keys_keyboard

// --- 官方

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func minStepsO1(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				dp[i] = min(dp[i], dp[j]+i/j)
				dp[i] = min(dp[i], dp[i/j]+j)
			}
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
