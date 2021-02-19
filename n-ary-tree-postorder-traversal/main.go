/*
https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/

590. N叉树的后序遍历

给定一个 N 叉树，返回其节点值的后序遍历。

说明:
	递归法很简单，你可以使用迭代法完成此题吗?

*/
package n_ary_tree_postorder_traversal

/**
 * Definition for a Node.
 */
type Node struct {
	Val      int
	Children []*Node
}

// --- 自己

/*
方法一: 递归

时间复杂度：
空间复杂度：
*/
func postorderS1(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	for _, child := range root.Children {
		ans = append(ans, postorderS1(child)...)
	}

	ans = append(ans, root.Val)
	return ans
}

/*
方法二: 迭代

时间复杂度：时间复杂度：O(M)，其中 M 是 N 叉树中的节点个数。每个节点只会入栈和出栈各一次。
空间复杂度：O(M)。在最坏的情况下，这棵 N 叉树只有 2 层，所有第 2 层的节点都是根节点的孩子。将根节点推出栈后，需要将这些节点都放入栈，
	共有 M - 1 个节点，因此栈的大小为 O(M)。
*/
func postorderS2(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	stack := []*Node{root}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append([]int{node.Val}, ans...)
		for _, child := range node.Children {
			stack = append(stack, child)
		}
	}

	return ans
}
