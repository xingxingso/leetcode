/*
Package interval_list_intersections
https://leetcode-cn.com/problems/interval-list-intersections/

986. 区间列表的交集

给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi]
而 secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。
返回这 两个区间列表的交集 。
形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。
两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。

提示：
	0 <= firstList.length, secondList.length <= 1000
	firstList.length + secondList.length >= 1
	0 <= start_i < end_i <= 10^9
	end_i < start_{i+1}
	0 <= start_j < end_j <= 10^9
	end_j < start_{j+1}

*/
package interval_list_intersections

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := make([][]int, 0)
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		if !(firstList[i][0] > secondList[j][1] || secondList[j][0] > firstList[i][1]) {
			temp := []int{max(firstList[i][0], secondList[j][0]), min(firstList[i][1], secondList[j][1])}
			ans = append(ans, temp)
		}
		if firstList[i][1] > secondList[j][1] {
			j++
		} else {
			i++
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
