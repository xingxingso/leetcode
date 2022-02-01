/*
Package basic_calculator_ii
https://leetcode-cn.com/problems/basic-calculator-ii/

227. 基本计算器 II

给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。

整数除法仅保留整数部分。

提示：
	1 <= s.length <= 3 * 10^5
	s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
	s 表示一个 有效表达式
	表达式中的所有整数都是非负整数，且在范围 [0, 2^31 - 1] 内
	题目数据保证答案是一个 32-bit 整数

*/
package basic_calculator_ii

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func calculate(s string) int {
	stackOper := make([]byte, 0)
	stackNum := make([]int, 0)
	num := 0
	start, end := 0, len(s)-1
	for s[start] == ' ' {
		start++
	}
	for s[end] == ' ' {
		end--
	}

	for i := start; i <= end; i++ {
		v := s[i]
		if v == ' ' {
			//stackNum = append(stackNum, num)
			//num = 0
			continue
		}
		//fmt.Println(num)

		if isOperator(v) {
			stackNum = append(stackNum, num)
			num = 0
			// 查看上面操作符
			if len(stackOper) > 0 {
				top := stackOper[len(stackOper)-1]
				if top == '*' || top == '/' {
					stackOper = stackOper[:len(stackOper)-1]
					n1, n2 := stackNum[len(stackNum)-2], stackNum[len(stackNum)-1]
					stackNum = stackNum[:len(stackNum)-2]
					if top == '*' {
						stackNum = append(stackNum, n1*n2)
					} else {
						stackNum = append(stackNum, n1/n2)
					}
				}
			}
			stackOper = append(stackOper, v)
		} else {
			num = 10*num + int(v-'0')
		}
	}
	stackNum = append(stackNum, num)

	if len(stackOper) > 0 {
		top := stackOper[len(stackOper)-1]
		if top == '*' || top == '/' {
			stackOper = stackOper[:len(stackOper)-1]
			n1, n2 := stackNum[len(stackNum)-2], stackNum[len(stackNum)-1]
			stackNum = stackNum[:len(stackNum)-2]
			if top == '*' {
				stackNum = append(stackNum, n1*n2)
			} else if top == '/' {
				stackNum = append(stackNum, n1/n2)
			}
		}
	}

	//fmt.Println(stackNum, stackOper)
	for i := 1; i < len(stackNum); i++ {
		first := stackOper[0]
		stackOper = stackOper[1:]
		if first == '+' {
			stackNum[i] = stackNum[i-1] + stackNum[i]
		} else {
			stackNum[i] = stackNum[i-1] - stackNum[i]
		}
	}

	//fmt.Println(stackNum)
	return stackNum[len(stackNum)-1]
}

func isOperator(n byte) bool {
	return n == '+' || n == '-' || n == '*' || n == '/'
}
