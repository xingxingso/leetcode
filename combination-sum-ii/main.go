/*
Package combination_sum_ii
https://leetcode-cn.com/problems/combination-sum-ii/

40. 组合总和 II

给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

注意：
	解集不能包含重复的组合。

提示:
	1 <= candidates.length <= 100
	1 <= candidates[i] <= 50
	1 <= target <= 30

*/
package combination_sum_ii

import (
	"sort"
)

// --- 官方

/*
方法一: 回溯

时间复杂度：O(2^n x n)，其中 n 是数组 candidates 的长度。
空间复杂度：O(n)。
*/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	// 统计每个数字出现的次数
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	ans := make([][]int, 0)
	var sequence []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		// pos 当前这个不取
		dfs(pos+1, rest)

		// 至少取1个
		most := min(rest/freq[pos][0], freq[pos][1]) // 最多能用多少个, 有多少个重复的
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		// 回溯
		sequence = sequence[:len(sequence)-most]
	}

	dfs(0, target)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
