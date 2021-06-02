/*
Package two_sum_ii_input_array_is_sorted
https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/

167. 两数之和 II - 输入有序数组

给定一个已按照 升序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。
函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数 ，
所以答案数组应当满足 1 <= answer[0] < answer[1] <= numbers.length 。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。

提示：
	2 <= numbers.length <= 3 * 10^4
	-1000 <= numbers[i] <= 1000
	numbers 按 递增顺序 排列
	-1000 <= target <= 1000
	仅存在一个有效答案

*/
package two_sum_ii_input_array_is_sorted

// --- 自己

/*
方法一: 双指针

时间复杂度：
空间复杂度：
*/
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{-1, -1}
}
