/*
Package rotate_array
https://leetcode-cn.com/problems/rotate-array/

189. 旋转数组

给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

进阶：
	尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
	你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

*/
package rotate_array

// --- 他人
// 参考 https://www.xuetangx.com/learn/THU08091000384/THU08091000384/7755489/video/12847446

/*
方法一: 环状替换

时间复杂度：O(n)
空间复杂度：O(1)
*/
func rotateP1(nums []int, k int) {
	k = k % len(nums)
	for s, mov := 0, 0; mov < len(nums); s++ { // O(g) = O(GCD(n, k))
		mov += shift(nums, s, k)
	}
}

// s->s+k, s+k->s+2k
func shift(nums []int, s, k int) int { //从A[s]出发，以k为间隔循环右移，O(n/GCD(n, k))
	mov := 0                   //移动次数
	i, j := s, (s+k)%len(nums) //从该元素出发
	pre := nums[i]
	for s != j { //以k为间隔
		next := nums[j]
		nums[j] = pre //依次右移k位
		i, j = j, (j+k)%len(nums)
		pre = next
		mov++
	}
	nums[s] = pre
	return mov + 1
}

/*
方法二: 数组倒置
	虽然此方法的复杂度略高于方法一，O(3n), 但是这种方式更充分的使用到了 Cache,
	所以当数组比较大的时候，这个执行速度更快于方法一

时间复杂度：O(n)
空间复杂度：O(1)
*/
func rotateP2(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)     // O(3n/2)
	reverse(nums[:k]) // O(3k/2)
	reverse(nums[k:]) // O(3(n-k)/2)
} // O(3n)

// 数组翻转
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
