/*
Package _4sum
https://leetcode-cn.com/problems/4sum/

18. 四数之和

给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：答案中不可以包含重复的四元组。

提示：
	0 <= nums.length <= 200
	-109 <= nums[i] <= 109
	-109 <= target <= 109

*/
package _4sum

import (
	"sort"
)

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for first := 0; first < len(nums); first++ {
		if first > 0 && first < len(nums) && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && second < len(nums) && nums[second] == nums[second-1] {
				continue
			}
			fourth := len(nums) - 1
			remain := target - nums[first] - nums[second]
			for third := second + 1; third < fourth; third++ {
				if third > second+1 && third < fourth && nums[third] == nums[third-1] {
					continue
				}
				for third < fourth && nums[third]+nums[fourth] > remain {
					fourth--
				}
				if third == fourth {
					break
				}
				if nums[third]+nums[fourth] == remain {
					res = append(res, []int{nums[first], nums[second], nums[third], nums[fourth]})
				}
			}
		}
	}
	return res
}
