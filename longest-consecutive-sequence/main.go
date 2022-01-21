/*
Package longest_consecutive_sequence
https://leetcode-cn.com/problems/longest-consecutive-sequence/

128. 最长连续序列

给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

提示：
	0 <= nums.length <= 10^5
	-10^9 <= nums[i] <= 10^9

*/
package longest_consecutive_sequence

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func longestConsecutive(nums []int) int {
	h := make(map[int]int)
	for _, n := range nums {
		h[n]++
	}
	ans := 0
	for n, c := range h {
		if c == 0 {
			continue
		}
		count := 1
		for p := n + 1; h[p] > 0; {
			count++
			h[p] = 0
			p++
		}
		for p := n - 1; h[p] > 0; {
			count++
			h[p] = 0
			p--
		}

		if count > ans {
			ans = count
		}
	}

	return ans
}
