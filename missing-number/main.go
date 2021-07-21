/*
Package missing_number
https://leetcode-cn.com/problems/missing-number/

268. 丢失的数字

给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

进阶：
	你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?

提示：
	n == nums.length
	1 <= n <= 10^4
	0 <= nums[i] <= n
	nums 中的所有数字都 独一无二

*/
package missing_number

// --- 自己

/*
方法一: 异或运算
时间复杂度：
空间复杂度：
*/
func missingNumber(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i] ^ i
	}
	return ans ^ len(nums)
}

// --- 他人
//https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484477&idx=1&sn=13834cfd618377385226c3dc598b2c28&scene=21#wechat_redirect

/*
方法二: 等差数列
时间复杂度：
空间复杂度：
*/
func missingNumberP1(nums []int) int {
	n := len(nums)
	// 公式：(首项 + 末项) * 项数 / 2
	expect := (0 + n) * (n + 1) / 2

	sum := 0
	for _, x := range nums {
		sum += x
	}
	return expect - sum
}

/*
方法二: 等差数列 优化
	原来求和可能溢出
时间复杂度：
空间复杂度：
*/
func missingNumberP2(nums []int) int {
	n := len(nums)
	res := 0
	// 新补的索引
	res += n - 0
	// 剩下索引和元素的差加起来
	for i := 0; i < n; i++ {
		res += i - nums[i]
	}
	return res
}
