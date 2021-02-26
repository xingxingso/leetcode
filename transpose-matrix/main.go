/*
https://leetcode-cn.com/problems/transpose-matrix/

867. 转置矩阵

给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。
矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。

提示：
	m == matrix.length
	n == matrix[i].length
	1 <= m, n <= 1000
	1 <= m * n <= 105
	-109 <= matrix[i][j] <= 109

*/
package transpose_matrix

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}

	ans := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				ans[j] = make([]int, len(matrix))
			}
			ans[j][i] = matrix[i][j]
		}
	}
	return ans
}

// --- 官方

/*
方法一: 模拟

时间复杂度：O(mn)，其中 m 和 n 分别是矩阵 matrix 的行数和列数。需要遍历整个矩阵，并对转置后的矩阵进行赋值操作。
空间复杂度：O(1)。除了返回值以外，额外使用的空间为常数。
*/
func transposeO1(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	t := make([][]int, m)
	for i := range t {
		t[i] = make([]int, n)
		for j := range t[i] {
			t[i][j] = -1
		}
	}
	for i, row := range matrix {
		for j, v := range row {
			t[j][i] = v
		}
	}
	return t
}
