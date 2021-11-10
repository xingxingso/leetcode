/*
Package sum_of_square_numbers
https://leetcode-cn.com/problems/sum-of-square-numbers/

633. 平方数之和

给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a^2 + b^2 = c 。

提示：
	0 <= c <= 2^31 - 1

*/
package sum_of_square_numbers

import (
	"math"
)

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func judgeSquareSum(c int) bool {
	k := int(math.Sqrt(float64(c)))
	for i := 0; i <= k; i++ {
		if isSqrt(c - i*i) {
			return true
		}
	}
	return false
}

func isSqrt(n int) bool {
	k := int(math.Sqrt(float64(n)))
	return k*k == n
}

// --- 官方

/*
方法二: 双指针

时间复杂度：O(sqrt{c})。最坏情况下 a 和 b 一共枚举了 0 到 sqrt{c} 里的所有整数。
空间复杂度：O(1)。
*/
func judgeSquareSumO1(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}
