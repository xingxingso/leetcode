/*
https://leetcode-cn.com/problems/remove-covered-intervals/

1288. 删除被覆盖区间

给你一个区间列表，请你删除列表中被其他区间所覆盖的区间。
只有当 c <= a 且 b <= d 时，我们才认为区间 [a,b) 被区间 [c,d) 覆盖。
在完成所有删除操作后，请你返回列表中剩余区间的数目。

提示：
	1 <= intervals.length <= 1000
	0 <= intervals[i][0] < intervals[i][1] <= 10^5
	对于所有的 i != j：intervals[i] != intervals[j]

*/
package remove_covered_intervals

import (
	"sort"
)

// --- 自己

/*
方法一: 贪心算法

时间复杂度：O(NlogN)，排序过程占主导地位。
空间复杂度：O(1)。
*/
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] { // 优先按起点排序
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] > intervals[j][1] // 否则按终点降序
		}
	})
	n := 0
	for i := 0; i < len(intervals); {
		j := i + 1
		for ; j < len(intervals) && intervals[i][1] >= intervals[j][1]; j++ {
			n++
		}
		i = j
	}

	return len(intervals) - n
}
