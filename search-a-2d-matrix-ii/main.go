/*
Package search_a_2d_matrix_ii
https://leetcode-cn.com/problems/search-a-2d-matrix-ii/

240. 搜索二维矩阵 II

编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
	每行的元素从左到右升序排列。
	每列的元素从上到下升序排列。

提示：
	m == matrix.length
	n == matrix[i].length
	1 <= n, m <= 300
	-10^9 <= matrix[i][j] <= 10^9
	每行的所有元素从左到右升序排列
	每列的所有元素从上到下升序排列
	-10^9 <= target <= 10^9

*/
package search_a_2d_matrix_ii

// --- 自己

/*
方法一：

时间复杂度：
空间复杂度：
*/
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix)-1, len(matrix[0])-1
	for i, j := m, 0; i >= 0 && j <= n; {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return false
}
