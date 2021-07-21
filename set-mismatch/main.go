/*
Package set_mismatch
https://leetcode-cn.com/problems/set-mismatch/

645. 错误的集合

集合 s 包含从 1 到 n 的整数。不幸的是，因为数据错误，
导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有一个数字重复 。
给定一个数组 nums 代表了集合 S 发生错误后的结果。
请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

提示：
	2 <= nums.length <= 10^4
	1 <= nums[i] <= 10^4

*/
package set_mismatch

// --- 他人
//https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485050&idx=1&sn=dac757454b2df9a1291f1e8027f56c1b&scene=21#wechat_redirect

/*
方法一:
时间复杂度：
空间复杂度：
*/
func findErrorNums(nums []int) []int {
	mis, dup := -1, -1
	for i := 0; i < len(nums); i++ {
		idx := abs(nums[i]) - 1
		if nums[idx] < 0 {
			dup = idx + 1
		} else {
			nums[idx] = -nums[idx]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			mis = i + 1
			break
		}
	}
	return []int{dup, mis}
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
