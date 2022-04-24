/*
Package rotting_oranges
https://leetcode-cn.com/problems/rotting-oranges/

994. 腐烂的橘子

在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
	值 0 代表空单元格；
	值 1 代表新鲜橘子；
	值 2 代表腐烂的橘子。

每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。

返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。

提示：
	m == grid.length
	n == grid[i].length
	1 <= m, n <= 10
	grid[i][j] 仅为 0、1 或 2

*/
package rotting_oranges

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func orangesRotting(grid [][]int) int {
	type point struct {
		x, y int
	}
	oranges := make([]point, 0)
	good := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				oranges = append(oranges, point{i, j})
			} else if grid[i][j] == 1 {
				good++
			}
		}
	}
	count, m, n := 0, len(grid), len(grid[0])
	dirs := []int{1, 0, -1, 0, 1}
	for len(oranges) > 0 {
		size := len(oranges)
		for i := 0; i < size; i++ {
			for j := 0; j < 4; j++ {
				x, y := oranges[i].x+dirs[j], oranges[i].y+dirs[j+1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					oranges = append(oranges, point{x, y})
					good--
				}
			}
		}
		oranges = oranges[size:]
		if len(oranges) > 0 {
			count++
		}
	}
	if good > 0 {
		return -1
	}

	return count
}
