/*
Package best_sightseeing_pair
https://leetcode-cn.com/problems/best-sightseeing-pair/

1014. 最佳观光组合

给你一个正整数数组 values，其中 values[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的 距离 为 j - i。

一对景点（i < j）组成的观光组合的得分为 values[i] + values[j] + i - j ，也就是景点的评分之和 减去 它们两者之间的距离。

返回一对观光景点能取得的最高分。

提示：
	2 <= values.length <= 5 * 10^4
	1 <= values[i] <= 1000

*/
package best_sightseeing_pair

// --- 自己

/*
方法一: 暴力解法
	// 超出时间限制

时间复杂度：
空间复杂度：
*/
func maxScoreSightseeingPairS1(values []int) int {
	ans := 0
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			n := values[j] + values[i] + i - j
			if n > ans {
				ans = n
			}
		}
	}
	return ans
}

// --- 官方

/*
方法一: 遍历

时间复杂度：
空间复杂度：
*/
func maxScoreSightseeingPairO1(values []int) int {
	ans := 0
	lastMax := values[0] + 0
	for i := 1; i < len(values); i++ {
		tmp := lastMax + values[i] - i
		if tmp > ans {
			ans = tmp
		}
		if values[i]+i > lastMax {
			lastMax = values[i] + i
		}
	}
	return ans
}
