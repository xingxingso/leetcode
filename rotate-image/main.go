/*
Package rotate_image
https://leetcode-cn.com/problems/russian-doll-envelopes/

48. 旋转图像

给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

提示：
	matrix.length == n
	matrix[i].length == n
	1 <= n <= 20
	-1000 <= matrix[i][j] <= 1000

*/
package rotate_image

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func rotate(matrix [][]int) {
	n := len(matrix) - 1
	// 对于 n x n 矩阵， 坐标(x, y)旋转90度 得到 (n-y, x)

	var whirl func(i, j int)
	whirl = func(i, j int) {
		if i >= j {
			return
		}

		for k := i; k < j; k++ {
			x, y := i, k // 起点
			temp := matrix[x][y]
			matrix[x][y] = matrix[n-y][x]
			matrix[n-y][x] = matrix[n-x][n-y]
			matrix[n-x][n-y] = matrix[y][n-x]
			matrix[y][n-x] = temp
		}
		//fmt.Println(i, j)
		whirl(i+1, j-1)
	}

	whirl(0, n)

	//fmt.Println(matrix)

	return
}
