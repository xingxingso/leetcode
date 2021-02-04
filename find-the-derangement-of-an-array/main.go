/*
https://leetcode-cn.com/problems/find-the-derangement-of-an-array/

634. 寻找数组的错位排列

在组合数学中，如果一个排列中所有元素都不在原先的位置上，那么这个排列就被称为错位排列。
给定一个从 1 到 n 升序排列的数组，你可以计算出总共有多少个不同的错位排列吗？
由于答案可能非常大，你只需要将答案对 109+7 取余输出即可。

注释:
	n 的范围是 [1, 106]。

*/
package find_the_derangement_of_an_array

// --- 官方

/*
方法一: 动态规划

时间复杂度：O(N)。
空间复杂度：O(N)。
*/
func findDerangementO1(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = (i - 1) * (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}

/*
方法二: 动态规划 + 空间优化

时间复杂度：O(N)。
空间复杂度：O(1)。
*/
func findDerangementO2(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}
	first, second := 1, 0
	for i := 2; i <= n; i++ {
		temp := second
		second = (i - 1) * (first + second) % 1000000007
		first = temp
	}
	return second
}

/*
方法三: 容斥原理

时间复杂度：O(N)。
空间复杂度：O(1)。
*/
func findDerangementO3(n int) int {
	mul, sum := 1, 0
	M := 1000000007
	for i := n; i >= 0; i-- {
		fl := -1
		if i%2 == 0 {
			fl = 1
		}

		sum = (sum + M + mul*fl) % M
		mul = (mul * i) % M
	}
	return sum
}
