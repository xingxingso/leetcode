/*
Package squares_of_a_sorted_array
https://leetcode-cn.com/problems/squares-of-a-sorted-array/

977. 有序数组的平方

给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

提示：
	1 <= nums.length <= 10^4
	-10^4 <= nums[i] <= 10^4
	nums 已按 非递减顺序 排序

进阶：
	请你设计时间复杂度为 O(n) 的算法解决本问题

*/
package squares_of_a_sorted_array

// --- 自己

/*
方法一：

时间复杂度：
空间复杂度：
*/
func sortedSquares(nums []int) []int {
	pos := 0
	for ; pos < len(nums); pos++ {
		if nums[pos] > 0 {
			break
		}
	}
	res := make([]int, 0)
	for l, r := pos-1, pos; l >= 0 || r < len(nums); {
		if l < 0 {
			res = append(res, nums[r]*nums[r])
			r++
		} else if r >= len(nums) {
			res = append(res, nums[l]*nums[l])
			l--
		} else if nums[r] > -nums[l] {
			res = append(res, nums[l]*nums[l])
			l--
		} else {
			res = append(res, nums[r]*nums[r])
			r++
		}
	}

	return res
}

// --- 官方

/*
方法一：双指针

时间复杂度：O(n)，其中 n 是数组 nums 的长度。
空间复杂度：O(1)。除了存储答案的数组以外，我们只需要维护常量空间。
*/
func sortedSquaresO1(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := nums[i]*nums[i], nums[j]*nums[j]; v > w {
			ans[pos] = v
			i++
		} else {
			ans[pos] = w
			j--
		}
	}
	return ans
}
