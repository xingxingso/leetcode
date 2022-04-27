/*
Package combinations
https://leetcode-cn.com/problems/combinations/

77. 组合

给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

*/
package combinations

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	item := make([]int, 0)
	var backTack func(idx int)
	backTack = func(idx int) {
		if k == len(item) {
			//temp := make([]int, k)
			//copy(temp, item)
			//ans = append(ans, temp)
			ans = append(ans, append([]int{}, item...))
			return
		}
		for i := idx; i <= n; i++ {
			item = append(item, i)
			backTack(i + 1)
			item = item[:len(item)-1]
		}
		return
	}

	backTack(1)
	return ans
}
