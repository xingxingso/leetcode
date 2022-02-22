/*
Package reachable_nodes_in_subdivided_graph
https://leetcode-cn.com/problems/reachable-nodes-in-subdivided-graph/

882. 细分图中的可到达结点

给你一个无向图（原始图），图中有 n 个节点，编号从 0 到 n - 1 。你决定将图中的每条边 细分 为一条节点链，每条边之间的新节点数各不相同。

图用由边组成的二维数组 edges 表示，其中 edges[i] = [ui, vi, cnt_i] 表示原始图中节点 ui 和 vi 之间存在一条边，
cnt_i 是将边 细分 后的新节点总数。注意，cnt_i == 0 表示边不可细分。

要 细分 边 [ui, vi] ，需要将其替换为 (cnt_i + 1) 条新边，和 cnt_i 个新节点。
新节点为 x1, x2, ..., x_{cnt_i} ，新边为 [ui, x1], [x1, x2], [x2, x3], ..., [x_{cnt_i+1}, x_{cnt_i}], [x_{cnt_i}, vi] 。

现在得到一个 新的细分图 ，请你计算从节点 0 出发，可以到达多少个节点？如果节点间距离是 maxMoves 或更少，则视为 可以到达 。

给你原始图和 maxMoves ，返回 新的细分图中从节点 0 出发 可到达的节点数 。

提示：
	0 <= edges.length <= min(n * (n - 1) / 2, 10^4)
	edges[i].length == 3
	0 <= ui < vi < n
	图中 不存在平行边
	0 <= cnt_i <= 10^4
	0 <= maxMoves <= 10^9
	1 <= n <= 3000

*/
package reachable_nodes_in_subdivided_graph

import (
	"container/heap"
)

// --- 官方

/*
方法一: Dijkstra 算法

时间复杂度：
空间复杂度：
*/
func reachableNodesO1(edges [][]int, maxMoves int, n int) int {
	graph := make(map[int]map[int]int)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if graph[u] == nil {
			graph[u] = make(map[int]int)
		}
		if graph[v] == nil {
			graph[v] = make(map[int]int)
		}
		graph[u][v] = w
		graph[v][u] = w
	}

	//fmt.Println(graph)

	// 优先队列
	pq := &simpleHeap{}
	heap.Push(pq, aNode{0, 0})
	dist := make(map[int]int)
	dist[0] = 0
	used := make(map[int]int)
	ans := 0

	for pq.Len() > 0 {
		anode := heap.Pop(pq).(aNode) // 每次选择最小距离的下一个目标
		//fmt.Println("anode", anode)
		node := anode.node
		d := anode.dist

		if d > dist[node] { // 不是最小步数达到该点的方案， 直接废弃这种方案
			//fmt.Println(d, dist, node)
			continue
		}
		// Each node is only visited once.  We've reached
		// a node in our original graph.
		ans++
		if graph[node] == nil {
			continue
		}
		for nei, weight := range graph[node] {
			// M - d is how much further we can walk from this node;
			// weight is how many new nodes there are on this edge.
			// v is the maximum utilization of this edge.
			v := min(weight, maxMoves-d)
			used[n*node+nei] = v
			//fmt.Println("used", used, nei)
			//fmt.Println("dist", dist)

			// d2 is the total distance to reach 'nei' (nei***or) node
			// in the original graph.
			d2 := d + weight + 1
			d3, ok := dist[nei]
			if !ok {
				d3 = maxMoves + 1
			}
			if d2 < d3 { // 说明可达
				heap.Push(pq, aNode{node: nei, dist: d2})
				dist[nei] = d2 // 到达 nei 需要的消耗的步数
			}
		}
	}

	// At the end, each edge (u, v, w) can be used with a maximum
	// of w new nodes: a max of used[u, v] nodes from one side,
	// and used[v, u] nodes from the other.
	// [We use the encoding (u, v) = u * N + v.]
	for _, edge := range edges {
		ans += min(edge[2], used[edge[0]*n+edge[1]]+used[edge[1]*n+edge[0]])
	}
	//fmt.Println("allused", used)

	return ans
}

type aNode struct {
	node, dist int
}

// An simpleHeap is a min-heap of aNode.
type simpleHeap []aNode

func (h simpleHeap) Len() int           { return len(h) }
func (h simpleHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h simpleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *simpleHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(aNode))
}

func (h *simpleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
