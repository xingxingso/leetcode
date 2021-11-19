/*
Package sort_colors
https://leetcode-cn.com/problems/sort-colors/

75. 颜色分类

给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

提示：
	n == nums.length
	1 <= n <= 300
	nums[i] 为 0、1 或 2

进阶：
	你可以不使用代码库中的排序函数来解决这道题吗？
	你能想出一个仅使用常数空间的一趟扫描算法吗？

*/
package sort_colors

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func sortColors(nums []int) {
	left, right := 0, len(nums)-1

	for i := 0; i <= right; {
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		} else if nums[i] == 2 {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		} else {
			i++
		}
	}
}
