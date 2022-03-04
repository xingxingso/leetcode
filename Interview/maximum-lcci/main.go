/*
Package maximum_lcci
https://leetcode-cn.com/problems/maximum-lcci/

面试题 16.07. 最大数值

编写一个方法，找出两个数字a和b中最大的那一个。不得使用if-else或其他比较运算符。

*/
package maximum_lcci

// --- 他人

/*
方法一:

时间复杂度：
空间复杂度：
*/
func maximum(a int, b int) int {
	//非负数右移高位补0,负数右移高位补1, 即 arithmetic right shift
	//a>b,ret>0,temp&1=0
	//a<b,ret<0,temp&1=1
	ret := int64(a - b)
	ret = int64(a) - ret&(ret>>63)
	return int(ret)
}
