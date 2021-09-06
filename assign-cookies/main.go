/*
Package assign_cookies
https://leetcode-cn.com/problems/assign-cookies/

455. 分发饼干

假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；
并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，
我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。
你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

提示：
	1 <= g.length <= 3 * 10^4
	0 <= s.length <= 3 * 10^4
	1 <= g[i], s[j] <= 2^31 - 1

*/
package assign_cookies

import (
	"sort"
)

// --- 自己

/*
方法一: 贪心

时间复杂度：O(mlogm+nlogn)，其中 m 和 n 分别是数组 g 和 s 的长度。
空间复杂度：O(logm+logn)，其中 m 和 n 分别是数组 g 和 s 的长度。
*/
func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	count := 0
	for i := 0; count < len(g) && i < len(s); i++ {
		if s[i] >= g[count] {
			count++
		}
	}
	return count
}
