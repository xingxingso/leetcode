/*
Package single_number_iii
https://leetcode-cn.com/problems/single-number-iii/

260. 只出现一次的数字 III

给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。

进阶：
	你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

提示：
	2 <= nums.length <= 3 * 10^4
	-2^31 <= nums[i] <= 2^31 - 1
	除两个只出现一次的整数外，nums 中的其他数字都出现两次

*/
package single_number_iii

// --- 官方

/*
方法一:

时间复杂度：O(n)，其中 n 是数组 nums 的长度。
空间复杂度：O(1)。
*/
func singleNumber(nums []int) []int {
	x := 0
	for _, n := range nums {
		x ^= n
	}

	lsb := x ^ (-x) // 取最低位的1
	type1, type2 := 0, 0
	for _, n := range nums {
		if n&lsb > 0 {
			type1 ^= n
		} else {
			type2 ^= n
		}
	}
	ans := []int{type1, type2}

	return ans
}
