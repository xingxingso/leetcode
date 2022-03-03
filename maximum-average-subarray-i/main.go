/*
Package maximum_average_subarray_i
https://leetcode-cn.com/problems/maximum-average-subarray-i/

643. 子数组最大平均数 I

给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。

请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。

任何误差小于 10^-5 的答案都将被视为正确答案。

提示：
	n == nums.length
	1 <= k <= n <= 10^5
	-10^4 <= nums[i] <= 10^4

*/
package maximum_average_subarray_i

// --- 自己

/*
方法一: 滑动窗口

时间复杂度：
空间复杂度：
*/
func findMaxAverage(nums []int, k int) float64 {
	kthSum := 0
	for i := 0; i < k; i++ {
		kthSum += nums[i]
	}
	maxKthSum := kthSum
	for i := k; i < len(nums); i++ {
		kthSum = kthSum + nums[i] - nums[i-k]
		if kthSum > maxKthSum {
			maxKthSum = kthSum
		}
	}

	return float64(maxKthSum) / float64(k)
}
