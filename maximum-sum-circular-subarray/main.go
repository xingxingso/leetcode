/*
Package maximum_sum_circular_subarray
https://leetcode-cn.com/problems/maximum-sum-circular-subarray/

918. 环形子数组的最大和

给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。

环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ，
nums[i] 的前一个元素是 nums[(i - 1 + n) % n] 。

子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，
不存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。

提示：
	n == nums.length
	1 <= n <= 3 * 10^4
	-3 * 10^4 <= nums[i] <= 3 * 10^4

*/
package maximum_sum_circular_subarray

import (
	"math"
)

// --- 自己

/*
方法一: 动态规划
	对数组进行扩展，这样问题简化
	// runtime: out of memory

时间复杂度：
空间复杂度：
*/
func maxSubarraySumCircularS1(nums []int) int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		nums = append(nums, nums[i])
	}
	//fmt.Println(nums)

	dp := make([][]int, l+1)
	// 以 i 开头， 长度为 j 的和
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, l+1)
	}
	ans := math.MinInt64
	for i := 1; i <= l; i++ {
		for j := i; j < l+i; j++ {
			dp[i][j-i+1] = dp[i][j-i] + nums[j-1]
			if dp[i][j-i+1] > ans {
				ans = dp[i][j-i+1]
			}
		}
	}
	//fmt.Println(dp)

	return ans
}

/*
方法一: 动态规划
	// 超出时限

时间复杂度：
空间复杂度：
*/
func maxSubarraySumCircularS2(nums []int) int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		nums = append(nums, nums[i])
	}
	//fmt.Println(nums)

	// 长度为 j 的最大和
	dp := make([]int, l+1)
	ans := math.MinInt64
	for i := 1; i <= l; i++ {
		for j := i; j < l+i; j++ {
			dp[j-i+1] = dp[j-i] + nums[j-1]
			if dp[j-i+1] > ans {
				ans = dp[j-i+1]
			}
		}
	}
	//fmt.Println(dp)

	return ans
}

// --- 官方

/*
方法一: 邻接数组

时间复杂度：O(N)，其中 N 是 A 的长度。
空间复杂度：O(N)。
*/
func maxSubarraySumCircularO1(nums []int) int {
	// 单区间数组
	cur, ans := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		//cur = max(nums[i], cur+nums[i])
		cur = nums[i] + max(cur, 0) // 这样可以避免加法
		ans = max(ans, cur)
	}

	// 双区间
	// For each i, we want to know
	// the maximum of sum(A[j:]) with j >= i+2
	// rightSums[i] = A[i] + A[i+1] + ... + A[N-1]
	rightSums := make([]int, len(nums))
	rightSums[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		rightSums[i] = rightSums[i+1] + nums[i]
	}

	// maxRight[i] = max_{j >= i} rightSums[j]
	maxRight := make([]int, len(nums))
	maxRight[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], rightSums[i])
	}

	leftSum := 0
	for i := 0; i < len(nums)-2; i++ {
		leftSum += nums[i]
		ans = max(ans, leftSum+maxRight[i+2]) // maxRight[i+1] 则成了单区间
	}

	//fmt.Println(ans)

	return ans
}

/*
方法二: 前缀和 + 单调队列

时间复杂度：O(N)，其中 N 是 A 的长度。
空间复杂度：O(N)。
*/
func maxSubarraySumCircularO2(nums []int) int {
	n := len(nums)

	// Compute P[j] = B[0] + B[1] + ... + B[j-1]
	// for fixed array B = A+A
	p := make([]int, 2*n+1)
	for i := 0; i < len(p)-1; i++ {
		p[i+1] = p[i] + nums[i%n]
	}
	//fmt.Println(p)

	// Want largest P[j] - P[i] with 1 <= j-i <= N
	// For each j, want smallest P[i] with j-N<= i < j
	ans := nums[0]
	// deque: i's, increasing by P[i]
	deque := make([]int, 0)
	deque = append(deque, 0)

	for j := 1; j < len(p); j++ {
		// If the smallest i is too small, remove it.
		if deque[0] < j-n {
			deque = deque[1:]
		}

		// The optimal i is deque[0], for cand. answer P[j] - P[i].
		ans = max(ans, p[j]-p[deque[0]])

		// Remove any i1's with P[i2] <= P[i1].
		for len(deque) > 0 && p[j] <= p[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, j)
		//fmt.Println(deque)
	}

	return ans
}

/*
方法三: Kadane 算法（符号变种）

时间复杂度：O(N)，其中 N 是 A 的长度。
空间复杂度：O(1)。
*/
func maxSubarraySumCircularO3(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ans1 := kadane(nums, 1)
	ans2 := sum + kadane(nums[1:], -1)
	ans3 := sum + kadane(nums[:len(nums)-1], -1)
	return max(ans1, max(ans2, ans3))
}

func kadane(nums []int, sign int) int {
	// The maximum non-empty subarray for array
	// [sign * A[i], sign * A[i+1], ..., sign * A[j]]
	ans, cur := math.MinInt64/2, math.MinInt64/2
	for i := 0; i < len(nums); i++ {
		cur = sign*nums[i] + max(cur, 0)
		ans = max(ans, cur)
	}
	return ans
}

/*
方法四: Kadane 算法（最小值变种）

时间复杂度：O(N)，其中 N 是 A 的长度。
空间复杂度：O(1)。
*/
func maxSubarraySumCircularO4(A []int) int {
	sum := 0
	for _, v := range A {
		sum += v
	}

	// ans1: answer for one-interval subarray
	ans1 := math.MinInt64 / 2
	cur := math.MinInt64 / 2
	for _, v := range A {
		cur = v + max(cur, 0)
		ans1 = max(ans1, cur)
	}

	// ans2: answer for two-interval subarray, interior in A[1:]
	ans2 := math.MaxInt64 / 2
	cur = math.MaxInt64 / 2
	for i := 1; i < len(A); i++ {
		cur = A[i] + min(cur, 0)
		ans2 = min(ans2, cur)
	}
	ans2 = sum - ans2

	// ans3: answer for two-interval subarray, interior in A[:len(A)-1]
	ans3 := math.MaxInt64 / 2
	cur = math.MaxInt64 / 2
	for i := 0; i < len(A)-1; i++ {
		cur = A[i] + min(cur, 0)
		ans3 = min(ans3, cur)
	}
	ans3 = sum - ans3

	return max(ans1, max(ans2, ans3))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
