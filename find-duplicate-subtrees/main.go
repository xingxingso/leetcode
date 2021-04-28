/*
https://leetcode-cn.com/problems/find-duplicate-subtrees/

652. 寻找重复的子树

给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
两棵树重复是指它们具有相同的结构以及相同的结点值。

注意:
	合并必须从两个树的根节点开始。

*/
package find_duplicate_subtrees

import (
	"strconv"
)

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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	memo := make(map[string]int)
	var traverse func(root *TreeNode) string
	traverse = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		left := traverse(root.Left)
		right := traverse(root.Right)
		treeString := strconv.Itoa(root.Val) + "," + left + "," + right
		if n, ok := memo[treeString]; n == 1 && ok {
			res = append(res, root)
		}
		memo[treeString]++
		// fmt.Println(memo, res)
		return treeString
	}

	_ = traverse(root)
	return res
}
