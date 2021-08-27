/*
Package perfect_rectangle
https://leetcode-cn.com/problems/perfect-rectangle/

391. 完美矩形

我们有 N 个与坐标轴对齐的矩形, 其中 N > 0, 判断它们是否能精确地覆盖一个矩形区域。

每个矩形用左下角的点和右上角的点的坐标来表示。例如， 一个单位正方形可以表示为 [1,1,2,2]。 ( 左下角的点的坐标为 (1, 1) 以及右上角的点的坐标为 (2, 2) )。

*/
package perfect_rectangle

import (
	"fmt"
	"math"
)

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247488941&idx=1&sn=eda94ebcd26f01ee017419d1ac31d689&scene=21#wechat_redirect

/*
方法一：

时间复杂度：
空间复杂度：
*/
func isRectangleCoverP1(rectangles [][]int) bool {
	type point struct {
		x int
		y int
	}

	gx1, gy1 := math.MaxInt64, math.MaxInt64
	gx2, gy2 := math.MinInt64, math.MinInt64

	actualArea := 0
	// 哈希集合，记录最终图形的顶点
	points := make(map[string]bool, 0)
	for _, v := range rectangles {
		// 计算完美矩形的理论顶点坐标
		gx1, gy1 = min(gx1, v[0]), min(gy1, v[1])
		gx2, gy2 = max(gx2, v[2]), max(gy2, v[3])
		// 累加小矩形的面积
		actualArea += (v[2] - v[0]) * (v[3] - v[1])
		// 记录最终形成的图形中的顶点
		p1, p2 := point{v[0], v[1]}, point{v[0], v[3]}
		p3, p4 := point{v[2], v[1]}, point{v[2], v[3]}
		for _, p := range []point{p1, p2, p3, p4} {
			k := genKey(p.x, p.y)
			if points[k] {
				delete(points, k)
			} else {
				points[k] = true
			}
		}
	}
	// 判断面积是否相同
	expectedArea := (gx2 - gx1) * (gy2 - gy1)
	if actualArea != expectedArea {
		return false
	}
	// 判断最终留下的顶点个数是否为 4
	if len(points) != 4 {
		return false
	}
	// 判断留下的 4 个顶点是否是完美矩形的顶点
	if !points[genKey(gx1, gy1)] {
		return false
	}
	if !points[genKey(gx1, gy2)] {
		return false
	}
	if !points[genKey(gx2, gy1)] {
		return false
	}
	if !points[genKey(gx2, gy2)] {
		return false
	}
	// 面积和顶点都对应，说明矩形符合题意
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func genKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}
