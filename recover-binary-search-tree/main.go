/*
Package recover_binary_search_tree
https://leetcode-cn.com/problems/recover-binary-search-tree/

99. 恢复二叉搜索树

给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。

提示：
	树上节点的数目在范围 [2, 1000] 内
	-2^31 <= Node.val <= 2^31 - 1

进阶：
	使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？

*/
package recover_binary_search_tree

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 中序遍历

时间复杂度：
空间复杂度：
*/
func recoverTree(root *TreeNode) {
	var pre, first, second *TreeNode
	flag := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if flag == 2 {
			return
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		// 比较
		if flag == 0 && pre != nil && root.Val < pre.Val {
			//fmt.Println("pre", pre, root.Val, pre.Val)
			// get pre
			first = pre
			second = root // 如果是 1 3 2 4 这种，就要先把 2 保留到 second
			flag = 1
		} else if flag == 1 && pre != nil && root.Val < pre.Val {
			//fmt.Println("pre", pre, root.Val, pre.Val)
			// get cur
			second = root // 如果是 3 2 1 这种，后面会替换掉 second
			flag = 2
		}
		pre = root
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	//fmt.Println(first.Val, second.Val)
	first.Val, second.Val = second.Val, first.Val
	return
}
