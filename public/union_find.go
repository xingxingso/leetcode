package public

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
