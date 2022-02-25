/*
Package merge_intervals
https://leetcode-cn.com/problems/merge-intervals/

56. 合并区间

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

提示：
	1 <= intervals.length <= 104
	intervals[i].length == 2
	0 <= starti <= endi <= 104

*/
package merge_intervals

import (
	"sort"
)

// --- 自己

/*
方法一: 排序

时间复杂度：
空间复杂度：
*/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	ans := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		for j := i + 1; j < len(intervals); j++ {
			// 没有交集只有一种情况
			if intervals[j][0] > end {
				break
			} else {
				end = max(end, intervals[j][1])
				i++
			}
		}
		ans = append(ans, []int{start, end})
	}
	return ans
}

/*
方法一: 排序优化

时间复杂度：
空间复杂度：
*/
func mergeS1(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		l, r := intervals[i][0], intervals[i][1]
		if len(ans) == 0 || ans[len(ans)-1][1] < l {
			ans = append(ans, []int{l, r})
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], r)
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
