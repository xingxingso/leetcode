/*
https://leetcode-cn.com/problems/edit-distance/

72. 编辑距离

给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数。
你可以对一个单词进行如下三种操作：
	插入一个字符
	删除一个字符
	替换一个字符

提示：
	0 <= word1.length, word2.length <= 500
	word1 和 word2 由小写英文字母组成

*/
package edit_distance

import "fmt"

// --- 他人

/*
! 超出时限
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func minDistance(word1 string, word2 string) int {
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		} else {
			return minN(dp(i, j-1)+1, dp(i-1, j)+1, dp(i-1, j-1)+1)
		}
	}

	return dp(len(word1)-1, len(word2)-1)
}

func minN(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}

/*
方法一: 动态规划优化

时间复杂度：
空间复杂度：
*/
func minDistanceS1(word1 string, word2 string) int {
	memo := make(map[string]int)
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if n, ok := memo[fmt.Sprintf("%d,%d", i, j)]; ok {
			return n
		}
		if i == -1 {
			memo[fmt.Sprintf("%d,%d", i, j)] = j + 1
			return j + 1
		}
		if j == -1 {
			memo[fmt.Sprintf("%d,%d", i, j)] = i + 1
			return i + 1
		}
		if word1[i] == word2[j] {
			n := dp(i-1, j-1)
			memo[fmt.Sprintf("%d,%d", i, j)] = n
			return n
		} else {
			//           插入        删除           替换
			n := minN(dp(i, j-1)+1, dp(i-1, j)+1, dp(i-1, j-1)+1)
			memo[fmt.Sprintf("%d,%d", i, j)] = n
			return n
		}
	}

	return dp(len(word1)-1, len(word2)-1)
}

/*
方法一: 自底向上

时间复杂度：O(mn)，其中 m 为 word1 的长度，n 为 word2 的长度。
空间复杂度：O(mn)，我们需要大小为 O(mn) 的 DD 数组来记录状态值。
*/
func minDistanceS2(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}
	for i := 1; i <= l2; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minN(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[l1][l2]
}
