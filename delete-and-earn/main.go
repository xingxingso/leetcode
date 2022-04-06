/*
Package delete_and_earn
https://leetcode-cn.com/problems/delete-and-earn/

740. 删除并获得点数

给你一个整数数组 nums ，你可以对它进行一些操作。

每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。

开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。

提示：
	1 <= nums.length <= 2 * 10^4
	1 <= nums[i] <= 10^4

*/
package delete_and_earn

import "sort"

// --- 官方

/*
方法一: 动态规划

时间复杂度：O(NlogN)，其中 NN 是数组 nums 的长度。对 nums 排序需要花费 O(NlogN) 的时间，
	遍历计算需要花费 O(N) 的时间，故总的时间复杂度为 O(NlogN)。
空间复杂度：O(N)。统计 sum 至多需要花费 O(N) 的空间。
*/
func deleteAndEarn(nums []int) int {
	sort.Ints(nums)
	sum := []int{nums[0]}
	ans := 0
	for i := 1; i < len(nums); i++ {
		if val := nums[i]; val == nums[i-1] {
			sum[len(sum)-1] += val
		} else if val == nums[i-1]+1 {
			sum = append(sum, val)
		} else {
			ans += rob(sum)
			sum = []int{val}
		}
	}
	ans += rob(sum)
	return ans
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
