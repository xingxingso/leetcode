/*
Package minimum_height_trees
https://leetcode-cn.com/problems/minimum-height-trees/

310. 最小高度树

树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。

给你一棵包含 n 个节点的树，标记为 0 到 n - 1 。
给定数字 n 和一个有 n - 1 条无向边的 edges 列表（每一个边都是一对标签），
其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。

可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。
在所有可能的树中，具有最小高度的树（即，min(h)）被称为 最小高度树 。

请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。

树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。

提示：
	1 <= n <= 2 * 10^4
	edges.length == n - 1
	0 <= ai, bi < n
	ai != bi
	所有 (ai, bi) 互不相同
	给定的输入 保证 是一棵树，并且 不会有重复的边

*/
package minimum_height_trees

// --- 自己

/*
方法一:
	// 超出时间限制

时间复杂度：
空间复杂度：
*/
func findMinHeightTrees(n int, edges [][]int) []int {
	if len(edges) == 0 {
		return []int{0}
	}

	tree := make(map[int][]int, 0)
	for _, v := range edges {
		tree[v[0]] = append(tree[v[0]], v[1])
		tree[v[1]] = append(tree[v[1]], v[0])
	}

	//fmt.Println(tree)

	var dfs func(node, pre int) int
	dfs = func(node, pre int) int {
		if len(tree[node]) == 1 && tree[node][0] == pre {
			return 1
		}

		maxH := 0
		for _, v := range tree[node] {
			if v == pre {
				continue
			}
			if h := dfs(v, node); h > maxH {
				maxH = h
			}
		}
		return maxH + 1
	}

	ans := make([]int, 0)
	minHeight := n
	for root := range tree {
		height := dfs(root, root)
		if height < minHeight {
			minHeight = height
			ans = []int{root}
		} else if height == minHeight {
			ans = append(ans, root)
		}
	}

	//fmt.Println(ans)
	return ans
}

// --- 他人

/*
方法二:
	https://leetcode-cn.com/problems/minimum-height-trees/solution/zui-rong-yi-li-jie-de-bfsfen-xi-jian-dan-zhu-shi-x/

时间复杂度：
空间复杂度：
*/
func findMinHeightTreesP1(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	//建立各个节点的出度表
	degree := make([]int, n)
	//建立图关系，在每个节点的list中存储相连节点
	tree := make(map[int][]int, 0)
	for _, edge := range edges {
		degree[edge[0]]++                              // 出度++
		degree[edge[1]]++                              // 出度++
		tree[edge[0]] = append(tree[edge[0]], edge[1]) // 添加相邻节点
		tree[edge[1]] = append(tree[edge[1]], edge[0]) // 添加相邻节点
	}

	// 建立队列
	queue := make([]int, 0)
	// 把所有出度为1的节点，也就是叶子节点入队
	for i, d := range degree {
		if d == 1 {
			queue = append(queue, i)
		}
	}

	//fmt.Println(tree, degree)

	var res []int
	for len(queue) > 0 {
		// 这个地方注意，我们每层循环都要new一个新的结果集合，
		// 这样最后保存的就是最终的最小高度树了
		res = []int{}
		for size := len(queue) - 1; size >= 0; size-- {
			cur := queue[0]
			queue = queue[1:]
			// 把当前节点加入结果集，不要有疑问，为什么当前只是叶子节点为什么要加入结果集呢?
			// 因为我们每次循环都会新建一个list，所以最后保存的就是最后一个状态下的叶子节点，
			// 这也是很多题解里面所说的剪掉叶子节点的部分，你可以想象一下图，每层遍历完，
			// 都会把该层（也就是叶子节点层）这一层从队列中移除掉，
			// 不就相当于把原来的图给剪掉一圈叶子节点，形成一个缩小的新的图吗
			res = append(res, cur)
			// 这里就是经典的bfs了，把当前节点的相邻接点都拿出来，
			// 把它们的出度都减1，因为当前节点已经不存在了，所以，
			// 它的相邻节点们就有可能变成叶子节点
			for _, neighbor := range tree[cur] {
				degree[neighbor]--
				if degree[neighbor] == 1 {
					//如果是叶子节点我们就入队
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return res // 返回最后一次保存的list
}
