/*
Package range_sum_query_2d_immutable
https://leetcode-cn.com/problems/range-sum-query-2d-immutable/

304. 二维区域和检索 - 矩阵不可变

给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。

提示：
	m == matrix.length
	n == matrix[i].length
	1 <= m, n <= 200
	-10^5 <= matrix[i][j] <= 10^5
	0 <= row1 <= row2 < m
	0 <= col1 <= col2 < n
	最多调用 10^4 次 sumRegion 方法

*/
package range_sum_query_2d_immutable

//type NumMatrix struct {}
//func Constructor(matrix [][]int) NumMatrix {}
//func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

// --- 自己

/*
方法一: 积分图
	integral[i][j] 表示以位置 (0, 0) 为左上角、位置 (i-1, j-1) 为右下角的长方形中所有数字的和。
	integral[i][j]=matrix[i-1][j-1]+integral[i-1][j]+integral[i][j-1]-integral[i-1][j-1]

时间复杂度：
空间复杂度：
*/

type NumMatrix struct {
	integral [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	integral := make([][]int, m+1)
	integral[0] = make([]int, n+1)
	for i := 1; i <= m; i++ {
		integral[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			integral[i][j] = matrix[i-1][j-1] + integral[i-1][j] + integral[i][j-1] - integral[i-1][j-1]
		}
	}

	return NumMatrix{
		integral: integral,
	}
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return n.integral[row2+1][col2+1] - n.integral[row2+1][col1] - n.integral[row1][col2+1] + n.integral[row1][col1]
}

/*
方法二:

时间复杂度：
空间复杂度：
*/

type NumMatrixS2 struct {
	prefixSum [][]int
}

func ConstructorS2(matrix [][]int) NumMatrixS2 {
	prefixSum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		prefixSum[i] = make([]int, len(matrix[i])+1)
		for j := 0; j < len(matrix[i]); j++ {
			prefixSum[i][j+1] = prefixSum[i][j] + matrix[i][j]
		}
	}

	return NumMatrixS2{
		prefixSum: prefixSum,
	}
}

func (n *NumMatrixS2) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := 0
	for i := row1; i <= row2; i++ {
		ans += n.prefixSum[i][col2+1] - n.prefixSum[i][col1]
	}

	return ans
}
