/*
Package game_of_life
https://leetcode-cn.com/problems/game-of-life/

289. 生命游戏

根据 百度百科 ， 生命游戏 ，简称为 生命 ，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。
每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：
	1. 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
	2. 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
	3. 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
	4. 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；

下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。
给你 m x n 网格面板 board 的当前状态，返回下一个状态。

提示：
	m == board.length
	n == board[i].length
	1 <= m, n <= 25
	board[i][j] 为 0 或 1

进阶：
	- 你可以使用原地算法解决本题吗？请注意，面板上所有格子需要同时被更新：你不能先更新某些格子，然后使用它们的更新后的值再更新其他格子。
	- 本题中，我们使用二维数组来表示面板。原则上，面板是无限的，但当活细胞侵占了面板边界时会造成问题。你将如何解决这些问题？

*/
package game_of_life

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	live := func(i, j int) {
		count := 0
		// 8个判断可优化
		if i-1 >= 0 && j-1 >= 0 && board[i-1][j-1]&0b01 == 0b01 {
			count++
		}
		if i-1 >= 0 && board[i-1][j]&0b01 == 0b01 {
			count++
		}
		if i-1 >= 0 && j+1 < n && board[i-1][j+1]&0b01 == 0b01 {
			count++
		}
		if j-1 >= 0 && board[i][j-1]&0b01 == 0b01 {
			count++
		}
		if j+1 < n && board[i][j+1]&0b01 == 0b01 {
			count++
		}
		if i+1 < m && j-1 >= 0 && board[i+1][j-1]&0b01 == 0b01 {
			count++
		}
		if i+1 < m && board[i+1][j]&0b01 == 0b01 {
			count++
		}
		if i+1 < m && j+1 < n && board[i+1][j+1]&0b01 == 0b01 {
			count++
		}

		if count < 2 || count > 3 {
			board[i][j] = board[i][j] & 0b01
		} else if count == 3 {
			board[i][j] = board[i][j] | 0b10
		} else {
			board[i][j] = board[i][j]<<1 | board[i][j]
		}

		//fmt.Println(i, j, count, board[i][j])
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			live(i, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = (board[i][j] & 0b10) >> 1
		}
	}
}
