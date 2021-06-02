/*
Package two_sum
https://leetcode-cn.com/problems/two-sum/

1. 两数之和

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

说明:
	1. 必须在原数组上操作，不能拷贝额外的数组。
	2. 尽量减少操作次数。

*/
package two_sum

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	l := len(nums)

	for i := 0; i < l; i++ {
		if j, ok := numsMap[target-nums[i]]; ok && i != j {
			return []int{j, i}
		}
		numsMap[nums[i]] = i
	}
	return []int{}
}
