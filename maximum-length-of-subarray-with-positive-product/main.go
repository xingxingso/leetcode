/*
Package maximum_length_of_subarray_with_positive_product
https://leetcode-cn.com/problems/maximum-length-of-subarray-with-positive-product/

1567. 乘积为正数的最长子数组长度

给你一个整数数组 nums ，请你求出乘积为正数的最长子数组的长度。

一个数组的子数组是由原数组中零个或者更多个连续数字组成的数组。

请你返回乘积为正数的最长子数组长度。

提示：
	1 <= nums.length <= 10^5
	-10^9 <= nums[i] <= 10^9

*/
package maximum_length_of_subarray_with_positive_product

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func getMaxLen(nums []int) int {
	var lastPos, lastNeg, ans int
	if nums[0] > 0 {
		lastPos = 1
		ans = 1
	} else if nums[0] < 0 {
		lastNeg = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			lastPos, lastNeg = 0, 0
			continue
		}
		lp, ln := lastPos, lastNeg
		if nums[i] > 0 {
			lastPos++
			if lastNeg > 0 {
				lastNeg++
			}
		} else {
			if ln == 0 {
				lastPos = 0
			} else {
				lastPos = ln + 1
			}
			lastNeg = lp + 1
		}
		// fmt.Println(lastPos, lastNeg)
		if lastPos > ans {
			ans = lastPos
		}
	}
	return ans
}
