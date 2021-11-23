/*
Package shortest_bridge
https://leetcode-cn.com/problems/shortest-bridge/

934. 最短的桥

在给定的二维二进制数组 A 中，存在两座岛。（岛是由四面相连的 1 形成的一个最大组。）

现在，我们可以将 0 变为 1，以使两座岛连接起来，变成一座岛。

返回必须翻转的 0 的最小数目。（可以保证答案至少是 1 。）

提示：
	2 <= A.length == A[0].length <= 100
	A[i][j] == 0 或 A[i][j] == 1

*/
package shortest_bridge

// --- 他人

/*
方法一:
	leetcode 101
	本题实际上是求两个岛屿间的最短距离，因此我们可以先通过任意搜索方法找到其中一个岛屿，
	然后利用广度优先搜索，查找其与另一个岛屿的最短距离。

时间复杂度：
空间复杂度：
*/
func shortestBridge(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	type point struct {
		x, y int
	}
	points := make([]point, 0)

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i == m || j < 0 || j == n || grid[i][j] == 2 {
			return
		}
		if grid[i][j] == 0 {
			points = append(points, point{i, j})
			return
		}
		grid[i][j] = 2
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	// dfs寻找第一个岛屿，并把1全部赋值为2
	flipped := false
	for i := 0; i < m; i++ {
		if flipped {
			break
		}
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				flipped = true
				break
			}
		}
	}

	// bfs寻找第二个岛屿，并把过程中经过的0赋值为2
	count := 0
	direction := []int{-1, 0, 1, 0, -1}
	for len(points) > 0 {
		count++
		for size := len(points); size > 0; size-- {
			p := points[0]
			points = points[1:]
			for i := 0; i < len(direction)-1; i++ {
				x, y := p.x+direction[i], p.y+direction[i+1]
				if x >= 0 && x < m && y >= 0 && y < n {
					if grid[x][y] == 1 {
						return count
					}
					if grid[x][y] == 0 {
						points = append(points, point{x, y})
						grid[x][y] = 2
					}
				}
			}
		}
	}

	return count
}
