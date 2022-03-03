/*
Package minimum_size_subarray_sum
https://leetcode-cn.com/problems/minimum-size-subarray-sum/

209. 长度最小的子数组

给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [nums_l, nums_l+1, ..., nums_r-1, nums_r] ，并返回其长度。
如果不存在符合条件的子数组，返回 0 。

提示：
	1 <= target <= 10^9
	1 <= nums.length <= 10^5
	1 <= nums[i] <= 10^5

进阶：
	如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。

*/
package minimum_size_subarray_sum

// --- 自己

/*
方法一: 滑动窗口

时间复杂度：O(n)，其中 n 是数组的长度。指针 start 和 end 最多各移动 n 次。
空间复杂度：O(1)。
*/
func minSubArrayLen(target int, nums []int) int {
	ans := 0
	sum := 0
	for l, r := -1, 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			if ans == 0 || r-l < ans {
				ans = r - l
			}
			l++
			sum -= nums[l]
		}
	}

	return ans
}
