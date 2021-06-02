/*
Package remove_element
https://leetcode-cn.com/problems/remove-element/

27. 移除元素

给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

提示：
	0 <= nums.length <= 100
	0 <= nums[i] <= 50
	0 <= val <= 100

*/
package remove_element

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func removeElement(nums []int, val int) int {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}
