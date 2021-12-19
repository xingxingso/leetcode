/*
Package beautiful_array
https://leetcode-cn.com/problems/beautiful-array/

932. 漂亮数组

对于某些固定的 N，如果数组 A 是整数 1, 2, ..., N 组成的排列，使得：

对于每个 i < j，都不存在 k 满足 i < k < j 使得 A[k] * 2 = A[i] + A[j]。

那么数组 A 是漂亮数组。

给定 N，返回任意漂亮数组 A（保证存在一个）。

提示：
	1 <= N <= 1000

*/
package beautiful_array

// --- 官方

/*
方法一: 分治

时间复杂度：O(NlogN)，代码中的函数 partition 执行的次数为 O(logN)，每次执行的时间复杂度为 O(N)。
空间复杂度：O(NlogN)。
*/
func beautifulArray(n int) []int {
	memo := make(map[int][]int)
	var partition func(n int) []int
	partition = func(n int) []int {
		if memo[n] != nil {
			return memo[n]
		}
		ans := make([]int, n)
		if n == 1 {
			ans[0] = 1
		} else {
			t := 0
			for _, x := range partition((n + 1) / 2) {
				ans[t] = 2*x - 1
				t++
			}
			for _, x := range partition(n / 2) {
				ans[t] = 2 * x
				t++
			}
		}

		memo[n] = ans
		return ans
	}

	return partition(n)
}
