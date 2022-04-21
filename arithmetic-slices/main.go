/*
Package arithmetic_slices
https://leetcode-cn.com/problems/arithmetic-slices/

413. 等差数列划分

如果一个数列 至少有三个元素 ，并且任意两个相邻元素之差相同，则称该数列为等差数列。

例如，[1,3,5,7,9]、[7,7,7,7] 和 [3,-1,-5,-9] 都是等差数列。
给你一个整数数组 nums ，返回数组 nums 中所有为等差数组的 子数组 个数。

子数组 是数组中的一个连续序列。

提示：
	1 <= nums.length <= 5000
	-1000 <= nums[i] <= 1000

*/
package arithmetic_slices

// --- 官方

/*
方法一:

时间复杂度：
空间复杂度：
*/
func numberOfArithmeticSlicesO1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	d, t := nums[0]-nums[1], 0
	ans := 0
	// 因为等差数列的长度至少为 3，所以可以从 i=2 开始枚举
	for i := 2; i < n; i++ {
		if nums[i-1]-nums[i] == d {
			t++
		} else {
			d, t = nums[i-1]-nums[i], 0
		}
		ans += t
	}
	return ans
}

// --- 他人

/*
方法一:
	Leetcode101 7章

时间复杂度：
空间复杂度：
*/
func numberOfArithmeticSlicesP1(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	dp := make([]int, n)
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	res := 0
	for _, v := range dp {
		res += v
	}
	return res
}
