/*
Package redundant_connection
https://leetcode-cn.com/problems/redundant-connection/

684. 冗余连接

树可以看成是一个连通且 无环 的 无向 图。

给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。
图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。

请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的边。

提示:
	n == edges.length
	3 <= n <= 1000
	edges[i].length == 2
	1 <= ai < bi <= edges.length
	ai != bi
	edges 中无重复元素
	给定的图是连通的

*/
package redundant_connection

// --- 自己

/*
方法一: 并查集

时间复杂度：
空间复杂度：
*/
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := NewUnionFind(n)
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		if uf.connected(x, y) {
			return edge
		}
		uf.union(x, y)
	}

	return nil
}

type unionFind struct {
	parent []int // 存储一棵树
	size   []int // 记录树的“重量”
}

func NewUnionFind(n int) *unionFind {
	parent, size := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{
		parent: parent,
		size:   size,
	}
}

// 返回某个节点 x 的根节点
func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
	// 下面是等价的
	//for uf.parent[x] != x {
	//	// 进行路径压缩
	//	uf.parent[x] = uf.parent[uf.parent[x]]
	//	x = uf.parent[x]
	//}
	//return uf.parent[x]
}

// 将 x 和 y 连接
func (uf *unionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	// 小树接到大树下面，较平衡
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
}

// x 和 y 是否连通
func (uf *unionFind) connected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}
