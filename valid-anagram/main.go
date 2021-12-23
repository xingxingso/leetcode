/*
Package valid_anagram
https://leetcode-cn.com/problems/valid-anagram/

242. 有效的字母异位词

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：
	若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

提示:
	1 <= s.length, t.length <= 5 * 10^4
	s 和 t 仅包含小写字母

进阶:
	如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

*/
package valid_anagram

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	memo := [26]int{}
	for i := 0; i < len(s); i++ {
		memo[s[i]-'a']++
		memo[t[i]-'a']--
	}

	for _, v := range memo {
		if v > 0 {
			return false
		}
	}

	return true
}
