/*
https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree/

255. 验证前序遍历序列二叉搜索树

给定一个整数数组，你需要验证它是否是一个二叉搜索树正确的先序遍历序列。
你可以假定该序列中的数都是不相同的。

进阶挑战：
	您能否使用恒定的空间复杂度来完成此题？

*/
package verify_preorder_sequence_in_binary_search_tree

import "math"

// --- 他人

/*
https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree/solution/python3-tu-jie-by-ml-zimingmeng/
方法一: 栈

时间复杂度：O(N)。遍历了一遍数组
空间复杂度：O(N)。最坏条件下为O(N)
*/
func verifyPreorder(preorder []int) bool {
	var stack []int
	minValue := math.MinInt32
	for i := 0; i < len(preorder); i++ {
		if preorder[i] < minValue {
			return false
		}
		for len(stack) > 0 && preorder[i] > stack[len(stack)-1] {
			minValue = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, preorder[i])
	}

	return true
}
