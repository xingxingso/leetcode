/*
Package sum_of_two_integers
https://leetcode-cn.com/problems/sum-of-two-integers/

371. 两整数之和

给你两个整数 a 和 b ，不使用 运算符 + 和 - ，计算并返回两整数之和。

提示：
	-1000 <= a, b <= 1000

*/
package sum_of_two_integers

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func getSum(a int, b int) int {
	half := func(x, y, r int) (int, int) {
		n := x ^ y ^ r
		r = (x|y)&r | (x & y)
		return n, r
	}

	r, k := 0, 0
	ans := 0
	for k < 64 {
		x := a & 1
		y := b & 1
		n := 0
		n, r = half(x, y, r)
		ans = ans | n<<k
		a >>= 1
		b >>= 1
		k++
	}

	return ans
}

// --- 官方

/*
方法一:

时间复杂度：
空间复杂度：
*/
func getSumO1(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}
