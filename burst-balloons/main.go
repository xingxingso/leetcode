/*
https://leetcode-cn.com/problems/burst-balloons/

312. 戳气球

有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。
这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。

求所能获得硬币的最大数量。

提示：
	n == nums.length
	1 <= n <= 500
	0 <= nums[i] <= 100

*/
package burst_balloons

// --- 自己

/*
方法一: 动态规划
	！！！超出时间限制

时间复杂度：O(n!)
空间复杂度：O(n)
*/
func maxCoins(nums []int) int {
	var dp func(nums []int) int
	dp = func(nums []int) int {
		ans := 0
		for i := 0; i < len(nums); i++ {
			cal := nums[i]
			if i > 0 && i < len(nums)-1 {
				cal *= nums[i-1] * nums[i+1]
			} else if i == 0 && i+1 < len(nums) {
				cal *= nums[i+1]
			} else if i > 0 && i == len(nums)-1 {
				cal *= nums[i-1]
			}
			newNum := make([]int, 0)
			newNum = append(newNum, nums[:i]...)
			newNum = append(newNum, nums[i+1:]...)
			// if len(newNum) == 1 { // 可以稍微提高点效率
			// 	cal += newNum[0]
			// } else {
			cal += dp(newNum)
			// }
			ans = max(ans, cal)
		}
		return ans
	}

	return dp(nums)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法一: 动态规划优化
	对于1个数，总是其自身，
	对于2个数，总是优先射小数
	对于3个数abc,
		1. ab+dp(bc)  = ab +bc+max(b,c)
		2. bc+dp(ab)  = bc +ab+max(a,b)
		3. abc+dp(ac) = abc+ac+max(a,c)

时间复杂度：
空间复杂度：
*/
func maxCoinsS2(nums []int) int {
	var dp func(nums []int) int
	dp = func(nums []int) int {
		ans := 0
		for i := 0; i < len(nums); i++ {
			cal := nums[i]
			if i > 0 && i < len(nums)-1 {
				cal *= nums[i-1] * nums[i+1]
			} else if i == 0 && i+1 < len(nums) {
				cal *= nums[i+1]
			} else if i > 0 && i == len(nums)-1 {
				cal *= nums[i-1]
			}
			newNum := make([]int, 0)
			newNum = append(newNum, nums[:i]...)
			newNum = append(newNum, nums[i+1:]...)
			// 可以稍微提高点效率
			if len(newNum) == 2 {
				cal += newNum[0]*newNum[1] + max(newNum[0], newNum[1])
			} else if len(newNum) == 1 {
				cal += newNum[0]
			} else {
				cal += dp(newNum)
			}
			ans = max(ans, cal)
		}
		return ans
	}

	return dp(nums)
}

/*
方法一: 记忆化搜索

时间复杂度：O(n^3)，其中 n 是气球数量。区间数为 n^2，区间迭代复杂度为 O(n)，最终复杂度为 O(n^2×n)=O(n^3)。
空间复杂度：O(n^2)，其中 n 是气球数量。缓存大小为区间的个数。
*/
func maxCoinsO1(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	val[0], val[n+1] = 1, 1
	rec := make([][]int, n+2)
	for i := 0; i < len(rec); i++ {
		rec[i] = make([]int, n+2)
		for j := 0; j < len(rec[i]); j++ {
			rec[i][j] = -1
		}
	}
	return solve(0, n+1, val, rec)
}

func solve(left, right int, val []int, rec [][]int) int {
	if left >= right-1 {
		return 0
	}
	if rec[left][right] != -1 {
		return rec[left][right]
	}
	for i := left + 1; i < right; i++ {
		sum := val[left] * val[i] * val[right]
		sum += solve(left, i, val, rec) + solve(i, right, val, rec)
		rec[left][right] = max(rec[left][right], sum)
	}
	return rec[left][right]
}

/*
方法二: 动态规划

时间复杂度：O(n^3)，其中 n 是气球数量。区间数为 n^2，区间迭代复杂度为 O(n)，最终复杂度为 O(n^2×n)=O(n^3)。
空间复杂度：O(n^2)，其中 n 是气球数量。
*/
func maxCoinsO2(nums []int) int {
	n := len(nums)
	rec := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		rec[i] = make([]int, n+2)
	}
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += rec[i][k] + rec[k][j]
				rec[i][j] = max(rec[i][j], sum)
			}
		}
	}
	return rec[0][n+1]
}

/*
方法一: 动态规划
	labuladong

时间复杂度：
空间复杂度：
*/
func maxCoinsP1(nums []int) int {
	// (i, j) 之间的气球能得到的分数， k 是最后一个被戳的气球 i < k < j
	// dp[i][j] = dp[i][k] + dp[k][j] + nums[i]*nums[j]*nums[k]
	n := len(nums)
	points := make([]int, n+2)
	points[0], points[n+1] = 1, 1
	copy(points[1:], nums)
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	for i := n; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
			}
		}
	}

	return dp[0][n+1]
}
