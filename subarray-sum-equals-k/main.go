/*
Package subarray_sum_equals_k
https://leetcode-cn.com/problems/subarray-sum-equals-k/

560. 和为K的子数组

给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

说明 :
	数组的长度为 [1, 20,000]。
	数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。

*/
package subarray_sum_equals_k

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484488&idx=1&sn=848f76e86fce722e70e265d0c6f84dc3&chksm=9bd7fa40aca07356a6f16db72f5a56529044b1bdb2dcce2de4efe59e0338f0c313de682aef29&scene=21#wechat_redirect

/*
方法一: 前缀和

时间复杂度：
空间复杂度：
*/
func subarraySumP1(nums []int, k int) int {
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if preSum[j+1]-preSum[i] == k {
				ans++
			}
		}
	}
	return ans
}

/*
方法一: 前缀和 优化src/cmd/compile/internal/gc/inl.go

时间复杂度：
空间复杂度：
*/
func subarraySumP2(nums []int, k int) int {
	preSum := make(map[int]int, len(nums)+1)
	preSum[0] = 1 // 必须有这个
	sum, ans := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		targetSum := sum - k
		if n, ok := preSum[targetSum]; ok {
			ans += n
		}
		preSum[sum]++
	}
	return ans
}
