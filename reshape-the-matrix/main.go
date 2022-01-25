/*
Package reshape_the_matrix
https://leetcode-cn.com/problems/reshape-the-matrix/

566. 重塑矩阵

颠倒给定的 32 位无符号整数的二进制位。

在 MATLAB 中，有一个非常有用的函数 reshape ，它可以将一个 m x n 矩阵重塑为另一个大小不同（r x c）的新矩阵，但保留其原始数据。

给你一个由二维数组 mat 表示的 m x n 矩阵，以及两个正整数 r 和 c ，分别表示想要的重构的矩阵的行数和列数。

重构后的矩阵需要将原始矩阵的所有元素以相同的 行遍历顺序 填充。

如果具有给定参数的 reshape 操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。

提示：
	m == mat.length
	n == mat[i].length
	1 <= m, n <= 100
	-1000 <= mat[i][j] <= 1000
	1 <= r, c <= 300

*/
package reshape_the_matrix

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}
	matrix := make([][]int, r)
	for i := 0; i < r; i++ {
		matrix[i] = make([]int, c)
	}

	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		p := i*n + j
	//		row := p / c
	//		col := p % c
	//		matrix[row][col] = mat[i][j]
	//	}
	//}

	all := n * m
	for i := 0; i < all; i++ {
		matrix[i/c][i%c] = mat[i/m][i%m]
	}

	return matrix
}
