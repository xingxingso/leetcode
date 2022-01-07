/*
Package spiral_matrix
https://leetcode-cn.com/problems/spiral-matrix/

54. 螺旋矩阵

给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

提示：
	m == matrix.length
	n == matrix[i].length
	1 <= m, n <= 10
	-100 <= matrix[i][j] <= 100

*/
package spiral_matrix

// --- 自己

/*
方法一：

时间复杂度：
空间复杂度：
*/
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	unUseNum := 101
	//directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	directions := []int{0, 1, 0, -1, 0}
	direction := 0

	ans := make([]int, m*n)

	for l, c, i := 0, 0, 0; matrix[l][c] != unUseNum; i++ {
		ans[i] = matrix[l][c]
		matrix[l][c] = unUseNum
		l2, c2 := l+directions[direction], c+directions[direction+1]
		if l2 < 0 || l2 >= m || c2 < 0 || c2 >= n || matrix[l2][c2] == unUseNum {
			direction = (direction + 1) % (len(directions) - 1)
		}

		l, c = l+directions[direction], c+directions[direction+1]
		//if l < 0 || l >= m || c < 0 || c >= n || matrix[l][c] == unUseNum {
		if l < 0 || l >= m || c < 0 || c >= n {
			break
		}
	}

	return ans
}
