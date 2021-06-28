/*
Package sudoku_solver
https://leetcode-cn.com/problems/sudoku-solver/

37. 解数独

编写一个程序，通过填充空格来解决数独问题。
数独的解法需 遵循如下规则：
	1. 数字 1-9 在每一行只能出现一次。
	2. 数字 1-9 在每一列只能出现一次。
	3. 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。

提示：
	board.length == 9
	board[i].length == 9
	board[i][j] 是一位数字或者 '.'
	题目数据 保证 输入数独仅有一个解

*/
package sudoku_solver

import (
	"fmt"
)

// --- 自己

/*
方法一: 回溯

时间复杂度：
空间复杂度：
*/
func solveSudoku(board [][]byte) {
	isFind := false
	var backTrack func(i, j int)
	backTrack = func(i, j int) {
		if isFind {
			return
		}
		if i == 9 && j == 0 {
			isFind = true
			return
		}

		if board[i][j] != '.' {
			if j == 8 {
				i++
				j = -1
			}
			backTrack(i, j+1)
			return
		}
		var k byte
		for k = '1'; k <= '9'; k++ {
			if isFind {
				return
			}
			board[i][j] = k
			if isValid(board, i, j) {
				p, q := i, j
				if q == 8 {
					p++
					q = -1
				}
				backTrack(p, q+1)
			}
			if !isFind {
				board[i][j] = '.'
			}
		}
	}
	// printBoard(board)
	backTrack(0, 0)
	// fmt.Println("after")
	// printBoard(board)
}

func isValid(board [][]byte, p, q int) bool {
	// 每行 每列
	for i := 0; i < 9; i++ {
		if q != i && board[p][i] == board[p][q] {
			return false
		}
		if p != i && board[i][q] == board[p][q] {
			return false
		}
	}
	// 每块
	w, h := p/3, q/3
	for i := w * 3; i < 3*w+3; i++ {
		for j := h * 3; j < 3*h+3; j++ {
			if i == p && j == q {
				continue
			}
			if board[i][j] == board[p][q] {
				return false
			}
		}
	}

	return true
}

func printBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Println()
	}
}
