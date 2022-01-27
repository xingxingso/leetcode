/*
Package super_ugly_number
https://leetcode-cn.com/problems/super-pow/

313. 超级丑数

超级丑数 是一个正整数，并满足其所有质因数都出现在质数数组 primes 中。

给你一个整数 n 和一个整数数组 primes ，返回第 n 个 超级丑数 。

题目数据保证第 n 个 超级丑数 在 32-bit 带符号整数范围内。

提示：
	1 <= n <= 10^6
	1 <= primes.length <= 100
	2 <= primes[i] <= 1000
	题目数据 保证 primes[i] 是一个质数
	primes 中的所有值都 互不相同 ，且按 递增顺序 排列

*/
package super_ugly_number

import "math"

// --- 官方

/*
方法一: 动态规划

时间复杂度：O(nm)，其中 n 是待求的超级丑数的编号，m 是数组 primes 的长度。需要计算数组 dp 中的 n 个元素，每个元素的计算都需要 O(m) 的时间。
空间复杂度：O(nm)，其中 n 是待求的超级丑数的编号，m 是数组 primes 的长度。需要计算数组 dp 中的 n 个元素，每个元素的计算都需要 O(m) 的时间。
*/
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	m := len(primes)
	pointers := make([]int, m)
	nums := make([]int, m)
	for i := range nums {
		nums[i] = 1
	}
	for i := 1; i <= n; i++ {
		minNum := math.MaxInt64
		for j := range pointers {
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j := range nums {
			if nums[j] == minNum {
				pointers[j]++
				nums[j] = dp[pointers[j]] * primes[j]
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
