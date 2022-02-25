/*
Package reverse_string
https://leetcode-cn.com/problems/reverse-string/

344. 反转字符串

编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

提示：
	1 <= s.length <= 10^5
	s[i] 都是 ASCII 码表中的可打印字符

*/
package reverse_string

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return
}
