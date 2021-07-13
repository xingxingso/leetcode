/*
Package single_number
https://leetcode-cn.com/problems/single-number/

136. 只出现一次的数字

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：
	你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

*/
package single_number

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func singleNumber(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
