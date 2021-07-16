/*
Package preimage_size_of_factorial_zeroes_function
https://leetcode-cn.com/problems/preimage-size-of-factorial-zeroes-function/

793. 阶乘函数后 K 个零

 f(x) 是 x! 末尾是 0 的数量。（回想一下 x! = 1 * 2 * 3 * ... * x，且 0! = 1 ）
例如， f(3) = 0 ，因为 3! = 6 的末尾没有 0 ；而 f(11) = 2 ，
因为 11!= 39916800 末端有 2 个 0 。给定 K，找出多少个非负整数 x ，能满足 f(x) = K 。

提示：
	K 是范围在 [0, 10^9] 的整数。

*/
package preimage_size_of_factorial_zeroes_function

import "math"

// --- 他人
//https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247486776&idx=1&sn=2bb7fb85dabe47711bf61fa4de0b92b2&scene=21#wechat_redirect

/*
方法一:

时间复杂度：
空间复杂度：
*/
func preimageSizeFZF(k int) int {
	// 左边界和右边界之差 + 1 就是答案
	return rightBound(k) - leftBound(k) + 1
}

// 满足 trailingZeroes(n) == k 的左边界
func leftBound(k int) int {
	l, r := 0, math.MaxInt64
	for l < r {
		mid := l + (r-l)/2
		if trailingZeroes(mid) < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// 满足 trailingZeroes(n) == k 的右边界
func rightBound(k int) int {
	l, r := 0, math.MaxInt64
	for l < r {
		mid := l + (r-l)/2
		if trailingZeroes(mid) > k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l - 1
}

// copy form factorial_trailing_zeroes
func trailingZeroes(n int) int {
	res := 0
	for n >= 5 {
		res += n / 5
		n /= 5
	}

	return res
}
