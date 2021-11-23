/*
Package word_search
https://leetcode-cn.com/problems/word-search/

79. 单词搜索

给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，
其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。

提示：
	m == board.length
	n = board[i].length
	1 <= m, n <= 6
	1 <= word.length <= 15
	board 和 word 仅由大小写英文字母组成

进阶：
	你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

*/
package word_search

// ---自己

/*
方法一: 回溯

时间复杂度：
空间复杂度：
*/
func exist(board [][]byte, word string) bool {
	direction := []int{-1, 0, 1, 0, -1}
	ans := false
	m, n, l := len(board), len(board[0]), len(word)
	var dfs func(i, j int, pos int)
	dfs = func(i, j int, pos int) {
		if pos == l {
			ans = true
			return
		}
		for k := 0; k < 4; k++ {
			p, q := i+direction[k], j+direction[k+1]
			if p >= 0 && p < m && q >= 0 && q < n && board[p][q] == word[pos] {
				c := board[p][q]
				board[p][q] = 0
				dfs(p, q, pos+1)
				board[p][q] = c
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ans {
				return ans
			}
			c := board[i][j]
			if c == word[0] {
				board[i][j] = 0
				dfs(i, j, 1)
				board[i][j] = c
			}
		}
	}
	return ans
}
