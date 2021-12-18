/*
Package balanced_binary_tree
https://leetcode-cn.com/problems/balanced-binary-tree/

110. 平衡二叉树

给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
	一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。

提示：
	树中的节点数在范围 [0, 5000] 内
	-10^4 <= Node.val <= 10^4

*/
package balanced_binary_tree

/*
TreeNode
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 官方

/*
方法一: 自顶向下的递归

时间复杂度：O(n^2)，其中 n 是二叉树中的节点个数。
	最坏情况下，二叉树是满二叉树，需要遍历二叉树中的所有节点，时间复杂度是 O(n)。
	对于节点 p，如果它的高度是 d，则 height(p) 最多会被调用 d 次（即遍历到它的每一个祖先节点时）。
	对于平均的情况，一棵树的高度 h 满足 O(h)=O(log n)，因为 d<=h，所以总时间复杂度为 O(nlogn)。
	对于最坏的情况，二叉树形成链式结构，高度为 O(n)，此时总时间复杂度为 O(n^2)。
空间复杂度：O(n)，其中 n 是二叉树中的节点个数。空间复杂度主要取决于递归调用的层数，递归调用的层数不会超过 n。
*/
func isBalancedO1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(heightO1(root.Left)-heightO1(root.Right)) <= 1 && isBalancedO1(root.Left) && isBalancedO1(root.Right)
}

func heightO1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(heightO1(root.Left), heightO1(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

/*
方法二: 自底向上的递归

时间复杂度：O(n)，其中 n 是二叉树中的节点个数。使用自底向上的递归，每个节点的计算高度和判断是否平衡都只需要处理一次，
	最坏情况下需要遍历二叉树中的所有节点，因此时间复杂度是 O(n)。
空间复杂度：O(n)，其中 n 是二叉树中的节点个数。空间复杂度主要取决于递归调用的层数，递归调用的层数不会超过 n。
*/
func isBalancedO2(root *TreeNode) bool {
	return heightO2(root) >= 0
}

func heightO2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := heightO2(root.Left)
	rightHeight := heightO2(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return max(leftHeight, rightHeight) + 1
}
