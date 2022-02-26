/*
Package super_egg_drop
https://leetcode-cn.com/problems/super-egg-drop/

887. 鸡蛋掉落

给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。
已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。
每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，
则可以在之后的操作中 重复使用 这枚鸡蛋。
请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？

提示：
	1 <= k <= 100
	1 <= n <= 10^4

*/
package super_egg_drop

import (
	"fmt"
	"math"
)

// --- 自己

/*
方法一: 动态规划
	!!! 超出时间限制

时间复杂度：
空间复杂度：
*/
func superEggDrop(k int, n int) int {
	// k 剩余鸡蛋 n 剩余楼层数
	var dp func(k, n int) int
	memo := make(map[string]int)
	dp = func(k, n int) int {
		ans := math.MaxInt64
		if n <= 0 { // 楼层为0 不用扔了
			return 0
		}

		if k <= 0 { // 没有鸡蛋，楼层大于0 不能完成测试了
			return -1
		}

		if k == 1 { // 只有一个鸡蛋，只能1层1层扔了
			return n
		}
		if ans, ok := memo[fmt.Sprintf("%d,%d", k, n)]; ok {
			return ans
		}

		for i := 1; i <= n; i++ {
			ans = min(ans, max(dp(k-1, i-1), dp(k, n-i))+1)
		}
		memo[fmt.Sprintf("%d,%d", k, n)] = ans
		return ans
	}

	ans := dp(k, n)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// --- 官方

/*
方法一: 动态规划 + 二分查找

时间复杂度：O(knlogn)。我们需要计算 O(kn) 个状态，每个状态计算时需要 O(logn) 的时间进行二分查找。
空间复杂度：O(kn)。我们需要 O(kn) 的空间存储每个状态的解。
*/
func superEggDropO1(k int, n int) int {
	memo := make(map[int]int)
	var dp func(k, n int) int
	dp = func(k, n int) int {
		if ans, ok := memo[n*1000+k]; ok {
			return ans
		}
		ans := 0
		if n == 0 {
			ans = 0
		} else if k == 1 {
			return n
		} else {
			lo, hi := 1, n
			for lo+1 < hi {
				mid := (lo + hi) / 2
				t1 := dp(k-1, mid-1)
				t2 := dp(k, n-mid)
				if t1 < t2 {
					lo = mid
				} else if t1 > t2 {
					hi = mid
				} else {
					lo, hi = mid, mid
				}
			}
			ans = 1 + min(max(dp(k-1, lo-1), dp(k, n-lo)), max(dp(k-1, hi-1), dp(k, n-hi)))
		}
		memo[n*1000+k] = ans
		return ans
	}

	return dp(k, n)
}

/*
方法二: 决策单调性

时间复杂度：O(kn)。我们需要计算 O(kn) 个状态，同时对于每个 k，最优解指针只会从 0 到 n 走一次，复杂度也是 O(kn)。因此总体复杂度为 O(kn)。
空间复杂度：O(n)。因为 dp 每一层的解只依赖于上一层的解，因此我们每次只保留一层的解，需要的空间复杂度为 O(n)。
*/
func superEggDropO2(k int, n int) int {
	// Right now, dp[i] represents dp(1, i)
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i
	}
	for j := 2; j <= k; j++ {
		// Now, we will develop dp2[i] = dp(j, i)
		dp2 := make([]int, n+1)
		x := 1
		for m := 1; m <= n; m++ {
			// Let's find dp2[m] = dp(j, m)
			// Increase our optimal x while we can make our answer better.
			// Notice max(dp[x-1], dp2[m-x]) > max(dp[x], dp2[m-x-1])
			// is simply max(T1(x-1), T2(x-1)) > max(T1(x), T2(x)).
			for x < m && max(dp[x-1], dp2[m-x]) > max(dp[x], dp2[m-x-1]) {
				x++
			}

			// The final answer happens at this x.
			dp2[m] = 1 + max(dp[x-1], dp2[m-x])
		}

		dp = dp2
	}

	return dp[n]
}

/*
方法二: 数学法

时间复杂度：O(kn)。事实上，更准确的时间复杂度应当为 O(kt)，我们不加证明地给出 n = O(t^k)，因此有 O(kt) = O(k\sqrt[k]{n})
空间复杂度：O(kn)
*/
func superEggDropO3(k int, n int) int {
	if n == 1 {
		return 1
	}
	f := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]int, k+1)
	}
	for i := 1; i <= k; i++ {
		f[1][i] = 1
	}
	ans := -1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			f[i][j] = 1 + f[i-1][j-1] + f[i-1][j]
		}
		if f[i][k] >= n {
			ans = i
			break
		}
	}
	return ans
}

// --- 他人

/*
方法二: 动态规划 + 递归

时间复杂度：
空间复杂度：
*/
func superEggDropP1(k int, n int) int {
	var calF func(k, t int) int
	calF = func(k, t int) int {
		// 如果只有一次尝试机会或者鸡蛋只有一个，则只能确定尝试机会+1数量的楼层哪一层鸡蛋会碎
		if t == 1 || k == 1 {
			return t + 1
		}
		// 非上述情况则递归（鸡蛋-1，则机会-1）和（鸡蛋没碎，机会-1）
		return calF(k-1, t-1) + calF(k, t-1)
	}

	t := 1
	// 问题 有 t 次尝试机会， k 个鸡蛋 能确定的最高楼层
	for calF(k, t) < n+1 { // 凑够n层的时候输出尝试的次数
		t++
	}
	return t
}
