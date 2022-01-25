/*
Package degree_of_an_array
https://leetcode-cn.com/problems/degree-of-an-array/

697. 数组的度

给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

提示：
	nums.length 在 1 到 50,000 范围内。
	nums[i] 是一个在 0 到 49,999 范围内的整数。

*/
package degree_of_an_array

// --- 自己

/*
方法一: 哈希表

时间复杂度：
空间复杂度：
*/
func findShortestSubArray(nums []int) int {
	hash := make(map[int][3]int, 0) // count, left index, right index
	max := 0
	length := 0
	for k, v := range nums {
		if m, ok := hash[v]; !ok {
			hash[v] = [3]int{1, k, k}
		} else {
			m[0]++
			hash[v] = [3]int{m[0], m[1], k}
		}
		if hash[v][0] > max || hash[v][0] == max && (length > hash[v][2]-hash[v][1]+1) { // 度相同，但是连续最短
			max = hash[v][0]
			length = hash[v][2] - hash[v][1] + 1
		}
	}

	return length
}
