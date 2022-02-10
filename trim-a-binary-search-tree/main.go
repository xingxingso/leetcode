/*
Package trim_a_binary_search_tree
https://leetcode-cn.com/problems/trim-a-binary-search-tree/

669. 修剪二叉搜索树

给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。
修剪树 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。

所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。

提示：
	树中节点数在范围 [1, 10^4] 内
	0 <= Node.val <= 10^4
	树中每个节点的值都是 唯一 的
	题目数据保证输入是一棵有效的二叉搜索树
	0 <= low <= high <= 10^4

*/
package trim_a_binary_search_tree

/*
TreeNode
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 深度优先遍历
	小于 low 的节点及其左子节点删除，其右节点中接续执行

时间复杂度：
空间复杂度：
*/
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
