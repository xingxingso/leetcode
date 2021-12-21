/*
Package heaters
https://leetcode-cn.com/problems/heaters/

475. 供暖器

冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。

在加热器的加热半径范围内的每个房屋都可以获得供暖。

现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。

说明：
	所有供暖器都遵循你的半径标准，加热的半径也一样。

提示：
	1 <= houses.length, heaters.length <= 3 * 10^4
	1 <= houses[i], heaters[i] <= 10^9

*/
package heaters

import (
	"math"
	"sort"
)

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	radius := 0
	for cur, i := 0, 0; i < len(houses); i++ {
		if abs(houses[i], heaters[cur]) <= radius {
			continue
		}

		// 已经是最大的了，只能提高半径
		if cur == len(heaters)-1 {
			radius = abs(houses[i], heaters[cur])
			continue
		}

		// 找到 house[i] 左边(cur-1) 和 右边(cur) 最近的加热器,可能只有一边, 至少有一边
		// 这里可以通过二分查找优化
		for ; cur < len(heaters) && heaters[cur] <= houses[i]; cur++ {
		}
		leftDis, rightDis := math.MaxInt64, math.MaxInt64
		if cur-1 >= 0 {
			leftDis = abs(heaters[cur-1], houses[i])
		}
		if cur < len(heaters) {
			rightDis = abs(heaters[cur], houses[i])
		}
		minDis := rightDis
		if leftDis < rightDis {
			cur-- // 使用左侧加热器
			minDis = leftDis
		}
		if radius < minDis {
			radius = minDis
		}
	}

	return radius
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
