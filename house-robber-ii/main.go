/*
Package house_robber_ii
https://leetcode-cn.com/problems/house-robber-ii/

213. 打家劫舍 II

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，
这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。

提示：
	0 <= nums.length <= 100
	0 <= nums[i] <= 1000

*/
package house_robber_ii

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	// 三种情况，第一家抢最后一家不抢，第一家不抢最后一家抢，第一家最后一家都不抢，
	// 最后一种情况肯定不是最大值
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange(nums []int, start, end int) int {
	first, second, profit := 0, 0, 0

	for i := end; i >= start; i-- {
		profit = max(first, nums[i]+second)
		second = first
		first = profit
	}

	return profit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
