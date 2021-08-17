/*
Package different_ways_to_add_parentheses
https://leetcode-cn.com/problems/different-ways-to-add-parentheses/

241. 为运算表达式设计优先级

给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。
你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 * 。

*/
package different_ways_to_add_parentheses

import (
	"strconv"
)

// --- 自己

/*
方法一: 分治


时间复杂度：
空间复杂度：
*/
func diffWaysToCompute(expression string) []int {
	express := trans(expression)
	var dfs func(express []string) []int
	dfs = func(express []string) []int {
		if len(express) == 1 {
			n, _ := strconv.Atoi(express[0])
			return []int{n}
		}

		ans := make([]int, 0)
		for i := 1; i < len(express); i += 2 {
			a := dfs(express[:i])
			b := dfs(express[i+1:])
			temp := calculate(a, b, express[i])
			ans = append(ans, temp...)
		}
		return ans
	}

	return dfs(express)
}

func trans(expression string) []string {
	ans := make([]string, 0)
	i := 0
	for i < len(expression) {
		j := i
		for i < len(expression) && expression[i] >= '0' && expression[i] <= '9' {
			i++
		}
		ans = append(ans, expression[j:i])
		if i < len(expression) {
			if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
				ans = append(ans, string(expression[i]))
			}
			i++
		}
	}
	return ans
}

func calculate(num1, num2 []int, expr string) []int {
	ans := make([]int, len(num1)*len(num2))
	k := 0
	switch expr {
	case "+":
		for i := 0; i < len(num1); i++ {
			for j := 0; j < len(num2); j++ {
				ans[k] = num1[i] + num2[j]
				k++
			}
		}
	case "-":
		for i := 0; i < len(num1); i++ {
			for j := 0; j < len(num2); j++ {
				ans[k] = num1[i] - num2[j]
				k++
			}
		}
	case "*":
		for i := 0; i < len(num1); i++ {
			for j := 0; j < len(num2); j++ {
				ans[k] = num1[i] * num2[j]
				k++
			}
		}
	}
	return ans
}
