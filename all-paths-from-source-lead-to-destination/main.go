/*
https://leetcode-cn.com/problems/all-paths-from-source-lead-to-destination/

1059. 从始点到终点的所有路径

给定有向图的边 edges，以及该图的始点 source 和目标终点 destination，确定从始点 source 出发的所有路径是否最终结束于目标终点 destination，即：
	从始点 source 到目标终点 destination 存在至少一条路径
	如果存在从始点 source 到没有出边的节点的路径，则该节点就是路径终点。
	从始点source到目标终点 destination 可能路径数是有限数字

当从始点 source 出发的所有路径都可以到达目标终点 destination 时返回 true，否则返回 false。

提示：
	给定的图中可能带有自环和平行边。
	图中的节点数 n 介于 1 和 10000 之间。
	图中的边数在 0 到 10000 之间。
	0 <= edges.length <= 10000
	edges[i].length == 2
	0 <= source <= n - 1
	0 <= destination <= n - 1

*/
package all_paths_from_source_lead_to_destination

// --- 自己

/*
方法一: 递归
	每次从移除目标点后，更新目标点
	这个 n 没用到

时间复杂度：
空间复杂度：
*/
func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	newEdges := make([][]int, 0, len(edges))
	newSources := make([]int, 0)

	hasSource := false // 标志是否有下一路径
	for _, v := range edges {
		if v[0] == source {
			newSources = append(newSources, v[1])
			hasSource = true
		} else {
			newEdges = append(newEdges, v)
		}
	}

	// 当起点==终点 且没有下一路径的时候 说明到达终点
	if source == destination {
		return !hasSource
	}

	// 起点 != 终点的时候 没有下一路径 说明不符合
	if !hasSource {
		return false
	}

	for _, v := range newSources {
		if !leadsToDestination(n-1, newEdges, v, destination) {
			return false
		}
	}

	// 未到达最终 则认为在路径中 返回 true
	return true
}
