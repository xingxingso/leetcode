/*
Package climbing_stairs
https://leetcode-cn.com/problems/climbing-stairs/

70. 爬楼梯

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

*/
package climbing_stairs

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	pre, cur := 1, 1
	res := 0
	for i := 2; i <= n; i++ {
		res = pre + cur
		pre, cur = cur, res
	}
	return res
}
