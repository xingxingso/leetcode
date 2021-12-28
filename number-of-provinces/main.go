/*
Package number_of_provinces
https://leetcode-cn.com/problems/number-of-provinces/

547. 省份数量

有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，
而 isConnected[i][j] = 0 表示二者不直接相连。

返回矩阵中 省份 的数量。

提示：
	1 <= n <= 200
	n == isConnected.length
	n == isConnected[i].length
	isConnected[i][j] 为 1 或 0
	isConnected[i][i] == 1
	isConnected[i][j] == isConnected[j][i]

*/
package number_of_provinces

// 此题和 number-of-islands 一样

// --- 自己

/*
方法一: union-find 并查集

时间复杂度：
空间复杂度：
*/
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	u := NewUF(n)
	ans := n
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 && !u.Connected(i, j) {
				u.Union(i, j)
				ans--
			}
		}
	}
	return ans
}

type UF struct {
	parent []int
	size   []int
}

func NewUF(n int) *UF {
	uf := &UF{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (u *UF) Find(p int) int {
	// 状态压缩
	if u.parent[p] != p {
		u.parent[p] = u.Find(u.parent[p])
	}
	return u.parent[p]
	//非压缩
	//for u.parent[p] != p {
	//	p = u.parent[p]
	//}
	//return p
}

func (u *UF) Union(p, q int) {
	i, j := u.Find(p), u.Find(q)
	if i == j {
		return
	}
	if u.size[i] < u.size[j] {
		u.parent[i] = j
		u.size[j] += u.size[i]
		return
	}
	u.parent[j] = i
	u.size[i] += u.size[j]
}

func (u *UF) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

// --- 他人

/*
方法一: 深度优先搜索
	LeetCode 101

时间复杂度：
空间复杂度：
*/
func findCircleNumP1(isConnected [][]int) int {
	n := len(isConnected)
	count := 0
	visited := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = true
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 && !visited[j] {
				dfs(j)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}
	return count
}
