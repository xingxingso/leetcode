/*
https://leetcode-cn.com/problems/path-with-minimum-effort/

1631. 最小体力消耗路径

你准备参加一场远足活动。给你一个二维 rows x columns 的地图 heights ，其中 heights[row][col] 表示格子 (row, col) 的高度。
一开始你在最左上角的格子 (0, 0) ，且你希望去最右下角的格子 (rows-1, columns-1) （注意下标从 0 开始编号）。
你每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。
一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。
请你返回从左上角走到右下角的最小 体力消耗值 。

提示：
	rows == heights.length
	columns == heights[i].length
	1 <= rows, columns <= 100
	1 <= heights[i][j] <= 10^6

*/
package path_with_minimum_effort

import (
	"container/heap"
	"math"
	"sort"
)

// --- 官方

/*
方法一: 二分查找 + 广度/深度 优先搜索
	下面是 广度优先搜索

时间复杂度：O(mn*logC)，其中 m 和 n 分别是地图的行数和列数，CC 是格子的最大高度，
	在本题中 C 不超过 10^6 。我们需要进行 O(logC) 次二分查找，每一步查找的过程中需要使用广度优先搜索，
	在 O(mn) 的时间判断是否可以从左上角到达右下角，因此总时间复杂度为 O(mn*logC)。
空间复杂度：O(mn)，即为广度优先搜索中使用的队列需要的空间。
*/

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPathO1(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	return sort.Search(1e6, func(maxHeightDiff int) bool {
		visit := make([][]bool, n)
		for i := 0; i < n; i++ {
			visit[i] = make([]bool, m)
		}
		visit[0][0] = true
		queue := []pair{{0, 0}}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			if p.x == n-1 && p.y == m-1 {
				return true
			}
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				if x >= 0 && x < n && y >= 0 && y < m && !visit[x][y] && abs(heights[x][y]-heights[p.x][p.y]) <= maxHeightDiff {
					visit[x][y] = true
					queue = append(queue, pair{x, y})
				}
			}
		}
		return false
	})
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

/*
方法二: 并查集

时间复杂度：O(mnlog(mn))，其中 m 和 n 分别是地图的行数和列数。图中的边数为 O(mn)，
	因此排序的时间复杂度为 O(mnlog(mn))。并查集的时间复杂度为 O(mn⋅α(mn))，其中 α 为阿克曼函数的反函数。
	由于后者在渐进意义下小于前者，因此总时间复杂度为 O(mnlog(mn))。
空间复杂度：O(mn)，即为存储所有边以及并查集需要的空间。
*/
type unionFind struct {
	parent, size []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent: parent, size: size}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
	//for uf.parent[x] != x {
	//	// 进行路径压缩
	//	uf.parent[x] = uf.parent[uf.parent[x]]
	//	x = uf.parent[x]
	//}
	//return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
}

func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

type edge struct {
	v, w, diff int
}

func minimumEffortPathO2(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	edges := []edge{}
	for i, row := range heights {
		for j, h := range row {
			id := i*m + j
			if i > 0 {
				edges = append(edges, edge{id - m, id, abs(h - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(h - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].diff < edges[j].diff })

	uf := newUnionFind(n * m)
	for _, e := range edges {
		uf.union(e.v, e.w)
		if uf.inSameSet(0, n*m-1) {
			return e.diff
		}
	}
	return 0
}

/*
方法三: 最短路

时间复杂度：O(mnlog(mn))，其中 m 和 n 分别是地图的行数和列数。对于节点数为 n_0，边数为 m_0 的图，
	使用优先队列优化的 Dijkstra 算法的时间复杂度为 O(m_0 \log m_0)。由于图中的边数为 O(mn)，
	带入即可得到时间复杂度 O(mnlog(mn))。
空间复杂度：O(mn)，即为 Dijkstra 算法需要使用的空间。
*/
type point struct{ x, y, maxDiff int }
type hp []point

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].maxDiff < h[j].maxDiff }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(point)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

//type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPathO3(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	maxDiff := make([][]int, n)
	for i := range maxDiff {
		maxDiff[i] = make([]int, m)
		for j := range maxDiff[i] {
			maxDiff[i][j] = math.MaxInt64
		}
	}
	maxDiff[0][0] = 0
	h := &hp{{}}
	for {
		p := heap.Pop(h).(point)
		if p.x == n-1 && p.y == m-1 {
			return p.maxDiff
		}
		if maxDiff[p.x][p.y] < p.maxDiff {
			continue
		}
		for _, d := range dir4 {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m {
				if diff := max(p.maxDiff, abs(heights[x][y]-heights[p.x][p.y])); diff < maxDiff[x][y] {
					maxDiff[x][y] = diff
					heap.Push(h, point{x, y, diff})
				}
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
