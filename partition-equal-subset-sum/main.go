/*
Package partition_equal_subset_sum
https://leetcode-cn.com/problems/partition-equal-subset-sum/

416. 分割等和子集

给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:
	每个数组中的元素不会超过 100
	数组的大小不会超过 200

*/
package partition_equal_subset_sum

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func canPartition(nums []int) bool {
	s := sum(nums)
	if s%2 != 0 {
		return false
	}
	target := s / 2
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = make([]bool, target+1)
			continue
		}

		dp[i] = make([]bool, target+1)
		for j := 0; j <= target; j++ {
			// if j == 0 { // 1. 1，2 选择一个
			// 	dp[i][j] = true
			// 	continue
			// }

			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else if nums[i] == j { // 2. 1,2 选择一个
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}

	return dp[len(nums)-1][target]
}

func sum(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans += v
	}

	return ans
}

// --- 官方

/*
方法一: 动态规划

时间复杂度：O(n×target)，其中 n 是数组的长度，target 是整个数组的元素和的一半。
	需要计算出所有的状态，每个状态在进行转移时的时间复杂度为 O(1)。
空间复杂度：O(target)，其中 target 是整个数组的元素和的一半。
	空间复杂度取决于 dp 数组，在不进行空间优化的情况下，空间复杂度是 O(n×target)，
	在进行空间优化的情况下，空间复杂度可以降到 O(target)。
*/
func canPartitionO1(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	if max > target {
		return false
	}

	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		v := nums[i]
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[target]
}
