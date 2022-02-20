/*
Package is_graph_bipartite
https://leetcode-cn.com/problems/is-graph-bipartite/

785. 判断二分图

存在一个 无向图 ，图中有 n 个节点。其中每个节点都有一个介于 0 到 n - 1 之间的唯一编号。
给你一个二维数组 graph ，其中 graph[u] 是一个节点数组，由节点 u 的邻接节点组成。
形式上，对于 graph[u] 中的每个 v ，都存在一条位于节点 u 和节点 v 之间的无向边。
该无向图同时具有以下属性：
	不存在自环（graph[u] 不包含 u）。
	不存在平行边（graph[u] 不包含重复值）。
	如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
	这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。

二分图 定义：
	如果能将一个图的节点集合分割成两个独立的子集 A 和 B ，
	并使图中的每一条边的两个节点一个来自 A 集合，一个来自 B 集合，就将这个图称为 二分图 。

如果图是二分图，返回 true ；否则，返回 false 。

提示：
	graph.length == n
	1 <= n <= 100
	0 <= graph[u].length < n
	0 <= graph[u][i] <= n - 1
	graph[u] 不会包含 u
	graph[u] 的所有值 互不相同
	如果 graph[u] 包含 v，那么 graph[v] 也会包含 u

*/
package is_graph_bipartite

// --- 他人

/*
方法一: 二分图算法(染色法, 一种广度优先算法)
	leetcode101 15.2
	利用队列和广度优先搜索，我们可以对未染色的节点进行染色，并且检查是否有颜色相同的 相邻节点存在。
	注意在代码中，我们用 0 表示未检查的节点，用 1 和 2 表示两种不同的颜色。

时间复杂度：
空间复杂度：
*/
func isBipartite(graph [][]int) bool {
	n := len(graph)
	if n == 0 {
		return false
	}
	color := make([]int, n)
	queue := make([]int, 0)

	for i := 0; i < n; i++ {
		if color[i] == 0 {
			queue = append(queue, i)
			color[i] = 1
		}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			for _, j := range graph[node] {
				//fmt.Println(color, j, queue)
				if color[j] == 0 {
					queue = append(queue, j)
					color[j] = 2
					if color[node] == 2 {
						color[j] = 1
					}
				} else if color[node] == color[j] {
					return false
				}
			}
		}
	}

	return true
}
