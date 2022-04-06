/*
Package reverse_words_in_a_string_iii
https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/

557. 反转字符串中的单词 III

给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

提示：
	1 <= s.length <= 5 * 10^4
	s 包含可打印的 ASCII 字符。
	s 不包含任何开头或结尾空格。
	s 里 至少 有一个词。
	s 中的所有单词都用一个空格隔开。

*/
package reverse_words_in_a_string_iii

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func reverseWords(s string) string {
	b := []byte(s)
	i, j := 0, 0
	for ; i < len(b); i++ {
		if b[i] != ' ' {
			continue
		}
		reverse(b[j:i])
		for ; i < len(b); i++ {
			if b[i] != ' ' {
				break
			}
		}
		j = i
		i--
	}
	reverse(b[j:])
	return string(b)
}

func reverse(b []byte) {
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
}
