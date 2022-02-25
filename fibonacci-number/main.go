/*
Package fibonacci_number
https://leetcode-cn.com/problems/fibonacci-number/

509. 斐波那契数

斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
	F(0) = 0，F(1) = 1
	F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给定 n ，请计算 F(n) 。

提示：
	0 <= n <= 30

*/
package fibonacci_number

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func fib(n int) int {
	if n == 0 {
		return 0
	}
	pre, cur := 0, 1
	for ; n > 1; n-- {
		temp := cur
		cur = pre + cur
		pre = temp
	}

	return cur
}
