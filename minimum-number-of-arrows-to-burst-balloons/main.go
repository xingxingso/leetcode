/*
Package minimum_number_of_arrows_to_burst_balloons
https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/

452. 用最少数量的箭引爆气球

在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。由于它是水平的，所以纵坐标并不重要，
因此只要知道开始和结束的横坐标就足够了。开始坐标总是小于结束坐标。
一支弓箭可以沿着 x 轴从不同点完全垂直地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 x_start，x_end，
且满足  x_start ≤ x ≤ x_end，则该气球会被引爆。可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。
我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。
给你一个数组 points ，其中 points [i] = [x_start,x_end] ，返回引爆所有气球所必须射出的最小弓箭数。

提示：
	0 <= points.length <= 10^4
	points[i].length == 2
	-2^31 <= x_start < x_end <= 2^31 - 1

*/
package minimum_number_of_arrows_to_burst_balloons

import "sort"

// --- 自己

/*
方法一: 贪心

时间复杂度：
空间复杂度：
*/
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	ans, end := 1, points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			ans++
			end = points[i][1]
		}
	}
	return ans
}

func findMinArrowShotsS2(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1] || (points[i][1] == points[j][1] && points[i][0] < points[j][0])
	})
	// fmt.Println(points)
	count := 0
	for i := 0; i < len(points); i++ {
		last := points[i][1]
		for ; i+1 < len(points) && points[i+1][0] <= last; i++ {
		}
		count++
		// fmt.Println(last)
	}

	return count
}
