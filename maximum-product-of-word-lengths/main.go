/*
Package maximum_product_of_word_lengths
https://leetcode-cn.com/problems/maximum-product-of-word-lengths/

318. 最大单词长度乘积

给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j]) 的最大值，并且这两个单词不含有公共字母。如果不存在这样的两个单词，返回 0 。

提示：
	2 <= words.length <= 1000
	1 <= words[i].length <= 1000
	words[i] 仅包含小写字母

*/
package maximum_product_of_word_lengths

// --- 他人

/*
方法一:
	leetcode101 10.3

时间复杂度：
空间复杂度：
*/
func maxProduct(words []string) int {
	hash := make(map[int]int) // mask => size
	ans := 0
	for _, word := range words {
		size := len(word)
		mask := 0
		for i := 0; i < size; i++ {
			mask |= 1 << (word[i] - 'a')
		}

		for m, s := range hash {
			if m&mask == 0 {
				ans = max(ans, s*size)
			}
		}

		hash[mask] = max(hash[mask], size) // 要使用最大的size, 比如 "fo", "foo", mask 相同，要取长度长的 foo
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
