/*
Package sqrtx
https://leetcode-cn.com/problems/sqrtx/

69. x 的平方根

给你一个非负整数 x ，计算并返回 x 的 平方根 。
由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：
	不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

提示：
	0 <= x <= 2^31 - 1

*/
package sqrtx

import "math"

// --- 自己

/*
方法一：

时间复杂度：
空间复杂度：
*/
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	l, r := 1, x
	for l <= r {
		mid := l + (r-l)>>1
		sqrt := x / mid
		if sqrt < mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l - 1
}

// 官方

/*
方法二：牛顿迭代法

时间复杂度：O(logx)，此方法是二次收敛的，相较于二分查找更快。
空间复杂度：O(1)。
*/
func mySqrtO1(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}
