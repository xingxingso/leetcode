/*
https://leetcode-cn.com/problems/3sum-smaller/

259. 较小的三数之和

给定一个长度为 n 的整数数组和一个目标值 target，寻找能够使条件 nums[i] + nums[j] + nums[k] < target 成立的
三元组  i, j, k 个数（0 <= i < j < k < n）。

进阶：
	是否能在 O(n^2) 的时间复杂度内解决？

*/
package _sum_smaller

import "sort"

// --- 自己

/*
方法一: 排序

时间复杂度：
空间复杂度：
*/
func threeSumSmaller(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := 0
	// 可优化为二分查找
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if target > nums[i]+nums[j]+nums[k] {
					ans++
				} else {
					break
				}
			}
		}
	}
	return ans
}

// --- 官方

/*
方法二: 二分查找

时间复杂度：O(n^2logn)。枚举 i 和 j 各需要 O(n) 的时间复杂度，二分查找 k 需要 O(logn) 的时间复杂度，
	因此总的时间复杂度为 O(n^2logn)。
空间复杂度：O(1)。
*/
func threeSumSmallerO2(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	sum := 0
	for i := 0; i < len(nums)-2; i++ {
		sum += twoSumSmaller(nums, i+1, target-nums[i])
	}

	return sum
}

func twoSumSmaller(nums []int, start, target int) int {
	sum := 0
	for i := start; i < len(nums)-1; i++ {
		j := binarySearch(nums, i, target-nums[i])
		sum += j - i
	}
	return sum
}

func binarySearch(nums []int, start, target int) int {
	left := start
	right := len(nums) - 1
	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

/*
方法三: 双指针

时间复杂度：O(n^2)。在每一步中，要么 left 向右移动一位，要么 right 向左移动一位。当 left>=right 时循环结束，因此它的时间复杂度为 O(n)。
	加上外层的循环，总的时间复杂度为 O(n^2)。
空间复杂度：O(1)。
*/
func threeSumSmallerO3(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	sum := 0
	for i := 0; i < len(nums)-2; i++ {
		sum += twoSumSmallerO3(nums, i+1, target-nums[i])
	}

	return sum
}

func twoSumSmallerO3(nums []int, start, target int) int {
	sum := 0
	left, right := start, len(nums)-1
	for left < right {
		if target > nums[left]+nums[right] {
			sum += right - left
			left++
		} else {
			right--
		}
	}

	return sum
}
