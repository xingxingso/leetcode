/*
Package longest_substring_without_repeating_characters
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

3. 无重复字符的最长子串

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

提示：
	0 <= s.length <= 5 * 10^4
	s 由英文字母、数字、符号和空格组成

*/
package longest_substring_without_repeating_characters

// --- 自己

/*
方法一:	滑动窗口
	尝试用下 labuladong 的公式
	https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485141&idx=1&sn=0e4583ad935e76e9a3f6793792e60734&chksm=9bd7f8ddaca071cbb7570b2433290e5e2628d20473022a5517271de6d6e50783961bebc3dd3b&scene=21#wechat_redirect

时间复杂度：
空间复杂度：
*/
func lengthOfLongestSubstring(s string) int {
	ans := 0
	window := map[byte]int{}
	for left, right := 0, 0; right < len(s); right++ {
		ch := s[right]
		// 进行窗口内数据的一系列更新
		window[ch]++
		// 有重复字符的时候缩小窗口
		for window[ch] > 1 {
			// 进行窗口内数据的一系列更新
			window[s[left]]--
			left++
		}
		// 在这里更新答案
		ans = max(ans, right-left+1)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法一:	滑动窗口
	和 lengthOfLongestSubstring 相比这里记录的是字符的下标，不需要遍历 map

时间复杂度：
空间复杂度：
*/
func lengthOfLongestSubstringS2(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := 0
	window := make(map[byte]int)
	for left, right := 0, 0; right < len(s); right++ {
		if idx, ok := window[s[right]]; ok && idx >= left {
			left = idx + 1 // idx >= left 说明 在窗口中, 此时窗口左边界直接移动到 idx + 1
			// ans := max(ans, right-left+1) // 这种情况下 长度不可能大于原窗口
		} else {
			ans = max(ans, right-left+1)
		}
		window[s[right]] = right // 更新最后出现的位置
	}

	return ans
}
