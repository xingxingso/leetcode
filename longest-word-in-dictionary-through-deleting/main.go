/*
Package longest_word_in_dictionary_through_deleting
https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting/

524. 通过删除字母匹配到字典里最长单词

给你一个字符串 s 和一个字符串数组 dictionary ，找出并返回 dictionary 中最长的字符串，
该字符串可以通过删除 s 中的某些字符得到。

如果答案不止一个，返回长度最长且字母序最小的字符串。如果答案不存在，则返回空字符串。

提示：
	1 <= s.length <= 1000
	1 <= dictionary.length <= 1000
	1 <= dictionary[i].length <= 1000
	s 和 dictionary[i] 仅由小写英文字母组成

*/
package longest_word_in_dictionary_through_deleting

import (
	"sort"
)

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) > len(dictionary[j]) || len(dictionary[i]) == len(dictionary[j]) && dictionary[i] <= dictionary[j]
	})
	//fmt.Println(dictionary)

	for _, v := range dictionary {
		if len(v) > len(s) {
			continue
		}
		i := 0
		for j := 0; i < len(v) && j < len(s) && len(s)-j >= len(v)-i; {
			if v[i] == s[j] {
				i++
				j++
			} else {
				j++
			}
		}
		if i == len(v) {
			return v
		}
	}

	return ""
}
