/*
Package longest_arithmetic_subsequence_of_given_difference
https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/

1218. 最长定差子序列

给你一个整数数组 arr 和一个整数 difference，请你找出并返回 arr 中最长等差子序列的长度，该子序列中相邻元素之间的差等于 difference 。

子序列 是指在不改变其余元素顺序的情况下，通过删除一些元素或不删除任何元素而从 arr 派生出来的序列。

提示：
	1 <= arr.length <= 10^5
	-10^4 <= arr[i], difference <= 10^4

*/
package longest_arithmetic_subsequence_of_given_difference

// --- 自己

/*
方法一: 动态规划
	// 超出时限

时间复杂度：
空间复杂度：
*/
func longestSubsequenceS1(arr []int, difference int) int {
	dp := make([]int, len(arr)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i <= len(arr); i++ {
		for j := i + 1; j <= len(arr); j++ {
			if arr[j-1]-arr[i-1] == difference {
				//dp[j] = max(dp[j], dp[i]+1)
				dp[j] = dp[i] + 1
				ans = max(ans, dp[j])
			}
		}
	}
	//fmt.Println(dp)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// --- 官方

/*
方法一: 动态规划
	令 dp[i] 表示以 arr[i] 为结尾的最长的等差子序列的长度
 	j 可以取满足 j<i 且 arr[j]=arr[i]−d
	dp[i]=dp[j]+1
	dp[v]=dp[v−d]+1

时间复杂度：
空间复杂度：
*/
func longestSubsequenceS2(arr []int, difference int) int {
	dp := make(map[int]int)
	ans := 0
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if ans < dp[v] {
			ans = dp[v]
		}
	}
	//fmt.Println(dp)
	return ans
}
