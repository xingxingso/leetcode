/*
Package sliding_puzzle
https://leetcode-cn.com/problems/sliding-puzzle/

773. 滑动谜题

在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用 0 来表示.
一次移动定义为选择 0 与一个相邻的数字（上下左右）进行交换.
最终当板 board 的结果是 [[1,2,3],[4,5,0]] 谜板被解开。
给出一个谜板的初始状态，返回最少可以通过多少次移动解开谜板，如果不能解开谜板，则返回 -1 。

提示：
	board 是一个如上所述的 2 x 3 的数组.
	board[i][j] 是一个 [0, 1, 2, 3, 4, 5] 的排列.

*/
package sliding_puzzle

// --- 自己

/*
方法一: BFS

时间复杂度：
空间复杂度：
*/
func slidingPuzzle(board [][]int) int {
	transTarget := trans([][]int{{1, 2, 3}, {4, 5, 0}})
	queue := make([][][]int, 0)
	queue = append(queue, board)
	visited := make(map[uint32]bool)
	visited[trans(board)] = true
	step := 0
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if trans(queue[i]) == transTarget {
				return step
			}
			for _, v := range move(queue[i]) {
				transV := trans(v)
				if visited[transV] {
					continue
				}
				queue = append(queue, v)
				visited[transV] = true
			}
		}
		step++
		queue = queue[l:]
	}
	return -1
}

// 压缩下作为 hash key
func trans(board [][]int) uint32 {
	var res uint32 = 0
	pos := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			temp := uint32(board[i][j] << pos)
			res |= temp
			pos += 4
		}
	}
	return res
}

func move(board [][]int) [][][]int {
	res := make([][][]int, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				// 上 下
				temp := make([][]int, 2)
				copy2(temp, board)
				temp[0][j], temp[1][j] = temp[1][j], temp[0][j]
				res = append(res, temp)

				// 左
				if j > 0 {
					temp := make([][]int, 2)
					copy2(temp, board)
					temp[i][j], temp[i][j-1] = temp[i][j-1], temp[i][j]
					res = append(res, temp)
				}
				// 右
				if j < 2 {
					temp := make([][]int, 2)
					copy2(temp, board)
					temp[i][j], temp[i][j+1] = temp[i][j+1], temp[i][j]
					res = append(res, temp)
				}

				return res
			}
		}
	}
	return nil
}

func copy2(dst, src [][]int) {
	for key, value := range src {
		dst[key] = make([]int, len(value))
		copy(dst[key], value)
	}
}
