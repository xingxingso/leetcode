/*
https://leetcode-cn.com/problems/surrounded-regions/

130. 被围绕的区域

给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

提示：
	m == board.length
	n == board[i].length
	1 <= m, n <= 200
	board[i][j] 为 'X' 或 'O'

*/
package surrounded_regions

// --- 自己

/*
方法一: 并查集
	参考 labuladong

时间复杂度：
空间复杂度：
*/
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	dummy := m * n
	uf := NewUnionFind(m*n + 1)
	// 边界上的 "O" 连通到 dummy
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			uf.union(i*n, dummy)
		}
		if board[i][n-1] == 'O' {
			uf.union(i*n+n-1, dummy)
		}
	}
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			uf.union(i, dummy)
		}
		if board[m-1][i] == 'O' {
			uf.union((m-1)*n+i, dummy)
		}
	}
	// 中心的 "O" 四个方向连通
	type dir struct{ x, y int }
	dirs := []dir{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				for _, dir := range dirs {
					if board[i+dir.x][j+dir.y] == 'O' {
						uf.union(i*n+j, (i+dir.x)*n+j+dir.y)
					}
				}
			}
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !uf.connected(i*n+j, dummy) {
				board[i][j] = 'X'
			}
		}
	}
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

// --- 官方

/*
方法一: 深度优先搜索

时间复杂度：O(n×m)，其中 n 和 m 分别为矩阵的行数和列数。深度优先搜索过程中，每一个点至多只会被标记一次。
空间复杂度：O(n×m)，其中 n 和 m 分别为矩阵的行数和列数。主要为深度优先搜索的栈的开销。
*/
func solveO1(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	n, m := len(board), len(board[0])
	var dfs func(board [][]byte, x, y int)
	dfs = func(board [][]byte, x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'
		dfs(board, x+1, y)
		dfs(board, x-1, y)
		dfs(board, x, y+1)
		dfs(board, x, y-1)
	}

	for i := 0; i < n; i++ {
		dfs(board, i, 0)
		dfs(board, i, m-1)
	}
	for i := 1; i < m-1; i++ {
		dfs(board, 0, i)
		dfs(board, n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

/*
方法二: 深度优先搜索

时间复杂度：O(n×m)，其中 n 和 m 分别为矩阵的行数和列数。深度优先搜索过程中，每一个点至多只会被标记一次。
空间复杂度：O(n×m)，其中 n 和 m 分别为矩阵的行数和列数。主要为深度优先搜索的栈的开销。
*/
func solveO2(board [][]byte) {
	var (
		dx = [4]int{1, -1, 0, 0}
		dy = [4]int{0, 0, 1, -1}
	)

	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m := len(board), len(board[0])
	queue := [][]int{}
	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			queue = append(queue, []int{i, 0})
		}
		if board[i][m-1] == 'O' {
			queue = append(queue, []int{i, m - 1})
		}
	}
	for i := 1; i < m-1; i++ {
		if board[0][i] == 'O' {
			queue = append(queue, []int{0, i})
		}
		if board[n-1][i] == 'O' {
			queue = append(queue, []int{n - 1, i})
		}
	}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1]
		board[x][y] = 'A'
		for i := 0; i < 4; i++ {
			mx, my := x+dx[i], y+dy[i]
			if mx < 0 || my < 0 || mx >= n || my >= m || board[mx][my] != 'O' {
				continue
			}
			queue = append(queue, []int{mx, my})
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
