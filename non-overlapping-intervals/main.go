/*
https://leetcode-cn.com/problems/non-overlapping-intervals/

435. 无重叠区间

给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意:
	1. 可以认为区间的终点总是大于它的起点。
	2. 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

*/
package non_overlapping_intervals

import (
	"sort"
)

// --- 自己

/*
方法一: 贪心

时间复杂度：
空间复杂度：
*/
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	first := intervals[0][1]
	count := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < first {
			count++
		} else {
			first = intervals[i][1]
		}
	}
	return count
}
