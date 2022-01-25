/*
Package contains_duplicate
https://leetcode-cn.com/problems/contains-duplicate/

217. 存在重复元素

给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。

提示：
	1 <= nums.length <= 10^5
	-10^9 <= nums[i] <= 10^9

*/
package contains_duplicate

// --- 自己

/*
方法一: 哈希表

时间复杂度：
空间复杂度：
*/
func containsDuplicate(nums []int) bool {
	hash := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		if hash[nums[i]] {
			return true
		}
		hash[nums[i]] = true
	}

	return false
}
