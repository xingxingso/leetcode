/*
Package number_of_islands
https://leetcode-cn.com/problems/number-of-islands/

200. 岛屿数量

给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

提示：
	m == grid.length
	n == grid[i].length
	1 <= m, n <= 300
	grid[i][j] 的值为 '0' 或 '1'

*/
package number_of_islands

// 此题和 number-of-provinces 一样

// --- 自己

/*
方法一: 广度优先遍历

时间复杂度：
空间复杂度：
*/
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	type point struct {
		x, y int
	}
	queue := make([]point, 0)
	ans := 0
	directs := []int{1, 0, -1, 0, 1} // 四个方向
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				queue = append(queue, point{i, j})
				grid[i][j] = '0'
				ans++
				for len(queue) > 0 {
					size := len(queue)
					for k := 0; k < size; k++ {
						x, y := queue[k].x, queue[k].y
						for p := 0; p < len(directs)-1; p++ {
							newX, newY := x+directs[p], y+directs[p+1]
							if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == '1' {
								grid[newX][newY] = '0'
								queue = append(queue, point{newX, newY})
							}
						}
					}
					queue = queue[size:]
				}
			}
		}
	}

	return ans
}
