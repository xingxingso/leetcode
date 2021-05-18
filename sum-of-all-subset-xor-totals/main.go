/*
https://leetcode-cn.com/problems/sum-of-all-subset-xor-totals/
https://leetcode-cn.com/contest/weekly-contest-241/problems/sum-of-all-subset-xor-totals/

5759. 找出所有子集的异或总和再求和

一个数组的 异或总和 定义为数组中所有元素按位 XOR 的结果；如果数组为 空 ，则异或总和为 0 。

例如，数组 [2,5,6] 的 异或总和 为 2 XOR 5 XOR 6 = 1 。
给你一个数组 nums ，请你求出 nums 中每个 子集 的 异或总和 ，计算并返回这些值相加之 和 。

注意：在本题中，元素 相同 的不同子集应 多次 计数。

数组 a 是数组 b 的一个 子集 的前提条件是：从 b 删除几个（也可能不删除）元素能够得到 a 。

提示：
	1 <= nums.length <= 12
	1 <= nums[i] <= 20

*/
package sum_of_all_subset_xor_totals

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func subsetXORSum(nums []int) int {
	var dfs func(cur int)
	var ans int
	set := make([]int, 0)
	dfs = func(cur int) {
		if cur == len(nums) {
			return
		}
		set = append(set, nums[cur])
		res := 0
		for i := 0; i < len(set); i++ {
			res ^= set[i]
		}
		ans += res

		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return ans
}
