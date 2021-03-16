/*
https://leetcode-cn.com/problems/target-sum/

494. 目标和

给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。
对于数组中的任意一个整数，你都可以从 + 或 - 中选择一个符号添加在前面。
返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

提示：
	数组非空，且长度不会超过 20 。
	初始的数组的和不会超过 1000 。
	保证返回的最终结果能被 32 位整数存下。

*/
package target_sum

import (
	"fmt"
)

// --- 自己

/*
方法一: 回溯

时间复杂度：O(2^N)，其中 N 是数组 nums 的长度。
空间复杂度：O(N)，为递归使用的栈空间大小。
*/
func findTargetSumWays(nums []int, S int) int {
	ans := 0
	var backTrack func(idx, sum int)
	backTrack = func(idx, sum int) {
		if idx >= len(nums) {
			if sum == S {
				ans++
			}
			return
		}
		// for i := 0; i < 2; i++ {
		// 	// if i == 0 {
		// 	// 	sum -= nums[idx]
		// 	// } else {
		// 	// 	sum += nums[idx]
		// 	// }
		// 	// backTrack(idx+1, sum)
		// 	// // !!! 要将sum 还原
		// 	// if i == 0 {
		// 	// 	sum += nums[idx]
		// 	// } else {
		// 	// 	sum -= nums[idx]
		// 	// }
		// 	if i == 0 {
		// 		backTrack(idx+1, sum-nums[idx])
		// 	} else {
		// 		backTrack(idx+1, sum+nums[idx])
		// 	}
		// }

		backTrack(idx+1, sum-nums[idx])
		backTrack(idx+1, sum+nums[idx])
	}

	backTrack(0, 0)
	return ans
}

/*
方法二: 动态规划
	labuladong

时间复杂度：
空间复杂度：
*/
func findTargetSumWaysS2(nums []int, S int) int {
	var dp func(i, rest int) int
	memo := make(map[string]int)
	dp = func(i, rest int) int {
		key := fmt.Sprintf("%d,%d", i, rest)
		if n, ok := memo[key]; ok {
			return n
		}

		if i == len(nums) {
			if rest == 0 {
				return 1
			}
			return 0
		}

		res := dp(i+1, rest-nums[i]) + dp(i+1, rest+nums[i])

		memo[key] = res

		return res
	}

	return dp(0, S)
}

/*
方法二: 动态规划优化
	labuladong
	A: 分配+号 B: 分配-号
	sum(A) - sum(B) = target
	sum(A) = target + sum(B)
	sum(A) + sum(A) = target + sum(B) + sum(A)
	2 * sum(A) = target + sum(nums)
	转为背包问题: nums 中存在几个子集 A，使得 A 中元素的和为 (target + sum(nums)) / 2？

时间复杂度：
空间复杂度：
*/
func findTargetSumWaysS3(nums []int, S int) int {
	n := sum(nums)
	if S > n || (S+n)%2 != 0 {
		return 0
	}
	// return subsets1(nums, (S+n)/2)
	return subsets(nums, (S+n)/2)
}

func sum(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans += v
	}

	return ans
}

func subsets(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= len(nums); i++ {
		for j := target; j >= 0; j-- {
			if j < nums[i-1] {
				dp[j] = dp[j]
			} else {
				dp[j] = dp[j] + dp[j-nums[i-1]]
			}
		}
	}
	return dp[target]
}

func subsets1(nums []int, target int) int {
	dp := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1
	}

	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= target; j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][target]
}
