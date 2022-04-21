/*
Package pascals_triangle
https://leetcode-cn.com/problems/pascals-triangle/

118. 杨辉三角

给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

提示:
	1 <= numRows <= 30

*/
package pascals_triangle

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	ans := [][]int{{1}, {1, 1}}
	for i := 2; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		for j := 0; j < len(ans[i-1])-1; j++ {
			temp[j+1] = ans[i-1][j] + ans[i-1][j+1]
		}
		temp[len(temp)-1] = 1
		ans = append(ans, temp)
	}

	return ans
}
