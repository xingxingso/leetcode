/*
Package daily_temperatures
https://leetcode-cn.com/problems/daily-temperatures/

739. 每日温度

请根据每日 气温 列表 temperatures ，请计算在每一天需要等几天才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。

提示：
	1 <= temperatures.length <= 10^5
	30 <= temperatures[i] <= 100

*/
package daily_temperatures

// --- 自己

/*
方法一: 单调栈

时间复杂度：
空间复杂度：
*/
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	ans := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		var top int
		for len(stack) > 0 {
			top = stack[len(stack)-1]
			if temperatures[top] > temperatures[i] {
				break
			}
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			ans[i] = top - i
		}
		stack = append(stack, i)
	}

	return ans
}
