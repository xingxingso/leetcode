/*
https://leetcode-cn.com/problems/delete-operation-for-two-strings/

583. 两个字符串的删除操作

给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。

提示：
	1. 给定单词的长度不超过500。
	2. 给定单词中的字符只含有小写字母。

*/
package delete_operation_for_two_strings

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func minDistance(word1 string, word2 string) int {
	// dp := make([][]int, len(word1)+1)
	d := make([]int, len(word2)+1)

	for i := 1; i <= len(word1); i++ {
		pre := 0
		for j := 1; j <= len(word2); j++ {
			temp := d[j]
			if word1[i-1] == word2[j-1] {
				// dp[i][j] = dp[i-1][j-1] + 1
				d[j] = pre + 1
			} else {
				// dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				d[j] = max(d[j], d[j-1])
			}
			pre = temp
		}
	}

	// return dp[len(word1)][len(word2)]
	return len(word1) + len(word2) - 2*d[len(word2)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
