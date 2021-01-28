/*
https://leetcode-cn.com/problems/factor-combinations/

254. 因子的组合

整数可以被看作是其因子的乘积。
例如：
	8 = 2 x 2 x 2;
  	  = 2 x 4.
请实现一个函数，该函数接收一个整数 n 并返回该整数所有的因子组合。

注意：
	1. 你可以假定 n 为永远为正数。
	2. 因子必须大于 1 并且小于 n。

*/
package factor_combinations

// --- 自己

/*
方法一:
	参考
	https://leetcode-cn.com/problems/factor-combinations/solution/c-di-gui-jie-fa-by-da-li-wang-2/

时间复杂度：
空间复杂度：
*/
func getFactors(n int) [][]int {
	// l 表示起始值， 保证有序性从而去重
	return dfs(2, n)
}

func dfs(l, n int) [][]int {
	var res [][]int
	for i := l; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, []int{n / i, i})
			// fmt.Println(res)
			for _, v := range dfs(i, n/i) {
				v := append(v, i)
				res = append(res, v)
			}
		}
	}
	return res
}
