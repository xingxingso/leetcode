/*
Package max_points_on_a_line
https://leetcode-cn.com/problems/max-points-on-a-line/

149. 直线上最多的点数

给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。

提示：
	1 <= points.length <= 300
	points[i].length == 2
	-10^4 <= xi, yi <= 10^4
	points 中的所有点 互不相同

*/
package max_points_on_a_line

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func maxPoints(points [][]int) int {
	ans := 0
	for i := 0; i < len(points); i++ {
		ratio := make(map[float64]int)
		infinite := 0
		x, y := points[i][0], points[i][1]
		for j := i + 1; j < len(points); j++ {
			point := points[j]
			diffX, diffY := float64(point[0]-x), float64(point[1]-y)
			if diffY == 0 {
				infinite++
			} else {
				ratio[diffX/diffY]++
			}
		}
		if infinite >= ans {
			ans = infinite + 1 // 原来那个点要算上
		}
		for _, v := range ratio {
			if v >= ans {
				ans = v + 1 // 原来那个点要算上
			}
		}
	}

	return ans
}
