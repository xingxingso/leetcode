/*
Package longest_common_subsequence
https://leetcode-cn.com/problems/longest-common-subsequence/

1143. 最长公共子序列

给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。

提示:
	1 <= text1.length <= 1000
	1 <= text2.length <= 1000
	输入的字符串只含有小写英文字符。

*/
package longest_common_subsequence

import (
	"fmt"
)

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func longestCommonSubsequenceS1(text1 string, text2 string) int {
	return 0
}

// --- 他人

/*
方法一: 动态规划
	https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247487860&idx=1&sn=f5759ae4f22f966db8ed5a85821edd34&chksm=9bd7ef7caca0666a628fe838dee6d5da44b05eadf01fd7e87fcef813430c8e6dc3eb3c23e15f&scene=21#wechat_redirect
	自顶向下

时间复杂度：
空间复杂度：
*/
func longestCommonSubsequenceP1(text1 string, text2 string) int {
	var dp func(i, j int) int
	memo := make(map[string]int)
	dp = func(i, j int) int {
		if n, ok := memo[fmt.Sprintf("%d,%d", i, j)]; ok {
			return n
		}

		if i == len(text1) || j == len(text2) {
			return 0
		}
		if text1[i] == text2[j] {
			return dp(i+1, j+1) + 1
		}
		// dp(i+1, j+1) 可以省略
		n := maxN(dp(i+1, j+1), dp(i+1, j), dp(i, j+1))
		memo[fmt.Sprintf("%d,%d", i, j)] = n
		return n
	}

	return dp(0, 0)
}

func maxN(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

/*
方法一: 动态规划
	自底向上

时间复杂度：
空间复杂度：
*/
func longestCommonSubsequenceP2(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxN(dp[i-1][j-1], dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

/*
方法一: 动态规划
	自底向上 状态压缩

时间复杂度：
空间复杂度：
*/
func longestCommonSubsequenceP3(text1 string, text2 string) int {
	// dp := make([][]int, len(text1)+1)
	// for i := 0; i < len(dp); i++ {
	// 	dp[i] = make([]int, len(text2)+1)
	// }
	d := make([]int, len(text2)+1)
	for i := 1; i <= len(text1); i++ {
		pre := 0 // 这里 表示 dp[i][0] = 0 !!!!
		for j := 1; j <= len(text2); j++ {
			temp := d[j] // ta 就是 dp[i-1][j], 所有 [i-1][x] 的都是上一轮的值，d[j] 未更新前就是 d[i-1][j]
			if text1[i-1] == text2[j-1] {
				// dp[i][j] = dp[i-1][j-1] + 1
				d[j] = pre + 1 // 这里的 pre 是上一轮 j 循环里的未更新 d[j], 也就是上一轮j循环里的 d[i-1][j-1] !!!
			} else {
				// dp[i][j] = maxN(dp[i-1][j-1], dp[i][j-1], dp[i-1][j])
				d[j] = maxN(d[j-1], d[j])
			}
			pre = temp
		}
	}
	// return dp[len(text1)][len(text2)]
	return d[len(text2)]
}
