/*
https://leetcode-cn.com/problems/one-edit-distance/

161. 相隔为 1 的编辑距离

给定两个字符串 s 和 t，判断他们的编辑距离是否为 1。

注意：
满足编辑距离等于 1 有三种可能的情形：
往 s 中插入一个字符得到 t
从 s 中删除一个字符得到 t
在 s 中替换一个字符得到 t
*/
package one_edit_distance

// --- 自己

/*
方法一: 参考官方题解后自己编写

时间复杂度：O(N)。最坏情况下，当字符串的长度足够接近，即 abs(N(s) - N(t)) <= 1，其中 N 是求一个字符串的长度的函数。
	最好情况下， abs(N(s)-N(t))>1 时，时间复杂度为 O(1)。
空间复杂度：O(1)。
*/
func isOneEditDistance(s string, t string) bool {
	// 这样只用考虑 插入 替换 两种情况
	// 删除的不用考虑了, 同时简化了逻辑
	if len(s) > len(t) {
		return isOneEditDistance(t, s)
	}

	// 差值大于2的肯定不满足
	if len(t)-len(s) > 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] { // 存在不同
			if len(s) == len(t) {
				return s[i+1:] == t[i+1:]
			}
			return s[i:] == t[i+1:]
		}
	}

	// 不存在不同
	return len(s)+1 == len(t) // 两个完全相同也不满足
}
