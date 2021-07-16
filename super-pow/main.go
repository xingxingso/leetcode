/*
Package super_pow
https://leetcode-cn.com/problems/super-pow/

372. 超级次方

你的任务是计算 a^b 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。

提示：
	1 <= a <= 2^31 - 1
	1 <= b.length <= 2000
	0 <= b[i] <= 9
	b 不含前导 0

*/
package super_pow

// --- 他人
//https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485035&idx=1&sn=c03c9056f412bf590480156e4357b433&scene=21#wechat_redirect

/*
方法一:
	a^m×a^n=a^(m+n)
	(a^m)^n=a^(m*n)
	(a*b)%k = (a%k)*(b%k)%k

时间复杂度：
空间复杂度：
*/
func superPowP1(a int, b []int) int {
	divisor := 1337
	var myPow func(a, pow int) int
	myPow = func(a, pow int) int {
		if pow == 0 {
			return 1
		}
		a %= divisor
		if pow%2 == 1 {
			return (a * myPow(a, pow-1)) % divisor
		} else {
			sub := myPow(a, pow/2)
			return sub * sub % divisor
		}
	}

	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]
	part1 := myPow(a, last)
	part2 := myPow(superPowP1(a, b), 10)
	return (part1 % divisor) * (part2 % divisor) % divisor
}
