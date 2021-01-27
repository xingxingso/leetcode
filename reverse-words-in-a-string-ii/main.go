/*
https://leetcode-cn.com/problems/reverse-words-in-a-string-ii/

186. 翻转字符串里的单词 II

给定一个字符串，逐个翻转字符串中的每个单词。

注意：
	单词的定义是不包含空格的一系列字符
	输入字符串中不会包含前置或尾随的空格
	单词与单词之间永远是以单个空格隔开的

进阶：
	使用 O(1) 额外空间复杂度的原地解法。

*/
package reverse_words_in_a_string_ii

// --- 自己

/*
方法一:
	the sky is blue
	各个单词翻转
	eht yks si eulb
	整体翻转
	blue is sky the

时间复杂度：O(N)
空间复杂度：O(1)
*/
func reverseWords(s []byte) {
	// fmt.Printf("%c\n", s)

	var wordStart, wordEnd int
	wordStart = -1
	for i := 0; i < len(s); i++ {
		if (s[i] == ' ' || i == len(s)-1) && wordStart != -1 {
			wordEnd = i - 1
			if s[i] != ' ' {
				wordEnd = i
			}
			for wordStart < wordEnd {
				tmp := s[wordStart]
				s[wordStart] = s[wordEnd]
				s[wordEnd] = tmp
				wordStart++
				wordEnd--
			}
			wordStart = -1
		} else if wordStart == -1 {
			wordStart = i
		}
	}

	wordStart, wordEnd = 0, len(s)-1

	// fmt.Printf("%c\n", s)

	for wordStart < wordEnd {
		tmp := s[wordStart]
		s[wordStart] = s[wordEnd]
		s[wordEnd] = tmp
		wordStart++
		wordEnd--
	}
	// fmt.Printf("%c\n", s)
}
