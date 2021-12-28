/*
Package combination_sum
https://leetcode-cn.com/problems/combination-sum/

39. 组合总和

给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有不同组合 ，
并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 target 的不同组合数少于 150 个。

提示：
	1 <= candidates.length <= 30
	1 <= candidates[i] <= 200
	candidate 中的每个元素都 互不相同
	1 <= target <= 500

*/
package combination_sum

// --- 自己

/*
方法一: 回溯

时间复杂度：
空间复杂度：
*/
func combinationSum(candidates []int, target int) [][]int {
	current := make([]int, 0)
	ans := make([][]int, 0)
	var backtrack func(target, pos int)
	backtrack = func(target, pos int) {
		if target == 0 {
			temp := make([]int, len(current))
			copy(temp, current)
			ans = append(ans, temp)
			return
		}
		for i := pos; i < len(candidates); i++ {
			if candidates[i] > target {
				continue
			}

			n := target / candidates[i]
			for j := 1; j <= n; j++ {
				current = append(current, candidates[i])
				backtrack(target-j*candidates[i], i+1)
			}
			current = current[:len(current)-n]
		}
	}

	if target > 0 {
		backtrack(target, 0)
	}

	return ans
}
