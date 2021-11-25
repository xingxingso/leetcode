/*
Package permutations_ii
https://leetcode-cn.com/problems/permutations-ii/

47. 全排列 II

给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

提示：
	1 <= nums.length <= 8
	-10 <= nums[i] <= 10

*/
package permutations_ii

import "sort"

// --- 官方

/*
方法一:

时间复杂度：
空间复杂度：
*/
func permuteUniqueO1(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	ans := make([][]int, 0)

	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			// 112 中 第一个1 必须被先访问
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return ans
}
