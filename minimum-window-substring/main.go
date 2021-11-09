/*
Package minimum_window_substring
https://leetcode-cn.com/problems/minimum-window-substring/

76. 最小覆盖子串

给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：
	如果 s 中存在这样的子串，我们保证它是唯一的答案。

提示：
	1 <= s.length, t.length <= 105
	s 和 t 由英文字母组成

进阶：
	你能设计一个在 o(n) 时间内解决此问题的算法吗？

*/
package minimum_window_substring

import "math"

// --- 自己

/*
方法一:	滑动窗口

时间复杂度：
空间复杂度：
*/
func minWindow(s string, t string) string {
	ans := ""
	if len(t) > len(s) {
		return ans
	}

	remain := len(t)
	memo := make(map[byte]int, 0) // 用于记录字符命中
	for i := 0; i < len(t); i++ {
		memo[t[i]]++
	}
	for left, right := 0, 0; right < len(s); right++ {
		if n, ok := memo[s[right]]; ok {
			memo[s[right]]--
			if n > 0 { // 大于0才扣除
				remain--
			}
		} else {
			continue
		}
		if remain > 0 {
			continue
		}

		// 收缩
		for ; left <= right; left++ {
			if n, ok := memo[s[left]]; ok {
				memo[s[left]]++
				if n >= 0 { // 不能再收缩了
					if len(ans) == 0 {
						ans = s[left : right+1]
					} else if len(ans) > right-left+1 {
						ans = s[left : right+1]
					}
					remain++ // 多了一个未命中数量
					left++   // 退出前 移动到下一位
					break
				}
			}
		}

	}

	return ans
}

// --- 官方

/*
方法一:	滑动窗口

时间复杂度：最坏情况下左右指针对 s 的每个元素各遍历一遍，哈希表中对 s 中的每个元素各插入、删除一次，对 t 中的元素各插入一次。
	每次检查是否可行会遍历整个 t 的哈希表，哈希表的大小与字符集的大小有关，设字符集大小为 C，则渐进时间复杂度为 O(C⋅∣s∣+∣t∣)。
空间复杂度：这里用了两张哈希表作为辅助空间，每张哈希表最多不会存放超过字符集大小的键值对，我们设字符集大小为 C ，则渐进空间复杂度为 O(C)。
*/
func minWindowO1(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	mLen := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < mLen {
				mLen = r - l + 1
				ansL, ansR = l, l+mLen
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
