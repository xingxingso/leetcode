/*
Package missing_ranges
https://leetcode-cn.com/problems/missing-ranges/

163. 缺失的区间

给定一个排序的整数数组 nums ，其中元素的范围在 闭区间 [lower, upper] 当中，返回不包含在数组中的缺失区间。

*/
package missing_ranges

import "strconv"

// --- 自己

/*
方法一:
时间复杂度：O(N)
空间复杂度：O(N)
*/
func findMissingRanges(nums []int, lower int, upper int) []string {
	res := make([]string, 0)

	if len(nums) == 0 {
		if lower == upper {
			res = append(res, strconv.Itoa(lower))
		} else {
			res = append(res, strconv.Itoa(lower)+"->"+strconv.Itoa(upper))
		}
		return res
	}

	for i := 0; i < len(nums); i++ {
		// 第一个特殊处理
		if i == 0 {
			if lower+1 == nums[i] {
				res = append(res, strconv.Itoa(lower))
			} else if lower+1 < nums[i] {
				res = append(res, strconv.Itoa(lower)+"->"+strconv.Itoa(nums[i]-1))
			}
			// continue // len==1
		}

		// 最后一个特殊处理
		if i == len(nums)-1 {
			if nums[i]+1 < upper {
				res = append(res, strconv.Itoa(nums[i]+1)+"->"+strconv.Itoa(upper))
			} else if nums[i]+1 == upper {
				res = append(res, strconv.Itoa(upper))
			}
			continue
		}

		if nums[i]+2 == nums[i+1] {
			res = append(res, strconv.Itoa(nums[i]+1))
		} else if nums[i]+2 < nums[i+1] {
			res = append(res, strconv.Itoa(nums[i]+1)+"->"+strconv.Itoa(nums[i+1]-1))
		}
	}

	return res
}
