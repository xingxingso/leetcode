/*
Package jump_game
https://leetcode-cn.com/problems/jump-game/

55. 跳跃游戏

给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。

提示：
	1 <= nums.length <= 3 * 10^4
	0 <= nums[i] <= 10^5

*/
package jump_game

// --- 自己

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func canJumpS1(nums []int) bool {
	// dp[i][j] 表示从i位置跳跃 是否可以到达终点
	// 目标值 dp[0]
	dp := make([]bool, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		// 已经位于终点
		if i == len(nums)-1 {
			dp[i] = true
			continue
		}
		// 下一次就可以到达终点
		if i+nums[i] >= len(nums)-1 {
			dp[i] = true
			continue
		}
		// 寻找下一次可以到达终点的点，如果找不到说明该点不可达
		for j := 1; j <= nums[i]; j++ {
			if dp[i+j] == true {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}

/*
方法一: 动态规划

时间复杂度：
空间复杂度：
*/
func canJumpS2(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			dp[i] = dp[j] && nums[j] >= i-j
			if dp[i] {
				break
			}
		}
	}
	// fmt.Println(dp)
	return dp[len(nums)-1]
}

/*
方法二: 贪心
	每次跳总选择`利益最大化的`, 即 跳跃步数 + 到达位置的可选步数 值最大

时间复杂度：
空间复杂度：
*/
func canJumpS3(nums []int) bool {
	for i := 0; i < len(nums); {
		if nums[i]+i >= len(nums)-1 {
			return true
		}
		cho := 0
		step := 0
		for j := 1; j <= nums[i]; j++ {
			if j+nums[j+i] > cho {
				cho = j + nums[j+i]
				step = j
			}
		}
		if step == 0 {
			return false
		}
		i += step
	}
	return true
}

/*
方法二: 倒推

时间复杂度：
空间复杂度：
*/
func canJumpS4(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	step := 1
	for i := len(nums) - 2; i > 0; {
		if nums[i] >= step {
			i--
			step = 1
		} else {
			i--
			step++
		}
		// fmt.Println(i, step)
	}
	return nums[0] >= step
}

// --- 官方

/*
方法一: 贪心

时间复杂度：O(n)，其中 n 为数组的大小。只需要访问 nums 数组一遍，共 n 个位置。
空间复杂度：O(1)，不需要额外的空间开销。
*/
func canJumpO1(nums []int) bool {
	rightmost := 0
	for i := 0; i < len(nums); i++ {
		if i <= rightmost {
			rightmost = max(rightmost, i+nums[i])
			if rightmost >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}

func canJumpO2(nums []int) bool {
	rightmost := 0
	for i := 0; i < len(nums)-1; i++ { //  i < len(nums)-1
		// 不断计算能跳到的最远距离
		rightmost = max(rightmost, i+nums[i])
		// 可能碰到了 0，卡住跳不动了
		if rightmost <= i {
			return false
		}
	}
	return true
	// return rightmost >= len(nums)-1
}

func canJumpO3(nums []int) bool {
	for rightmost, i := 0, 0; i < len(nums); i++ {
		rightmost = max(rightmost, i+nums[i])
		if rightmost >= len(nums)-1 {
			return true
		} else if rightmost <= i {
			return false
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
