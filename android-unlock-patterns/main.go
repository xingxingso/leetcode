/*
https://leetcode-cn.com/problems/android-unlock-patterns/

351. 安卓系统手势解锁

我们都知道安卓有个手势解锁的界面，是一个 3 x 3 的点所绘制出来的网格。用户可以设置一个 “解锁模式” ，通过连接特定序列中的点，
形成一系列彼此连接的线段，每个线段的端点都是序列中两个连续的点。如果满足以下两个条件，则 k 点序列是有效的解锁模式：
	- 解锁模式中的所有点 互不相同 。
	- 假如模式中两个连续点的线段需要经过其他点，那么要经过的点必须事先出现在序列中（已经经过），不能跨过任何还未被经过的点。

以下是一些有效和无效解锁模式的示例：
	无效手势：[4,1,3,6] ，连接点 1 和点 3 时经过了未被连接过的 2 号点。
	无效手势：[4,1,9,2] ，连接点 1 和点 9 时经过了未被连接过的 5 号点。
	有效手势：[2,4,1,3,6] ，连接点 1 和点 3 是有效的，因为虽然它经过了点 2 ，但是点 2 在该手势中之前已经被连过了。
	有效手势：[6,5,4,1,9,2] ，连接点 1 和点 9 是有效的，因为虽然它经过了按键 5 ，但是点 5 在该手势中之前已经被连过了。

给你两个整数，分别为 m 和 n ，那么请你统计一下有多少种 不同且有效的解锁模式 ，是 至少 需要经过 m 个点，但是 不超过 n 个点的。
两个解锁模式 不同 需满足：经过的点不同或者经过点的顺序不同。

提示：
	1 <= m, n <= 9

*/
package android_unlock_patterns

// --- 官方

/*
方法一: 回溯

时间复杂度：O(n!)，其中 n 是最大手势长度。
空间复杂度：O(n)，其中 n 是最大手势长度。
*/
func numberOfPatterns(m int, n int) int {
	res := 0

	used := make([]bool, 9)

	for i := m; i <= n; i++ {
		res += calcPatterns(-1, i, &used)
		// 下面应该可以去除
		for j := 0; j < 9; j++ {
			used[j] = false
		}
	}

	return res
}

func calcPatterns(last, len int, used *[]bool) int {
	if len == 0 {
		return 1
	}
	sum := 0
	for i := 0; i < 9; i++ {
		if isValid(i, last, used) {
			(*used)[i] = true
			sum += calcPatterns(i, len-1, used)
			(*used)[i] = false
		}
	}
	return sum
}

func isValid(index, last int, used *[]bool) bool {
	if (*used)[index] {
		return false
	}
	// first digit of the pattern
	if last == -1 {
		return true
	}
	// knight moves or adjacent cells (in a row or in a column)
	if (index+last)%2 == 1 {
		return true
	}
	// indexes are at both end of the diagonals for example 0,0, and 8,8
	mid := (index + last) / 2
	if mid == 4 {
		return (*used)[mid]
	}
	// adjacent cells on diagonal  - for example 0,0 and 1,0 or 2,0 and //1,1
	if index%3 != last%3 && index/3 != last/3 {
		// if ((index % 3) != (last % 3)) && ((index / 3) != (last / 3)) {
		return true
	}
	// all other cells which are not adjacent
	return (*used)[mid]
}
