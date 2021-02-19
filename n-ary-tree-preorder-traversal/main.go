/*
https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

589. N叉树的前序遍历

给定一个 N 叉树，返回其节点值的前序遍历。

说明:
	递归法很简单，你可以使用迭代法完成此题吗?

*/
package n_ary_tree_preorder_traversal

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
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	ans := make([]int, 0)
	ans = append(ans, root.Val)
	for _, child := range root.Children {
		ans = append(ans, preorder(child)...)
	}

	return ans
}

/*
方法二: 迭代

时间复杂度：
空间复杂度：
*/
func preorderS2(root *Node) []int {
	if root == nil {
		return nil
	}

	stack := []*Node{root}
	ans := make([]int, 0)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return ans
}
