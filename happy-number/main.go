/*
Package happy_number
https://leetcode-cn.com/problems/happy-number/

202. 快乐数

编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：
	对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
	然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
	如果 可以变为 1，那么这个数就是快乐数。

如果 n 是快乐数就返回 true ；不是，则返回 false 。

提示：
	1 <= n <= 2^31 - 1

*/
package happy_number

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func isHappy(n int) bool {
	memo := make(map[int]bool)

	var dfs func(n int) bool
	dfs = func(n int) bool {
		if n == 1 {
			return true
		}
		if memo[n] {
			return false
		}
		memo[n] = true

		m := 0
		for n > 0 {
			a := n % 10
			m += a * a
			n = n / 10
		}

		return dfs(m)
	}

	return dfs(n)
}
