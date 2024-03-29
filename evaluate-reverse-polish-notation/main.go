/*
Package evaluate_reverse_polish_notation
https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/

150. 逆波兰表达式求值

根据 逆波兰表示法，求表达式的值。
有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：
	整数除法只保留整数部分。
	给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

提示：
	1 <= tokens.length <= 10^4
	tokens[i] 要么是一个算符（"+"、"-"、"*" 或 "/"），要么是一个在范围 [-200, 200] 内的整数

逆波兰表达式：
	逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。
		- 平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
		- 该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。

	逆波兰表达式主要有以下两个优点：
		- 去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
		- 适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。

*/
package evaluate_reverse_polish_notation

import "strconv"

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func evalRPN(tokens []string) int {
	numStack := make([]int, 0)
	for _, v := range tokens {
		num := 0
		if !isOperator(v) {
			num, _ = strconv.Atoi(v)
		} else {
			num = calculate(numStack[len(numStack)-2], numStack[len(numStack)-1], v)
			numStack = numStack[:len(numStack)-2]
		}
		numStack = append(numStack, num)
	}
	return numStack[len(numStack)-1]
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}
