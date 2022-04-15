/*
Package minimum_time_to_make_rope_colorful
https://leetcode-cn.com/problems/minimum-time-to-make-rope-colorful/

1578. 使绳子变成彩色的最短时间

Alice 把 n 个气球排列在一根绳子上。给你一个下标从 0 开始的字符串 colors ，其中 colors[i] 是第 i 个气球的颜色。

Alice 想要把绳子装扮成 彩色 ，且她不希望两个连续的气球涂着相同的颜色，所以她喊来 Bob 帮忙。Bob 可以从绳子上移除一些气球使绳子变成 彩色 。
给你一个下标从 0 开始的整数数组 neededTime ，其中 neededTime[i] 是 Bob 从绳子上移除第 i 个气球需要的时间（以秒为单位）。

返回 Bob 使绳子变成 彩色 需要的 最少时间 。

提示：
	n == colors.length == neededTime.length
	1 <= n <= 10^5
	1 <= neededTime[i] <= 10^4
	colors 仅由小写英文字母组成

*/
package minimum_time_to_make_rope_colorful

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func minCost(colors string, neededTime []int) int {
	ans := 0
	for i := 0; i < len(colors)-1; {
		next := i + 1
		for next < len(colors) && colors[next] == colors[i] {
			next++
		}
		maxCost := neededTime[i]
		cost := 0
		for i < next {
			cost += neededTime[i]
			if neededTime[i] > maxCost {
				maxCost = neededTime[i]
			}
			i++
		}
		cost -= maxCost
		ans += cost
	}

	return ans
}
