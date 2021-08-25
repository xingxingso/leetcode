/*
Package split_array_into_consecutive_subsequences
https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/

659. 分割数组为连续子序列

给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个长度至少为 3 的子序列，其中每个子序列都由连续整数组成。
如果可以完成上述分割，则返回 true ；否则，返回 false 。

提示：
	1 <= nums.length <= 10000

*/
package split_array_into_consecutive_subsequences

// --- 他人
//https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247491005&idx=1&sn=36cdcb0098aca81c3c4061baf2474b82&scene=21#wechat_redirect

/*
方法一:

时间复杂度：
空间复杂度：
*/
func isPossible(nums []int) bool {
	freq, need := make(map[int]int), make(map[int]int)
	// 统计 nums 中元素的频率
	for _, v := range nums {
		freq[v]++
	}

	for _, v := range nums {
		if freq[v] == 0 { // 已经被用到其他子序列中
			continue
		}
		// 先判断 v 是否能接到其他子序列后面
		if need[v] > 0 { // v 可以接到之前的某个序列后面
			freq[v]--
			need[v]--   // 对 v 的需求减1
			need[v+1]++ // 对 v+1 的需求加1
		} else if freq[v] > 0 && freq[v+1] > 0 && freq[v+2] > 0 { // 将 v 作为开头，新建一个长度为 3 的子序列 [v,v+1,v+2]
			freq[v]--
			freq[v+1]--
			freq[v+2]--
			need[v+3]++ // 对 v+3 的需求加1
		} else { // 两种情况都不符合，则无法分配
			return false
		}
	}
	return true
}
