/*
https://leetcode-cn.com/problems/permutations/

46. 全排列

给定一个 没有重复 数字的序列，返回其所有可能的全排列。

*/
package permutations

import "fmt"

// --- 官方

/*
方法一: 搜索回溯

时间复杂度：O(n * n!)，其中 n 为序列的长度。
空间复杂度：空间复杂度：O(n)，其中 n 为序列的长度。除答案数组以外，递归函数在递归过程中需要为每一层递归函数分配栈空间，
	所以这里需要额外的空间且该空间取决于递归的深度，这里可知递归调用深度为 O(n)。
*/
func permuteO1(nums []int) [][]int {
	var result [][]int
	var currentArray []int
	length := len(nums)

	var backtrack func(position int, currentArray []int)
	backtrack = func(position int, currentArray []int) {
		fmt.Println(currentArray)
		if position == length {
			result = append(result, append([]int(nil), currentArray...))
			return
		}

		for i := position; i < length; i++ {
			currentArray = append(currentArray, nums[i])
			nums[position], nums[i] = nums[i], nums[position]
			backtrack(position+1, currentArray)
			nums[position], nums[i] = nums[i], nums[position]
			currentArray = currentArray[:len(currentArray)-1]
		}
	}
	backtrack(0, currentArray)

	return result
}

// --- 他人

/*
https://leetcode-cn.com/problems/permutations/solution/di-gui-shuang-bai-dao-shou-by-xiaoma_nmg/
方法一: 递归
	// 我也是这么想的，但是没有想到

时间复杂度：
空间复杂度：
*/
func permuteP1(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := [][]int{}

	for i, num := range nums {
		// 把num从 nums 拿出去 得到tmp
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])

		// sub 是把num 拿出去后，数组中剩余数据的全排列
		sub := permuteP1(tmp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}
	}
	return res
}

/*func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
*/
