/*
Package majority_element
https://leetcode-cn.com/problems/majority-element/

169. 多数元素

给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

进阶：
	尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

*/
package majority_element

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
func majorityElement(nums []int) int {
	var p, count int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			p = nums[i]
			count++
			continue
		}

		if nums[i] != p {
			count--
		} else {
			count++
		}
	}

	//if count == 0 { // 输入有问题
	//	return -1
	//}

	return p
}
