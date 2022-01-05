/*
Package excel_sheet_column_title
https://leetcode-cn.com/problems/excel-sheet-column-title/

168. Excel表列名称

给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。

提示：
	1 <= columnNumber <= 2^31 - 1

*/
package excel_sheet_column_title

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func convertToTitle(columnNumber int) string {
	ans := make([]byte, 0)

	for columnNumber > 0 {
		cur := (columnNumber - 1) % 26
		ans = append(ans, byte(cur)+'A')
		columnNumber = (columnNumber - 1) / 26
	}

	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return string(ans)
}
