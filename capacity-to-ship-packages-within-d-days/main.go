/*
Package capacity_to_ship_packages_within_d_days
https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/

1011. 在 D 天内送达包裹的能力

传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。
传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。
我们装载的重量不会超过船的最大运载重量。
返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。

提示：
	1 <= D <= weights.length <= 5 * 10^4
	1 <= weights[i] <= 500

*/
package capacity_to_ship_packages_within_d_days

import (
	"sort"
)

// --- 自己

/*
方法一: 二分查找

时间复杂度：
空间复杂度：
*/
func shipWithinDays(weights []int, D int) int {
	return sort.Search(500*5*10000, func(i int) bool {
		d := 0
		for j := 0; j < len(weights); {
			w := 0
			for ; j < len(weights) && w+weights[j] <= i; j++ {
				w += weights[j]
			}
			if w == 0 {
				return false
			}
			d++
		}
		return d <= D
	})
}

// --- 官方

/*
方法一: 二分查找转化为判定问题

时间复杂度：O(nlog(Σw))，其中 n 是数组 weights 的长度，Σw 是数组 weights 中元素的和。
	二分查找需要执行的次数为 O(log(Σw))，每一步中需要对数组 weights 进行依次遍历，时间为 O(n)，相乘即可得到总时间复杂度。
空间复杂度：O(1)。
*/
func shipWithinDaysO1(weights []int, D int) int {
	// 确定二分查找左右边界
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w // left 为其中单个最大值
		}
		right += w // right 为一次性运输最大值，即总和
	}
	return left + sort.Search(right-left, func(x int) bool {
		x += left
		day := 1 // 需要运送的天数
		sum := 0 // 当前这一天已经运送的包裹重量之和
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= D
	})
}
