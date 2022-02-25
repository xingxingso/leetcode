/*
Package permutations
https://leetcode-cn.com/problems/permutations/

46. 全排列

给定一个 没有重复 数字的序列，返回其所有可能的全排列。

*/
package permutations

// --- 自己

/*
方法一: 回溯

时间复杂度：
空间复杂度：
*/
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	item := make([]int, 0)
	var backTrack func()
	backTrack = func() {
		if len(item) == len(nums) {
			temp := make([]int, len(item))
			copy(temp, item)
			ans = append(ans, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if contains(item, nums[i]) {
				continue
			}
			item = append(item, nums[i])
			backTrack()
			item = item[:len(item)-1]
		}
		return
	}
	backTrack()
	return ans
}

func contains(s []int, target int) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == target {
			return true
		}
	}
	return false
}

// --- 官方

/*
方法一: 搜索回溯

时间复杂度：O(n * n!)，其中 n 为序列的长度。
空间复杂度：空间复杂度：O(n)，其中 n 为序列的长度。除答案数组以外，递归函数在递归过程中需要为每一层递归函数分配栈空间，
	所以这里需要额外的空间且该空间取决于递归的深度，这里可知递归调用深度为 O(n)。
*/
func permuteO1(nums []int) [][]int {
	var ret [][]int
	var cur []int

	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos == len(nums) {
			ret = append(ret, append([]int(nil), cur...))
			return
		}

		for i := pos; i < len(nums); i++ {
			cur = append(cur, nums[i])
			nums[pos], nums[i] = nums[i], nums[pos]
			backtrack(pos + 1)
			nums[pos], nums[i] = nums[i], nums[pos]
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0)

	return ret
}

// --- 他人

/*
https://leetcode-cn.com/problems/permutations/solution/di-gui-shuang-bai-dao-shou-by-xiaoma_nmg/
方法一: 递归

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
