/*
Package decode_string
https://leetcode-cn.com/problems/decode-string/

394. 字符串解码

给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

提示：
	1 <= s.length <= 30
	s 由小写英文字母、数字和方括号 '[]' 组成
	s 保证是一个 有效 的输入。
	s 中所有整数的取值范围为 [1, 300]

*/
package decode_string

// --- 自己

/*
方法一:
	// 不支持嵌套

时间复杂度：
空间复杂度：
*/
func decodeStringS1(s string) string {
	ans := make([]byte, 0)
	symbols := make([]byte, 0)
	count := 0
	for i := 0; i < len(s); i++ {
		v := s[i]
		switch {
		case v >= 'a' && v <= 'z':
			symbols = append(symbols, v)
		case v >= '0' && v <= '9':
			if len(symbols) > 0 {
				ans = writeSymbols(ans, symbols, 1)
				symbols = symbols[:0]
			}
			count = count*10 + int(v-'0')
		case v == '[':
			symbols = symbols[:0]
		case v == ']':
			ans = writeSymbols(ans, symbols, count)
			count = 0
			symbols = symbols[:0]
			// fmt.Println(symbols, len(symbols))
		}
		// fmt.Println(string(ans), i)
	}
	if len(symbols) > 0 {
		ans = writeSymbols(ans, symbols, 1)
		symbols = symbols[:0]
	}
	return string(ans)
}

func writeSymbols(target, symbols []byte, n int) []byte {
	if len(symbols) == 0 {
		return target
	}
	for i := 0; i < n; i++ {
		target = append(target, symbols...)
	}
	return target
}

/*
方法一: 递归

时间复杂度：
空间复杂度：
*/
func decodeStringS2(s string) string {
	var dfs func(pos int) (string, int)
	dfs = func(pos int) (string, int) {
		if pos >= len(s) {
			return "", 0
		}
		ans := ""
		for pos < len(s) {
			pre := make([]byte, 0)
			for pos < len(s) && s[pos] >= 'a' && s[pos] <= 'z' {
				pre = append(pre, s[pos])
				pos++
			}
			ans += string(pre)
			if pos < len(s) && s[pos] == ']' {
				return ans, pos
			}

			count := 0
			for pos < len(s) && s[pos] >= '0' && s[pos] <= '9' {
				count = count*10 + int(s[pos]-'0')
				pos++
			}

			str := ""
			if pos < len(s) && s[pos] == '[' {
				pos++
				str, pos = dfs(pos)
				pos++
			}

			//fmt.Println(str, ans, count, pos)
			for i := 0; i < count; i++ {
				ans += str
			}
		}

		return ans, pos
	}

	ans, _ := dfs(0)
	return ans
}
