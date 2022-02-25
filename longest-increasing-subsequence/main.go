/*
Package longest_increasing_subsequence
https://leetcode-cn.com/problems/longest-increasing-subsequence/

300. 最长递增子序列

给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

提示：
	1 <= nums.length <= 2500
	-10^4 <= nums[i] <= 10^4

进阶：
	你可以设计时间复杂度为 O(n^2) 的解决方案吗？
	你能将算法的时间复杂度降低到 O(nlog(n)) 吗?

*/
package longest_increasing_subsequence

// --- 自己

/*
方法一: 动态规划
	参考别人题解

时间复杂度：O(n^2)，其中 n 为数组 nums 的长度。动态规划的状态数为 n，计算状态 dp[i] 时，需要 O(n) 的时间遍历 dp[0…i−1] 的所有状态，
	所以总时间复杂度为 O(n^2)。
空间复杂度：O(n)，需要额外使用长度为 n 的 dp 数组。
*/
func lengthOfLIS(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			ans = max(ans, dp[i])
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func lengthOfLISS2(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	ans := 1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] && dp[i] >= dp[j] {
				dp[j] = dp[i] + 1
				if dp[j] > ans {
					ans = dp[j]
				}
			}
		}
	}
	//fmt.Println(dp)
	return ans
}

/*
方法二: 贪心 + 二分查找

时间复杂度：O(nlogn)。数组 nums 的长度为 n，我们依次用数组中的元素去更新 d 数组，而更新 d 数组时需要进行 O(logn) 的二分搜索，
	所以总时间复杂度为 O(nlogn)。
空间复杂度：O(n)，需要额外使用长度为 n 的 d 数组。
*/
func lengthOfLISS3(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	d := make([]int, len(nums))
	d[0] = nums[0]
	lens := 1
	for i := 1; i < len(nums); i++ {
		l, r := 0, lens-1
		// 找到第一个大于等于他的数所在位置,即左侧搜索
		for l <= r {
			mid := l + (r-l)>>1
			if nums[i] > d[mid] {
				l = mid + 1
			} else if nums[i] < d[mid] {
				r = mid - 1
			} else {
				r = mid - 1
			}
		}
		d[l] = nums[i]
		if l >= lens { // 表示在新增了一个大数
			lens++
		}
	}
	return lens
}
