/*
Package number_of_1_bits
https://leetcode-cn.com/problems/number-of-1-bits/

191. 位1的个数

编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。

提示：
	请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，
	因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
	在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 3 中，输入表示有符号整数 -3。

提示：
	输入必须是长度为 32 的 二进制串 。

进阶：
	如果多次调用这个函数，你将如何优化你的算法？

*/
package number_of_1_bits

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func hammingWeightS1(num uint32) int {
	res := 0
	for i := 32; i > 0; i, num = i-1, num>>1 {
		//if num&1 == 1 {
		//	res++
		//}
		res += int(num & 1)
	}

	return res
}

/*
方法二:
	n&(n-1)

时间复杂度：
空间复杂度：
*/
func hammingWeightS2(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num - 1)
		res++
	}

	return res
}
