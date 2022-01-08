/*
Package minimum_knight_moves
https://leetcode-cn.com/problems/minimum-knight-moves/

1197. 进击的骑士

一个坐标可以从 -infinity 延伸到 +infinity 的 无限大的 棋盘上，你的 骑士 驻扎在坐标为 [0, 0] 的方格里。
骑士的走法和中国象棋中的马相似，走 “日” 字：即先向左（或右）走 1 格，再向上（或下）走 2 格；或先向左（或右）走 2 格，再向上（或下）走 1 格。
每次移动，他都可以按图示八个方向之一前进。
现在，骑士需要前去征服坐标为 [x, y] 的部落，请你为他规划路线。
最后返回所需的最小移动次数即可。本题确保答案是一定存在的。

提示：
	|x| + |y| <= 300

*/
package minimum_knight_moves

import "fmt"

// --- 他人

/*
方法一:
https://leetcode-cn.com/problems/minimum-knight-moves/solution/di-ke-si-te-la-suan-fa-shu-xue-fa-can-kao-liao-pin/

时间复杂度：
空间复杂度：
*/
func minKnightMovesP1(x int, y int) int {
	// 以原点为中心，四个象限都是对称的，所以 x,y的符号不会对路径的长度产生影响
	// x, y 直接取绝对值，简单处理
	x, y = abs(x), abs(y)
	// 特殊情况判断
	if x+y == 1 {
		return 3
	}
	// (x + 1)/2，(y + 1)/2 代表 走到x，y所需要的最小步数
	// (x + y + 2)/3  每次最多走三步，所以走到 x+y 的最小步数为 (x + y + 2)/3
	res := max((x+1)/2, (y+1)/2, (x+y+2)/3)
	// 结论：如果x，y同奇同偶，只有偶数步 才能走到x，y点，如果一奇一偶，则必须走 奇数步才能走到x，y点 （细品）
	res += (res ^ x ^ y) & 1
	return res
}

func max(x, y, z int) int {
	if x > y {
		if x > z {
			return x
		}
		return z
	}

	if y > z {
		return y
	}

	return z
}

/*
方法一:
https://leetcode-cn.com/problems/minimum-knight-moves/solution/hua-tu-huo-qu-yi-ge-xiao-fan-wei-de-nei-jie-zai-bi/

时间复杂度：
空间复杂度：
*/
func minKnightMovesP2(x int, y int) int {
	x = abs(x)
	y = abs(y)

	arr := [6][6]int{
		{0, 3, 2, 3, 2, 3},
		{3, 2, 1, 2, 3, 4},
		{2, 1, 4, 3, 2, 3},
		{3, 2, 3, 2, 3, 4},
		{2, 3, 2, 3, 4, 3},
		{3, 4, 3, 4, 3, 4},
	}
	ans := 0
	for x > 5 || y > 5 {
		if x > y {
			x -= 2
			if y < 2 {
				y++
			} else {
				y--
			}
		} else {
			y -= 2
			if x < 2 {
				x++
			} else {
				x--
			}
		}
		ans++
	}
	fmt.Println(x, y)
	return ans + arr[x][y]
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
