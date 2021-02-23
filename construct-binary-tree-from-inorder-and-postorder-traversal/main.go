/*
https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

106. 从中序与后序遍历序列构造二叉树

根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
	你可以假设树中没有重复的元素。

*/
package construct_binary_tree_from_inorder_and_postorder_traversal

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
方法一:

时间复杂度：
空间复杂度：
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) <= 0 {
		return nil
	}

	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	if len(inorder) == 1 {
		return root
	}

	index := searchIndex(inorder, root.Val)
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

func searchIndex(inorder []int, target int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == target {
			return i
		}
	}

	return -1
}
