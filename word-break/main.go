/*
Package word_break
https://leetcode-cn.com/problems/word-break/

139. 单词拆分

给你一个字符串 s 和一个字符串列表 wordDict 作为字典，判定 s 是否可以由空格拆分为一个或多个在字典中出现的单词。

说明：
	拆分时可以重复使用字典中的单词。

提示：
	1 <= s.length <= 300
	1 <= wordDict.length <= 1000
	1 <= wordDict[i].length <= 20
	s 和 wordDict[i] 仅有小写英文字母组成
	wordDict 中的所有字符串 互不相同

*/
package word_break

// --- 自己

/*
方法一: 深度优先搜索
	// 超出时限

时间复杂度：
空间复杂度：
*/
func wordBreak(s string, wordDict []string) bool {
	hash := make(map[string]bool)
	for _, word := range wordDict {
		hash[word] = true
	}
	ans := false

	var dfs func(pos int)
	dfs = func(pos int) {
		if ans {
			return
		}
		if pos == len(s) {
			ans = true
			return
		}

		for i := pos + 1; i <= len(s); i++ {
			if ans {
				break
			}
			if hash[s[pos:i]] {
				dfs(i)
			}
		}
	}
	dfs(0)
	return ans
}

/*
方法二: 动态规划

时间复杂度：
空间复杂度：
*/
func wordBreakS2(s string, wordDict []string) bool {
	hash := make(map[string]bool)
	for _, word := range wordDict {
		hash[word] = true
	}

	dp := make([]bool, len(s)+1)
	for i := 1; i <= len(s); i++ {
		if hash[s[:i]] {
			dp[i] = true
			continue
		}
		for j := 1; j < i; j++ {
			if dp[j] && hash[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
