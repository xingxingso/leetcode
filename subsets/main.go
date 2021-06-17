/*
Package subsets
https://leetcode-cn.com/problems/subsets/

78. 子集

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

提示：
    1 <= nums.length <= 10
    -10 <= nums[i] <= 10
    nums 中的所有元素 互不相同

*/
package subsets

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	q := make([]int, 0)
	var backTrack func(idx int)
	l := len(nums)
	backTrack = func(idx int) {
		res := make([]int, len(q))
		copy(res, q)
		ans = append(ans, res)
		for i := idx; i < l; i++ {
			q = append(q, nums[i])
			backTrack(i + 1)
			q = q[:len(q)-1]
		}
	}
	backTrack(0)
	return ans
}
