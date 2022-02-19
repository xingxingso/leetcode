/*
Package two_sum_iv_input_is_a_bst
https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/

653. 两数之和 IV - 输入 BST

给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

提示:
	二叉树的节点个数的范围是 [1, 10^4].
	-10^4 <= Node.val <= 10^4
	root 为二叉搜索树
	-10^5 <= k <= 10^5

*/
package two_sum_iv_input_is_a_bst

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一:
	中序遍历有序， 有序数组目标和问题

时间复杂度：
空间复杂度：
*/
func findTarget(root *TreeNode, k int) bool {
	arr := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}
		arr = append(arr, root.Val)
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	for left, right := 0, len(arr)-1; left < right; {
		sum := arr[left] + arr[right]
		if sum == k {
			return true
		}
		if sum < k {
			left++
		} else {
			right--
		}
	}

	return false
}
