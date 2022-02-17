/*
Package construct_binary_tree_from_preorder_and_postorder_traversal
https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

889. 根据前序和后序遍历构造二叉树

给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，
postorder 是同一棵树的后序遍历，重构并返回二叉树。

如果存在多个答案，您可以返回其中 任何 一个。

提示：
	1 <= preorder.length <= 30
	1 <= preorder[i] <= preorder.length
	preorder 中所有值都 不同
	postorder.length == preorder.length
	1 <= postorder[i] <= postorder.length
	postorder 中所有值都 不同
	保证 preorder 和 postorder 是同一棵二叉树的前序遍历和后序遍历

*/
package construct_binary_tree_from_preorder_and_postorder_traversal

import "fmt"

// TreeNode Definition for a binary tree node.
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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	preorder = preorder[1:]
	postorder = postorder[:len(postorder)-1]
	if len(preorder) == 0 {
		return root
	}

	leftIndex := 0
	for i, v := range postorder {
		if v == preorder[0] {
			leftIndex = i
			break
		}
	}
	root.Left = constructFromPrePost(preorder[:leftIndex+1], postorder[:leftIndex+1])
	root.Right = constructFromPrePost(preorder[leftIndex+1:], postorder[leftIndex+1:])
	//printTree(root)
	return root
}

func printTree(root *TreeNode) {
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		fmt.Printf("l:%d", level)
		size := len(queue)
		noLeft := true
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				fmt.Print(" nil ")
				queue = append(queue, nil)
				queue = append(queue, nil)
				continue
			}
			fmt.Printf(" %d ", node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
			if node.Left != nil || node.Right != nil {
				noLeft = false
			}
		}
		queue = queue[size:]
		level++
		fmt.Println()
		if noLeft {
			break
		}
	}
}
