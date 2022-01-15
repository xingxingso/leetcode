/*
Package find_all_numbers_disappeared_in_an_array
https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/

448. 找到所有数组中消失的数字

给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

提示：
	n == nums.length
	1 <= n <= 10^5
	1 <= nums[i] <= n

进阶：
	你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。

*/
package find_all_numbers_disappeared_in_an_array

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[i] == i+1 {
			i++
			continue
		}

		p := nums[i]
		if nums[p-1] == p {
			i++
			continue
		}

		nums[p-1], nums[i] = nums[i], nums[p-1]
	}

	//fmt.Println(nums)
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ans = append(ans, i+1)
		}
	}

	return ans
}
