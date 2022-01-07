/*
Package product_of_array_except_self
https://leetcode-cn.com/problems/product-of-array-except-self/

238. 除自身以外数组的乘积

给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

提示：
	题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。

说明:
	请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
	你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

*/
package product_of_array_except_self

// --- 自己

/*
方法一:
	// 前缀积 + 后缀积

时间复杂度：O(n)
空间复杂度：O(1)
*/
func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))

	prefix[0] = 1
	for i := 0; i < len(nums)-1; i++ {
		prefix[i+1] = prefix[i] * nums[i]
	}
	suffix[len(nums)-1] = 1
	for i := len(nums) - 1; i > 0; i-- {
		suffix[i-1] = suffix[i] * nums[i]
	}

	//fmt.Println(prefix, suffix)

	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = prefix[i] * suffix[i]
	}

	return ans
}

/*
方法一: 优化
	// 前缀积 + 后缀积

时间复杂度：O(n)
空间复杂度：O(1)
*/
func productExceptSelfS1(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 0; i < len(nums)-1; i++ {
		ans[i+1] = ans[i] * nums[i]
	}
	//fmt.Println(ans)

	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * r
		r *= nums[i]
	}

	return ans
}
