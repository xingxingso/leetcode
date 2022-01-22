/*
Package rearrange_string_k_distance_apart
https://leetcode-cn.com/problems/rearrange-string-k-distance-apart/

358. K 距离间隔重排字符串

给你一个非空的字符串 s 和一个整数 k，你要将这个字符串中的字母进行重新排列，使得重排后的字符串中相同字母的位置间隔距离至少为 k。
所有输入的字符串都由小写字母组成，如果找不到距离至少为 k 的重排结果，请返回一个空字符串 ""。

*/
package rearrange_string_k_distance_apart

import (
	"sort"
)

// --- 自己

/*
方法一: 排序分组

时间复杂度：
空间复杂度：
*/
func rearrangeString(s string, k int) string {
	if len(s) <= 1 || k <= 0 {
		return s
	}

	group := make([]string, 0, len(s))
	hash := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		if idx, ok := hash[s[i]]; !ok {
			hash[s[i]] = len(group)
			group = append(group, string(s[i]))
		} else {
			group[idx] += string(s[i])
		}
	}

	// 按重复次数 多到少 排序
	sort.Slice(group, func(i, j int) bool {
		return len(group[i]) > len(group[j])
	})
	// fmt.Println(group)

	if len(group[0]) <= 1 {
		return s
	}

	if k > len(group) {
		return ""
	}
	maxL := len(group[0])
	if len(s) < k*(maxL-1)+1 {
		return ""
	}

	ans := ""
	for len(group[0]) > 0 {
		i := 0
		for ; i < k; i++ {
			if len(group[i]) <= 0 {
				break
			}
			ans += string(group[i][0])
			group[i] = group[i][1:]
		}

		if i < k && len(group[0]) > 0 {
			return ""
		}

		// 依次弹出 k 个字符后 保持排序 这里可以优化
		// 前 k 个元素 和 后 k 个元素 是分别有序的
		// sort.Slice(group, func(i, j int) bool {
		// 	return len(group[i]) > len(group[j])
		// })

		group = sortGroup(group, i)
		// fmt.Println(group, i)
	}
	// fmt.Println(ans)
	return ans
}

func sortGroup(group []string, p int) []string {
	res := make([]string, 0, len(group))

	i := 0
	sp := p
	for i < sp && p < len(group) {
		if len(group[i]) > len(group[p]) {
			res = append(res, group[i])
			i++
		} else {
			res = append(res, group[p])
			p++
		}
	}

	for i < sp {
		res = append(res, group[i])
		i++
	}
	for p < len(group) {
		res = append(res, group[p])
		p++
	}

	return res
}
