/*
https://leetcode-cn.com/problems/jump-game-ii/

45. 跳跃游戏 II

给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。

说明:
	假设你总是可以到达数组的最后一个位置。

*/
package jump_game_ii

// --- 自己

/*
方法一: 贪心
	每次跳总选择`利益最大化的`, 即 跳跃步数 + 到达位置的可选步数 值最大

时间复杂度：
空间复杂度：
*/
func jump(nums []int) int {
	count := 0
	for i := 0; i < len(nums)-1; {
		if nums[i]+i >= len(nums)-1 {
			return count + 1
		}
		cho := 0  // 下一次跳跃后的最大可能跳跃位置
		step := 0 // 选择跳多少步
		for j := 1; j <= nums[i]; j++ {
			if j+nums[j+i] > cho {
				cho = j + nums[j+i]
				step = j
			}
		}
		if step == 0 { // 没有办法前进了
			return -1
		}
		i += step
		count++
	}
	return count
}

// --- 官方

/*
方法一: 贪心

时间复杂度：
空间复杂度：
*/
func jumpO1(nums []int) int {
	steps, end, maxPosition := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
