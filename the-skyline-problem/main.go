/*
Package the_skyline_problem
https://leetcode-cn.com/problems/the-skyline-problem/

218. 天际线问题

城市的天际线是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。给你所有建筑物的位置和高度，请返回由这些建筑物形成的 天际线 。

每个建筑物的几何信息由数组 buildings 表示，其中三元组 buildings[i] = [left_i, right_i, height_i] 表示：
	- left_i 是第 i 座建筑物左边缘的 x 坐标。
	- right_i 是第 i 座建筑物右边缘的 x 坐标。
	- height_i 是第 i 座建筑物的高度。

天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，并按 x 坐标 进行 排序 。
关键点是水平线段的左端点。列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0 ，仅用于标记天际线的终点。
此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。

注意：
	输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答案；
	三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]

提示：
	1 <= buildings.length <= 10^4
	0 <= left_i < right_i <= 2^31 - 1
	1 <= height_i <= 2^31 - 1
	buildings 按 left_i 非递减排序

*/
package the_skyline_problem

import (
	"container/heap"
	"fmt"
	"sort"
)

// --- 官方

/*
方法一: 堆

时间复杂度：
空间复杂度：
*/
func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	boundaries := make([]int, 0, 2*n)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1]) // 左右节点
	}
	sort.Ints(boundaries)
	fmt.Println(boundaries)

	idx := 0
	h := &hp{}
	ans := make([][]int, 0)

	for _, boundary := range boundaries {
		for idx < n && buildings[idx][0] <= boundary {
			heap.Push(h, pair{buildings[idx][1], buildings[idx][2]})
			idx++
		}
		fmt.Println(*h)
		for h.Len() > 0 && h.Peek().(pair).right <= boundary {
			heap.Pop(h)
		}

		maxn := 0
		if h.Len() > 0 {
			maxn = h.Peek().(pair).height
		}
		if len(ans) == 0 || maxn != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxn})
		}
	}
	return ans
}

// An hp is a min-heap.
type pair struct{ right, height int }
type hp []pair // [高度，右端]

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Peek 获取堆顶元素 不弹出
func (h *hp) Peek() interface{} { return (*h)[0] }
