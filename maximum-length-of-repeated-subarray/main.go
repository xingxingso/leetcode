/*
Package maximum_length_of_repeated_subarray
https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/

718. 最长重复子数组

给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。

提示：
	1 <= nums1.length, nums2.length <= 1000
	0 <= nums1[i], nums2[i] <= 100

*/
package maximum_length_of_repeated_subarray

// --- 自己

/*
方法一: 动态规划
	状态可优化

时间复杂度：
空间复杂度：
*/
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := 0; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	ans := 0
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			}
		}
	}

	return ans
}

// --- 官方

/*
方法一: 滑动窗口
	我们可以枚举 A 和 B 所有的对齐方式。对齐的方式有两类：
	第一类为 A 不变，B 的首元素与 A 中的某个元素对齐；
	第二类为 B 不变，A 的首元素与 B 中的某个元素对齐。
	对于每一种对齐方式，我们计算它们相对位置相同的重复子数组即可。

时间复杂度：
空间复杂度：
*/
func findLengthO1(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	ans := 0
	for i := 0; i < m; i++ {
		l := min(m-i, n)
		length := maxLength(nums1, i, nums2, 0, l)
		ans = max(ans, length)
	}
	for i := 0; i < n; i++ { // 这种情况也必须考虑,
		l := min(n-i, m)
		length := maxLength(nums1, 0, nums2, i, l)
		ans = max(ans, length)
	}

	return ans
}

func maxLength(nums1 []int, pos1 int, nums2 []int, pos2 int, l int) int {
	k := 0
	ret := 0
	for i := 0; i < l; i++ {
		if nums1[i+pos1] == nums2[i+pos2] {
			k++
			ret = max(ret, k)
		} else {
			k = 0
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
