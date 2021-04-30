/*
https://leetcode-cn.com/problems/path-with-minimum-effort/

1631. 最小体力消耗路径

你准备参加一场远足活动。给你一个二维 rows x columns 的地图 heights ，其中 heights[row][col] 表示格子 (row, col) 的高度。
一开始你在最左上角的格子 (0, 0) ，且你希望去最右下角的格子 (rows-1, columns-1) （注意下标从 0 开始编号）。
你每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。
一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。
请你返回从左上角走到右下角的最小 体力消耗值 。

提示：
	rows == heights.length
	columns == heights[i].length
	1 <= rows, columns <= 100
	1 <= heights[i][j] <= 106

*/
package path_with_minimum_effort

import "fmt"

// --- 自己

/*
方法一: // todo: 未完成

时间复杂度：
空间复杂度：
*/
func minimumEffortPath(heights [][]int) int {
	rows := len(heights)
	columns := len(heights[0])
	// maxL, maxR := 0, 0
	fmt.Println(rows, columns)

	ans := 0

	var dfs func(l, r int, heights [][]int, m int) int
	dfs = func(l, r int, heights [][]int, m int) int {
		el, er := m, m
		if l >= rows-1 && r >= columns-1 {
			return m
		}

		if l < rows-1 {
			ce := abs(heights[l+1][r], heights[l][r])
			el = dfs(l+1, r, heights, max(ce, m))
		}

		if r < columns-1 {
			ce := abs(heights[l][r+1], heights[l][r])
			er = dfs(l, r+1, heights, max(ce, m))
		}

		fmt.Println(l, r, el, er)
		return min(el, er)
	}

	// el := dfs(1, 0, heights, maxL)
	// er := dfs(0, 1, heights, maxR)
	// ans := min(el, er)

	ans = dfs(0, 0, heights, 0)

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}
