/*
Package permutation_in_string
https://leetcode-cn.com/problems/permutation-in-string/

567. 字符串的排列

给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
换句话说，第一个字符串的排列之一是第二个字符串的子串。

提示：
	输入的字符串只包含小写字母
	两个字符串的长度都在 [1, 10,000] 之间
*/
package permutation_in_string

// --- 自己

/*
方法一: 双指针

时间复杂度：
空间复杂度：
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s1); i++ {
		ori[s1[i]]++
	}

	var check func() bool
	check = func() bool {
		for k, v := range ori {
			if cnt[k] != v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < len(s2); r++ {
		if _, ok := ori[s2[r]]; !ok {
			cnt = map[byte]int{}
			l = r
			continue
		} else {
			cnt[s2[r]]++
		}

		for ; r-l >= len(s1); l++ {
			if _, ok := cnt[s2[l]]; ok {
				cnt[s2[l]]--
			}
		}
		if check() {
			return true
		}
	}

	return false
}

// --- 官方

/*
方法一: 滑动窗口 未优化

时间复杂度：
空间复杂度：
*/
func checkInclusionO1(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var cnt1, cnt2 [26]int
	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := n; i < m; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

/*
方法一: 滑动窗口

时间复杂度：O(n+m+∣Σ∣)，其中 n 是字符串 s1 的长度，m 是字符串 s2 的长度，Σ 是字符集，这道题中的字符集是小写字母，∣Σ∣=26。
空间复杂度：O(∣Σ∣)。
*/
func checkInclusionO2(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for i, ch := range s1 {
		cnt[ch-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, c := range cnt[:] {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		x, y := s2[i]-'a', s2[i-n]-'a'
		if x == y {
			continue
		}
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}

/*
方法二: 双指针

时间复杂度：O(n+m+∣Σ∣)。创建 cnt 需要 O(∣Σ∣) 的时间。 遍历 s1 需要 O(n) 的时间。
	双指针遍历 s2 时，由于 left 和 right 都只会向右移动，故这一部分需要 O(m) 的时间。
空间复杂度：O(∣Σ∣)。
*/
func checkInclusionO3(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		cnt[x]++
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}
