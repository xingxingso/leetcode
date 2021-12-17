/*
Package wiggle_subsequence
https://leetcode-cn.com/problems/wiggle-subsequence/

376. 摆动序列

如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。

例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。

相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。

给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。

提示：
	1 <= nums.length <= 1000
	0 <= nums[i] <= 1000

进阶：
	你能否用 O(n) 时间复杂度完成此题?

*/
package wiggle_subsequence

// --- 他人

/*
方法一:

时间复杂度：
空间复杂度：
*/
func wiggleMaxLength(nums []int) int {
	flag := 1

	pre := nums[0]
	start := 1
	for ; start < len(nums); start++ {
		if nums[start] != pre {
			break
		}
	}
	if start < len(nums) && nums[start]-pre < 0 {
		flag = -1
	}
	count := 1
	for ; start < len(nums); start++ {
		if (nums[start]-pre)*flag > 0 {
			flag *= -1
			count++
		}
		pre = nums[start]
	}
	return count
}
