/*
https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/

438. 找到字符串中所有字母异位词

给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：
	字母异位词指字母相同，但排列不同的字符串。
	不考虑答案输出的顺序。

*/
package find_all_anagrams_in_a_string

// --- 自己

/*
方法一: 双指针

时间复杂度：
空间复杂度：
*/
func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	sLen, pLen := len(s), len(p)
	if sLen < pLen || pLen == 0 {
		return ans
	}

	cnt := [26]int{}
	for _, ch := range p {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s {
		x := ch - 'a'
		cnt[x]++
		for cnt[x] > 0 {
			cnt[s[left]-'a']--
			left++
		}
		if right-left+1 == pLen {
			ans = append(ans, left)
		}
	}
	return ans
}
