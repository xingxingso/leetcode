/*
Package _3sum
https://leetcode-cn.com/problems/3sum/

15. 三数之和

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：
	答案中不可以包含重复的三元组。

提示：
	0 <= nums.length <= 3000
	-10^5 <= nums[i] <= 10^5

*/
package _3sum

import "sort"

// --- 自己

/*
方法一: 排序 + 双指针

时间复杂度：O(N^2)，其中 N 是数组 nums 的长度。
空间复杂度：O(logN)。我们忽略存储答案的空间，额外的排序的空间复杂度为 O(logN)。然而我们修改了输入的数组 nums，在实际情况下不一定允许，
	因此也可以看成使用了一个额外的数组存储了 nums 的副本并进行排序，空间复杂度为 O(N)。
*/
func threeSumS1(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	r := len(nums) - 1
	ans := make([][]int, 0)
	var twoSum func(l, r int, target int) [][]int
	twoSum = func(l, r int, target int) [][]int {
		res := make([][]int, 0)
		for l < r {
			left := nums[l]
			right := nums[r]
			sum := left + right
			if sum == target {
				res = append(res, []int{left, right})
				for l < r && nums[l] == left {
					l++
				}
				for l < r && nums[r] == right {
					r--
				}
			} else if sum < target {
				for l < r && nums[l] == left {
					l++
				}
			} else {
				for l < r && nums[r] == right {
					r--
				}
			}
		}
		return res
	}

	for i := 0; i < len(nums); {
		left := nums[i]
		res := twoSum(i+1, r, 0-left)
		for _, v := range res {
			ans = append(ans, []int{nums[i], v[0], v[1]})
		}
		for i++; i < len(nums) && nums[i] == left; i++ {
		}
	}

	return ans
}

// --- 官方

/*
方法一: 排序 + 双指针

时间复杂度：O(N^2)，其中 N 是数组 nums 的长度。
空间复杂度：O(logN)。我们忽略存储答案的空间，额外的排序的空间复杂度为 O(logN)。然而我们修改了输入的数组 nums，在实际情况下不一定允许，
	因此也可以看成使用了一个额外的数组存储了 nums 的副本并进行排序，空间复杂度为 O(N)。
*/
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
