/*
Package n_queens
https://leetcode-cn.com/problems/n-queens/

51. N 皇后

n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

提示：
	1 <= n <= 9
	皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。

*/
package n_queens

import (
	"math/bits"
	"strings"
)

// --- 自己

/*
方法一: 回溯

时间复杂度：
空间复杂度：
*/
func solveNQueens(n int) [][]string {
	var ans [][]string

	var backTrack func(row int, board []string)
	backTrack = func(row int, board []string) {
		if row == n {
			newList := make([]string, n)
			copy(newList, board)
			ans = append(ans, newList)
			// ans = append(ans, board) // 直接这么写是个坑
			return
		}

		// 第n 行，尝试每一列
		for i := 0; i < n; i++ {
			if !isValid(board, i) {
				continue
			}
			board = append(board, draw(i, n))
			backTrack(row+1, board)
			board = board[:len(board)-1]
		}
	}

	backTrack(0, []string{})
	return ans
}

// 将 皇后 放在 p 列是否合法
// 行为 len(board) 位置，已保证不在同一行
func isValid(board []string, col int) bool {
	row := len(board)
	for r, l := range board {
		for c, v := range l {
			if v != 'Q' {
				continue
			}

			// 同一列
			if c == col {
				return false
			}
			// 对角线: 和或者差相等
			if r+c == row+col || r-c == row-col {
				return false
			}
		}
	}

	return true
}

func draw(col, n int) string {
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		if i == col {
			res[i] = 'Q'
		} else {
			res[i] = '.'
		}
	}

	return string(res)
}

// --- 官方
// https://leetcode-cn.com/problems/n-queens/solution/nhuang-hou-by-leetcode-solution/

/*
方法一: 基于集合的回溯

时间复杂度：O(N!)，其中 N 是皇后数量
空间复杂度：O(N)，其中 N 是皇后数量。空间复杂度主要取决于递归调用层数、记录每行放置的皇后的列下标的数组以及三个集合，
	递归调用层数不会超过 N，数组的长度为 N，每个集合的元素个数都不会超过 N。
*/
func solveNQueensO1(n int) [][]string {
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtrack(queens, n, 0, columns, diagonals1, diagonals2)
	return solutions
}

var solutions [][]string

func backtrack(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	for i := 0; i < n; i++ {
		if columns[i] {
			continue
		}
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		backtrack(queens, n, row+1, columns, diagonals1, diagonals2)
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

/*
方法二: 基于位运算的回溯

时间复杂度：O(N!)，其中 N 是皇后数量
空间复杂度：O(N)，其中 N 是皇后数量。由于使用位运算表示，因此存储皇后信息的空间复杂度是 O(1)，
	空间复杂度主要取决于递归调用层数和记录每行放置的皇后的列下标的数组，递归调用层数不会超过 N，数组的长度为 N。
*/
func solveNQueensO2(n int) [][]string {
	solutions2 = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	solve(queens, n, 0, 0, 0, 0)
	return solutions2
}

var solutions2 [][]string

func solve(queens []int, n, row, columns, diagonals1, diagonals2 int) {
	if row == n {
		board := generateBoard2(queens, n)
		solutions2 = append(solutions2, board)
		return
	}
	availablePositions := ((1 << n) - 1) & (^(columns | diagonals1 | diagonals2))
	for availablePositions != 0 {
		position := availablePositions & (-availablePositions)
		availablePositions = availablePositions & (availablePositions - 1)
		column := bits.OnesCount(uint(position - 1))
		queens[row] = column
		solve(queens, n, row+1, columns|position, (diagonals1|position)>>1, (diagonals2|position)<<1)
		queens[row] = -1
	}
}

func generateBoard2(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

/*
labuladong 44/562
方法一: 回溯算法

时间复杂度：
空间复杂度：
*/
func solveNQueensP1(n int) [][]string {
	// '.' 表示空，'Q' 表示皇后，初始化空棋盘。
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = []byte(strings.Repeat(".", n))
	}
	var res [][]string
	// 路径:board 中小于 row 的那些行都已经成功放置了皇后
	// 选择列表:第 row 行的所有列都是放置皇后的选择
	// 结束条件:row 超过 board 的最后一行
	var backTrack func(board [][]byte, row int)
	backTrack = func(board [][]byte, row int) {
		// 触发结束条件
		if row == len(board) {
			temp := make([]string, row)
			for i := 0; i < len(board); i++ {
				temp[i] = string(board[i])
			}
			res = append(res, temp)
			return
		}

		nn := len(board[row])
		for col := 0; col < nn; col++ {
			// 排除不合法选择
			if !isValid2(board, row, col) {
				continue
			}
			// 做选择
			board[row][col] = 'Q'
			// 进入下一行决策
			backTrack(board, row+1)
			// 撤销选择
			board[row][col] = '.'
		}
	}

	backTrack(board, 0)
	return res
}

/* 是否可以在 board[row][col] 放置皇后? */
func isValid2(board [][]byte, row, col int) bool {
	n := len(board)
	// 检查列是否有皇后互相冲突
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 检查左上方是否有皇后互相冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}
