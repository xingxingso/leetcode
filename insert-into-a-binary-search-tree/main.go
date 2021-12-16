/*
Package insert_into_a_binary_search_tree
https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/

701. 二叉搜索树中的插入操作

给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。
注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。

提示：
	给定的树上的节点数介于 0 和 10^4 之间
	每个节点都有一个唯一整数值，取值范围从 0 到 10^8
	-10^8 <= val <= 10^8
	新值和原始二叉搜索树中的任意节点值都不同

*/
package insert_into_a_binary_search_tree

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 深度优先搜索

时间复杂度：
空间复杂度：
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var dfs func(root *TreeNode, val int) *TreeNode
	dfs = func(root *TreeNode, val int) *TreeNode {
		if root == nil {
			return &TreeNode{Val: val}
		}
		if root.Val > val {
			root.Left = dfs(root.Left, val)
		} else { // 不存在相同
			root.Right = dfs(root.Right, val)
		}
		return root
	}
	return dfs(root, val)
}
