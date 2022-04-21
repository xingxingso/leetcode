/*
Package pascals_triangle_ii
https://leetcode-cn.com/problems/pascals-triangle-ii/

119. 杨辉三角 II

给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

提示:
	0 <= rowIndex <= 33

进阶：
	你可以优化你的算法到 O(rowIndex) 空间复杂度吗？

*/
package pascals_triangle_ii

// --- 官方

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func getRowO1(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}

/*
方法二: 线性递推

时间复杂度：
空间复杂度：
*/
func getRowO2(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}
