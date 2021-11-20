/*
Package max_area_of_island
https://leetcode-cn.com/problems/max-area-of-island/

695. 岛屿的最大面积

给你一个大小为 m x n 的二进制矩阵 grid 。

岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。
你可以假设 grid 的四个边缘都被 0（代表水）包围着。

岛屿的面积是岛上值为 1 的单元格的数目。

计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

提示：
	m == grid.length
	n == grid[i].length
	1 <= m, n <= 50
	grid[i][j] 为 0 或 1

*/
package max_area_of_island

// --- 自己

/*
方法一: 深度优先搜索

时间复杂度：
空间复杂度：
*/
func maxAreaOfIsland(grid [][]int) int {
	ans, area := 0, 0
	m, n := len(grid), len(grid[0])
	memo := map[int]bool{}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		memo[100*i+j] = true
		area++
		if j < n-1 && !memo[100*i+j+1] && grid[i][j+1] != 0 {
			dfs(i, j+1)
		}
		if i < m-1 && !memo[100*i+j+100] && grid[i+1][j] != 0 {
			dfs(i+1, j)
		}
		if i > 0 && !memo[100*i+j-100] && grid[i-1][j] != 0 {
			dfs(i-1, j)
		}
		if j > 0 && !memo[100*i+j-1] && grid[i][j-1] != 0 {
			dfs(i, j-1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !memo[100*i+j] && grid[i][j] != 0 {
				area = 0
				dfs(i, j)
				if area > ans {
					ans = area
				}
			}
		}
	}

	return ans
}

// --- 他人

/*
方法一: 深度优先搜索
	// LeetCode 101 6.2 深度优先搜索
	// 做了修改

时间复杂度：
空间复杂度：
*/
func maxAreaOfIslandP1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	area, ans := 0, 0
	type pointer struct {
		x, y int
	}
	direction := []int{-1, 0, 1, 0, -1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area = 1
				grid[i][j] = 0
				stack := []pointer{{i, j}}
				for len(stack) > 0 {
					p := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					for k := 0; k < len(direction)-1; k++ {
						x, y := p.x+direction[k], p.y+direction[k+1]
						if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
							grid[x][y] = 0
							area++
							stack = append(stack, pointer{x, y})
						}
					}
				}
				if area > ans {
					ans = area
				}
			}
		}
	}
	return ans
}
