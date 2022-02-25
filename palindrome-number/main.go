/*
Package palindrome_number
https://leetcode-cn.com/problems/palindrome-number/

9. 回文数

给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
	例如，121 是回文，而 123 不是。

提示：
	-2^31 <= x <= 2^31 - 1


进阶：
	你能不将整数转为字符串来解决这个问题吗？

*/
package palindrome_number

// --- 自己

/*
方法一: 反转全部

时间复杂度：
空间复杂度：
*/
func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}

	n, ori := 0, x
	for x > 0 {
		tail := x % 10
		x = x / 10
		n = n*10 + tail
	}

	return n == ori
}

// --- 官方

/*
方法一: 反转一半数字

时间复杂度：O(logn)，对于每次迭代，我们会将输入除以 10，因此时间复杂度为 O(logn)。
空间复杂度：O(1)。我们只需要常数空间存放若干变量。
*/
func isPalindromeO1(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}
