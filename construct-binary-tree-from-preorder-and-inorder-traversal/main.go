/*
https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

105. 从前序与中序遍历序列构造二叉树

根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
	你可以假设树中没有重复的元素。
*/
package construct_binary_tree_from_preorder_and_inorder_traversal

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
	利用前序遍历找到根节点，找到在根节点在中序遍历中的位置，左边为左树，右边为右树

时间复杂度：
空间复杂度：
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}

	if len(preorder) == 1 {
		return root
	}
	// 可以通过维护一个map 优化
	index := searchIndex(inorder, root.Val)

	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])

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
