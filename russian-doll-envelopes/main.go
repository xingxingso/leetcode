/*
https://leetcode-cn.com/problems/russian-doll-envelopes/

354. 俄罗斯套娃信封问题

给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

注意：不允许旋转信封。

提示：
	1 <= envelopes.length <= 5000
	envelopes[i].length == 2
	1 <= wi, hi <= 104

*/
package russian_doll_envelopes

import (
	"sort"
)

// --- 自己

/*
方法一: 排序+贪心+二分
	// 参考 别人解法

时间复杂度：nlog(n)
空间复杂度：
*/
func maxEnvelopes(envelopes [][]int) int {
	// 宽度正序 高度逆序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	nums := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		nums[i] = envelopes[i][1]
	}

	return lengthOfLIS(nums)
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	d := make([]int, len(nums))
	lens := 0
	for i := 0; i < len(nums); i++ {
		left, right := 0, lens-1
		for left <= right {
			mid := (left + right) / 2
			if d[mid] >= nums[i] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		d[left] = nums[i]
		if left >= lens {
			lens++
		}
	}
	return lens
}
