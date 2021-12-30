/*
Package minimum_ascii_delete_sum_for_two_strings
https://leetcode-cn.com/problems/minimum-ascii-delete-sum-for-two-strings/

712. 两个字符串的最小ASCII删除和

给定两个字符串s1, s2，找到使两个字符串相等所需删除字符的ASCII值的最小和。

注意:
	0 < s1.length, s2.length <= 1000。
	所有字符串中的字符ASCII值在[97, 122]之间。

*/
package minimum_ascii_delete_sum_for_two_strings

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func minimumDeleteSum(s1 string, s2 string) int {
	sum1, sum2 := 0, 0
	dp := make([][]int, len(s1)+1) // 取s1前i个字符与s2前j个字符的最大ascii子序列
	d := make([]int, len(s2)+1)
	for i := 0; i <= len(s1); i++ {
		if i != len(s1) {
			sum1 += int(s1[i])
		}
		dp[i] = make([]int, len(s2)+1)
	}
	for i := 0; i < len(s2); i++ {
		sum2 += int(s2[i])
	}
	for i := 1; i <= len(s1); i++ {
		pre := 0
		for j := 1; j <= len(s2); j++ {
			temp := d[j]
			if s1[i-1] == s2[j-1] {
				// dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
				d[j] = pre + int(s1[i-1])
			} else {
				// dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				d[j] = max(d[j], d[j-1])
			}
			pre = temp
		}
	}
	// fmt.Println(dp)

	return sum1 + sum2 - 2*d[len(s2)]
	// return sum1 + sum2 - 2*dp[len(s1)][len(s2)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
