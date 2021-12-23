/*
Package unique_binary_search_trees
https://leetcode-cn.com/problems/unique-binary-search-trees/

96. 不同的二叉搜索树

给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

*/
package unique_binary_search_trees

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func numTrees(n int) int {
	dp := make(map[int]int, n)
	dp[0] = 1
	dp[1] = 1
	for k := 2; k <= n; k++ {
		for i := 1; i <= k; i++ {
			dp[k] += dp[i-1] * dp[k-i]
		}
	}

	return dp[n]
}
