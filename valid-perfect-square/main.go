/*
Package valid_perfect_square
https://leetcode-cn.com/problems/valid-perfect-square/

367. 有效的完全平方数

给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

进阶：不要 使用任何内置的库函数，如 sqrt 。

提示：
	1 <= num <= 2^31 - 1

*/
package valid_perfect_square

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func isPerfectSquare(num int) bool {
	for i := 1; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
