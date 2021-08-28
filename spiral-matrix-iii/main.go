/*
Package spiral_matrix_iii
https://leetcode-cn.com/problems/spiral-matrix-iii/

885. 螺旋矩阵 III

在 R 行 C 列的矩阵上，我们从 (r0, c0) 面朝东面开始
这里，网格的西北角位于第一行第一列，网格的东南角位于最后一行最后一列。
现在，我们以顺时针按螺旋状行走，访问此网格中的每个位置。
每当我们移动到网格的边界之外时，我们会继续在网格之外行走（但稍后可能会返回到网格边界）。
最终，我们到过网格的所有 R * C 个空间。
按照访问顺序返回表示网格位置的坐标列表。

提示：
	1 <= R <= 100
	1 <= C <= 100
	0 <= r0 < R
	0 <= c0 < C

*/
package spiral_matrix_iii

// --- 自己

/*
方法一：

时间复杂度：
空间复杂度：
*/
func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	directors := [][]int{
		{0, 1},  // 东
		{1, 0},  // 南
		{0, -1}, // 西
		{-1, 0}, // 北
	}

	ans := make([][]int, 0)
	ans = append(ans, []int{rStart, cStart})

	direct := 0
	count := 1
	line := 1
	lineCount := 0
	total := rows * cols
	changeTimes := 0
	for count < total {
		rStart += directors[direct][0]
		cStart += directors[direct][1]
		if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
			ans = append(ans, []int{rStart, cStart})
			count++
		}
		lineCount++
		//1,1,2,2,3,3,4,4,5,5 每转向两次增加1个
		if lineCount >= line {
			lineCount = 0
			changeTimes++
			if changeTimes == 2 {
				changeTimes = 0
				line++
			}
			direct++
			if direct >= len(directors) {
				direct = 0
			}
		}
	}

	return ans
}
