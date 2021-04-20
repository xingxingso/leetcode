/*
https://leetcode-cn.com/problems/stone-game-vi/

1686. 石子游戏 VI

Alice 和 Bob 轮流玩一个游戏，Alice 先手。
一堆石子里总共有 n 个石子，轮到某个玩家时，他可以 移出 一个石子并得到这个石子的价值。Alice 和 Bob 对石子价值有 不一样的的评判标准 。
双方都知道对方的评判标准。给你两个长度为 n 的整数数组 aliceValues 和 bobValues 。aliceValues[i] 和 bobValues[i] 分别表示
Alice 和 Bob 认为第 i 个石子的价值。所有石子都被取完后，得分较高的人为胜者。如果两个玩家得分相同，那么为平局。
两位玩家都会采用 最优策略 进行游戏。
请你推断游戏的结果，用如下的方式表示：
	如果 Alice 赢，返回 1 。
	如果 Bob 赢，返回 -1 。
	如果游戏平局，返回 0 。

提示：
	n == aliceValues.length == bobValues.length
	1 <= n <= 10^5
	1 <= aliceValues[i], bobValues[i] <= 100

*/
package stone_game_vi

import (
	"sort"
)

// --- 他人

// https://leetcode-cn.com/problems/stone-game-vi/solution/tan-xin-by-o0o0o0o0o0-rzvp/
/*
方法一: 贪心算法
	alice - bob = sum(A) - (sum[a] + sum[b] +sum[c]....)
	a, b, c 为其中最大值

时间复杂度：
空间复杂度：
*/
func stoneGameVI(aliceValues []int, bobValues []int) int {
	sum := make([]int, len(aliceValues))
	sumA := 0
	for i := 0; i < len(aliceValues); i++ {
		sum[i] = aliceValues[i] + bobValues[i]
		sumA += aliceValues[i]
	}
	sort.Slice(sum, func(i, j int) bool {
		return sum[i] > sum[j]
	})
	// fmt.Println(sum)
	b := 0
	for i, v := range sum {
		if i%2 != 0 {
			b += v
		}
	}
	// fmt.Println(sumA, b)
	diff := sumA - b
	if diff > 0 {
		return 1
	} else if diff == 0 {
		return 0
	} else {
		return -1
	}
}
