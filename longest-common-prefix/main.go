/*
Package longest_common_prefix
https://leetcode-cn.com/problems/longest-common-prefix/

14. 最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

提示：
	1 <= strs.length <= 200
	0 <= strs[i].length <= 200
	strs[i] 仅由小写英文字母组成

*/
package longest_common_prefix

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func longestCommonPrefix(strs []string) string {
	ans := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for ; j < len(ans) && j < len(strs[i]); j++ {
			if strs[i][j] != ans[j] {
				break
			}
		}
		ans = ans[:j]
	}
	return ans
}
