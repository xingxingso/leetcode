/*
Package pacific_atlantic_water_flow
https://leetcode-cn.com/problems/pacific-atlantic-water-flow/

417. 太平洋大西洋水流问题

给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。

规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。

请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。

提示：
	输出坐标的顺序不重要
	m 和 n 都小于150

*/
package pacific_atlantic_water_flow

// --- 自己

/*
方法一: dfs

时间复杂度：
空间复杂度：
*/
func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacific, atlantic := false, false
	direction := []int{0, -1, 0, 1, 0}
	memo := make(map[int]bool)
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if pacific && atlantic {
			return
		}
		if memo[1000*i+j] {
			return
		}
		memo[1000*i+j] = true
		for k := 0; k < len(direction)-1; k++ {
			x, y := i+direction[k], j+direction[k+1]
			if x >= 0 && x < m && y >= 0 && y < n && heights[x][y] <= heights[i][j] {
				dfs(x, y)
			} else {
				if x < 0 || y < 0 {
					pacific = true
				}
				if x >= m || y >= n {
					atlantic = true
				}
			}
		}
	}

	ans := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pacific, atlantic = false, false
			memo = map[int]bool{}
			dfs(i, j)
			if pacific && atlantic {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

// --- 他人

/*
方法一:
	LeetCode 101

时间复杂度：
空间复杂度：
*/
func pacificAtlanticP1(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}
	m, n := len(heights), len(heights[0])
	direction := []int{-1, 0, 1, 0, -1}
	canReachP, canReachA := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		canReachP[i], canReachA[i] = make([]bool, n), make([]bool, n)
	}
	ans := make([][]int, 0)
	var dfs func(i, j int, canReach [][]bool)
	dfs = func(i, j int, canReach [][]bool) {
		if canReach[i][j] {
			return
		}
		canReach[i][j] = true
		for k := 0; k < len(direction)-1; k++ {
			x, y := i+direction[k], j+direction[k+1]
			if x >= 0 && x < m && y >= 0 && y < n && heights[i][j] <= heights[x][y] {
				dfs(x, y, canReach)
			}
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0, canReachP)
		dfs(i, n-1, canReachA)
	}
	for i := 0; i < n; i++ {
		dfs(0, i, canReachP)
		dfs(m-1, i, canReachA)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canReachP[i][j] && canReachA[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
