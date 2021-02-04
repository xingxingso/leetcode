/*
https://leetcode-cn.com/problems/two-sum-less-than-k/

1099. 小于 K 的两数之和

给你一个整数数组 nums 和整数 k ，返回最大和 sum ，满足存在 i < j 使得 nums[i] + nums[j] = sum 且 sum < k 。
如果没有满足此等式的 i,j 存在，则返回 -1 。

提示：
	1 <= nums.length <= 100
	1 <= nums[i] <= 1000
	1 <= k <= 2000

*/
package two_sum_less_than_k

import (
	"sort"
)

// --- 自己

/*
方法一: 排序后双指针
	双指针没用好

时间复杂度：O(n^2)
空间复杂度：O(1)
*/
func twoSumLessThanK(nums []int, k int) int {
	if len(nums) < 2 {
		return -1
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	// fmt.Println(nums)
	// todo: 优化，先二分查找到小于k的最大下标，right 赋值为该值 减少查找次数

	left := len(nums) - 2
	ans := -1
	for left >= 0 {
		right := len(nums) - 1
		for left < right {
			if k > nums[left]+nums[right] {
				ans = max(ans, nums[left]+nums[right])
			}
			right--
		}
		left--
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// --- 他人

/*
https://leetcode-cn.com/problems/two-sum-less-than-k/solution/tu-jie-xiao-yu-k-de-liang-shu-zhi-he-by-misterbooo/
方法一: 排序后双指针

时间复杂度：O(nlogn)
空间复杂度：O(1)
*/
func twoSumLessThanKP1(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	l, r := 0, len(nums)-1

	ans := -1

	for l < r {
		if nums[l]+nums[r] < k {
			ans = max(ans, nums[l]+nums[r])
			l++
		} else {
			r--
		}
	}

	return ans
}
