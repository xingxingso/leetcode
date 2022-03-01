/*
Package max_consecutive_ones_iii
https://leetcode-cn.com/problems/max-consecutive-ones-iii/

1004. 最大连续1的个数 III

给定一个二进制数组 nums 和一个整数 k ，如果可以翻转最多k 个 0 ，则返回 数组中连续 1 的最大个数 。

提示：
	1 <= nums.length <= 10^5
	nums[i] 不是 0 就是 1
	0 <= k <= nums.length

*/
package max_consecutive_ones_iii

// --- 自己

/*
方法一: 滑动窗口

时间复杂度：
空间复杂度：
*/
func longestOnes(nums []int, k int) int {
	ans := 0
	window := make([]int, 0)
	for l, r := -1, 0; r < len(nums); r++ {
		if nums[r] == 0 {
			window = append(window, r)
		}
		if len(window) > k {
			l = window[0]
			window = window[1:]
		}
		ans = max(ans, r-l)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// --- 他人

/*
方法一: 滑动窗口
	是的，这个写法维护的是一个只能单调变长的窗口。
	这种窗口经常出现在寻求”最大窗口“的问题中：因为要求的是”最大“，所以我们没有必要缩短窗口，于是代码就少了缩短窗口的部分；
	从另一个角度讲，本题里的K是消耗品，一旦透支，窗口就不能再增长了（也意味着如果K == 0还是有可能增长的）。
	所以K所代表的”资源“，通常是滑窗维护逻辑的核心，能这么写有两个先决条件：
		- 固定一个左端点，K随窗口增大是单调变化的。据此我们可以推知长度为n的窗口如若已经”透支“（K < 0）了，那么长度大于n的也一定不符合条件；
		- K的变化与数组元素有简单的算术关系。向窗口纳入（A[r++]）或移除（A[l++]）一个数组元素，可以在O(1)内更新K。
	虽说有条件，但仔细排查会发现许多滑窗问题都可以满足。

时间复杂度：
空间复杂度：
*/
func longestOnesP1(nums []int, k int) int {
	l, r := 0, 0
	for ; r < len(nums); r++ {
		if nums[r] == 0 {
			k--
		}
		if k < 0 {
			if nums[l] == 0 {
				k++
			}
			l++
		}
	}
	return r - l

	//int l = 0, r = 0;
	//while (r < A.length) {
	//	if (A[r++] == 0) K--;
	//if (K < 0 && A[l++] == 0) K++;
	//}
	//return r - l;
}
