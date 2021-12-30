/*
Package minimum_cost_to_merge_stones
https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/

1000. 合并石头的最低成本

有 N 堆石头排成一排，第 i 堆中有 stones[i] 块石头。

每次移动（move）需要将连续的 K 堆石头合并为一堆，而这个移动的成本为这 K 堆石头的总数。

找出把所有石头合并成一堆的最低成本。如果不可能，返回 -1 。

提示：
	1 <= stones.length <= 30
	2 <= K <= 30
	1 <= stones[i] <= 100

*/
package minimum_cost_to_merge_stones

import (
	"math"
)

// --- 自己

/*
// 错误，简单贪心不能通过测试 ex4
方法一: 贪心
	总是选择 连续和 最低的作为合并对象， 有点像构建 huffman 树
	连续和 求和技巧， sum(i+1)=sum(i)-stones[i]+stones[i+1+k-1], i为起始下标

时间复杂度：
空间复杂度：
*/
func mergeStonesS1(stones []int, k int) int {
	if t := len(stones) - k + 1; t < k && t > 1 { // 过滤
		return -1
	}

	var dfs func(stones []int) int
	dfs = func(stones []int) int {
		if len(stones) == 1 {
			return 0
		}
		if len(stones) < k {
			return -1
		}
		res := 0

		pre := 0
		for i := 0; i < k; i++ {
			pre += stones[i]
		}
		minIdx := 0
		min := pre
		for i := 1; i <= len(stones)-k; i++ {
			cur := pre + stones[i+k-1] - stones[i-1]
			if cur < min {
				minIdx = i
				min = cur
			}
			pre = cur
		}
		res += min
		// 这里将求和后的数组压缩到原数组中
		for i := 0; i < len(stones); i++ {
			if i < minIdx || (i > minIdx && i < minIdx+k) {
				continue
			} else if i == minIdx {
				stones[i] = min
			} else {
				stones[i-k+1] = stones[i]
			}
		}

		//fmt.Println("min", min, minIdx, stones[:len(stones)-k+1])
		n := dfs(stones[:len(stones)-k+1])
		if n == -1 {
			return -1
		}
		//fmt.Println(res, n)
		return res + n
	}
	return dfs(stones)
}

/*
方法二: 暴力解法
	超出时限

时间复杂度：
空间复杂度：
*/
func mergeStonesS2(stones []int, k int) int {
	if t := len(stones) - k + 1; t < k && t > 1 { // 过滤
		return -1
	}

	var dfs func(stones []int) int
	dfs = func(stones []int) int {
		if len(stones) == 1 {
			return 0
		}
		if len(stones) < k {
			return -1
		}

		min := math.MaxInt64
		newStones := make([]int, len(stones)-k+1)
		for i := 0; i <= len(stones)-k; i++ {
			res := 0
			newStones[i] = 0
			for j := 0; j < len(stones); j++ {
				if j < i {
					newStones[j] = stones[j]
				} else if j >= i && j < i+k {
					newStones[i] += stones[j]
				} else {
					newStones[j-k+1] = stones[j]
				}
			}
			//fmt.Println("newStones", newStones)
			nextSum := dfs(newStones)
			if nextSum == -1 {
				return -1
			}
			res = nextSum + newStones[i]
			if res < min {
				min = res
			}
		}

		return min
	}
	return dfs(stones)
}

// --- 他人

/*
方法一: 动态规划

	(5) JAVA Bottom-Up + Top-Down DP With Explaination - LeetCode Discuss
	https://leetcode.com/problems/minimum-cost-to-merge-stones/discuss/247657/JAVA-Bottom-Up-%2B-Top-Down-DP-With-Explaination

	（详解）由易到难，一步步说明思路和细节 - 合并石头的最低成本 - 力扣（LeetCode）
	https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/solution/yi-dong-you-yi-dao-nan-yi-bu-bu-shuo-ming-si-lu-he/

	【C++】AC的第1000题献给第1000题，经典区间DP - 合并石头的最低成本 - 力扣（LeetCode）
	https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/solution/c-acde-di-1000ti-xian-gei-di-1000ti-jing-ecs9/


时间复杂度：
空间复杂度：
*/
func mergeStonesP1(stones []int, k int) int {
	if (len(stones)-1)%(k-1) != 0 {
		return -1
	}

	prefixSum := make([]int, len(stones)+1)
	for i := 1; i <= len(stones); i++ { // 前缀和 用于求 sum[i,j] = prefixSum[j+1]- prefixSum[i]
		prefixSum[i] = prefixSum[i-1] + stones[i-1]
	}
	//fmt.Println(prefixSum)

	// dp[i][j][m] // [i,j] 区间石头合并成 m 堆的最小成本
	// 最终求 dp[0][len(stones)-1][1]
	// dp[i][i][1] = 0
	// dp[i][j][1] = dp[i][p][1] + dp[p+1][j][k-1] + sum[i,j]

	max := math.MaxInt64
	dp := make([][][]int, len(stones))
	for i := 0; i < len(stones); i++ {
		dp[i] = make([][]int, len(stones))
		for j := 0; j < len(stones); j++ {
			dp[i][j] = make([]int, k+1) // dp[i][j][0] 用不到
			for m := 0; m <= k; m++ {
				dp[i][j][m] = max
			}
		}
		dp[i][i][1] = 0
	}

	for l := 2; l <= len(stones); l++ { // 控制区间长度
		for i := 0; i <= len(stones)-l; i++ {
			j := i + l - 1
			for m := 2; m <= k; m++ { // 堆数
				for t := i; t < j; t++ { // 切分位置
					if dp[i][t][m-1] == max || dp[t+1][j][1] == max {
						continue
					}
					dp[i][j][m] = min(dp[i][j][m], dp[i][t][m-1]+dp[t+1][j][1])
				}
			}
			if dp[i][j][k] == max {
				continue
			}
			dp[i][j][1] = dp[i][j][k] + prefixSum[j+1] - prefixSum[i]
		}

	}
	//fmt.Println(dp)
	return dp[0][len(stones)-1][1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
