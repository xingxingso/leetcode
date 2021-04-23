/*
https://leetcode-cn.com/problems/stone-game-v/

1563. 石子游戏 V

几块石子 排成一行 ，每块石子都有一个关联值，关联值为整数，由数组 stoneValue 给出。
游戏中的每一轮：Alice 会将这行石子分成两个 非空行（即，左侧行和右侧行）；Bob 负责计算每一行的值，即此行中所有石子的值的总和。
Bob 会丢弃值最大的行，Alice 的得分为剩下那行的值（每轮累加）。如果两行的值相等，Bob 让 Alice 决定丢弃哪一行。下一轮从剩下的那一行开始。
只 剩下一块石子 时，游戏结束。Alice 的分数最初为 0 。
返回 Alice 能够获得的最大分数 。

提示：
	1 <= stoneValue.length <= 500
	1 <= stoneValue[i] <= 10^6

*/
package stone_game_v

// --- 自己

/*
方法一: 动态规划
	!!! 超出时间限制

时间复杂度：
空间复杂度：
*/
func stoneGameV(stoneValue []int) int {
	l := len(stoneValue)
	dp := make([][]int, l)

	for i := l - 1; i >= 0; i-- {
		dp[i] = make([]int, l)
		for j := i + 1; j < l; j++ {
			for k := i; k < j; k++ {
				left, right := sum(stoneValue[i:k+1]), sum(stoneValue[k+1:j+1])
				if left > right {
					dp[i][j] = max(dp[i][j], dp[k+1][j]+right)
				} else if left == right {
					dp[i][j] = max(dp[i][j], max(dp[i][k], dp[k+1][j])+left)
				} else {
					dp[i][j] = max(dp[i][j], dp[i][k]+left)
				}
			}
		}
	}
	return dp[0][l-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func sum(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans += v
	}

	return ans
}

// --- 官方

/*
方法一: 动态规划

时间复杂度：O(n^3)，其中 n 是数组 stoneValue 的长度。
空间复杂度：O(n^2)，为存储所有状态需要的空间。
*/
func stoneGameVO1(stoneValue []int) int {
	n := len(stoneValue)
	var dfs func(left, right int) int
	// f[i][j] 为 下标 [i,j] 的最大分数
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	dfs = func(left, right int) int {
		if left == right {
			return 0
		}
		if f[left][right] != 0 {
			return f[left][right]
		}
		s := 0
		for i := left; i <= right; i++ {
			s += stoneValue[i]
		}
		sumL := 0
		for i := left; i < right; i++ {
			sumL += stoneValue[i]
			sumR := s - sumL
			if sumL < sumR {
				f[left][right] = max(f[left][right], dfs(left, i)+sumL)
			} else if sumL > sumR {
				f[left][right] = max(f[left][right], dfs(i+1, right)+sumR)
			} else {
				f[left][right] = max(f[left][right], max(dfs(left, i), dfs(i+1, right))+sumL)
			}
		}
		return f[left][right]
	}

	return dfs(0, n-1)
}

/*
方法二: 动态规划优化
	// 没看懂， 先不写了

时间复杂度：O(n^2)，其中 n 是数组 stoneValue 的长度。
空间复杂度：O(n^2)，为存储所有状态以及辅助数组需要的空间。
*/
// func stoneGameVO2(stoneValue []int) int {
//
// }
