/*
Package shu_zu_zhong_zhong_fu_de_shu_zi_lcof
https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

剑指 Offer 03. 数组中重复的数字

找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字。

限制：
	2 <= n <= 100000

*/
package shu_zu_zhong_zhong_fu_de_shu_zi_lcof

// --- 自己

/*
方法一:
	nums[i]=n 表示 i 这个数字出现了 -n 次

时间复杂度：O(n)
空间复杂度：O(1)
*/
func findRepeatNumber(nums []int) int {
	max := 100000 // 不可能出现的值，2 <= n <= 100000
	for i := 0; i < len(nums); {
		p := nums[i]
		if p == max || p < 0 {
			i++
			continue
		}
		if nums[p] >= 0 { // 此位置上已被他人占据
			nums[i] = nums[p]
			nums[p] = -1 // 出现一次
		} else {
			return p // 这里只要求找出一个
			//nums[p]--     // 增加一次
			//nums[i] = max
		}
	}
	for k, v := range nums {
		if v < -1 {
			return k
		}
	}
	//fmt.Println(nums)
	return -1
}
