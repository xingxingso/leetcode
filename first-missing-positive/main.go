/*
Package first_missing_positive
https://leetcode-cn.com/problems/first-missing-positive/

41. 缺失的第一个正数

给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

提示：
	1 <= nums.length <= 5 * 10^5
	-2^31 <= nums[i] <= 2^31 - 1

*/
package first_missing_positive

// --- 自己

/*
方法一:
	// 长度为n的数组最大可能答案为 n+1, 最小为0
	// 利用数组本身的下标标准已出现的数字

时间复杂度：
空间复杂度：
*/
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); {
		t := nums[i]
		if t == i+1 { // 已经在标志位
			i++
			continue
		}

		if t > 0 && t <= len(nums) && nums[t-1] != t { // 不在标志位, 并且标志位没有被占据(因为允许重复数字）
			nums[t-1], nums[i] = nums[i], nums[t-1]
			continue
		}
		i++ // 非合理数字，大于数组长度或者0或者负数
	}

	//fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 { // 第一个出现的就是最小值
			return i + 1
		}
	}

	return len(nums) + 1 // 如果位置都填满了， 说明是数组长度+1
}
