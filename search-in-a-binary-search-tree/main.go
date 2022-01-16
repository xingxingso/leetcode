/*
Package search_in_a_binary_search_tree
https://leetcode-cn.com/problems/search-in-a-binary-search-tree/

700. 二叉搜索树中的搜索

给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。

*/
package search_in_a_binary_search_tree

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 递归
	复杂度分析来自官方

时间复杂度：O(H)，其中 H 是树高。平均时间复杂度为 O(logN)，最坏时间复杂度为 O(N)。
	使用主定理计算时间复杂度。
	T(N) = aT(N/b)+Θ(N^d)。 该公式表示花费 Θ(N^d) 时间分解一个问题，得到 a 个规模为 N/b 的子问题。
	在二叉搜索中，每次分解后只有一个子问题 a = 1，其规模为初始问题的一半 b = 2，
	每次分解花费恒定时间 d = 0，即 log_b(a)=d。 根据主定理，
	时间复杂度为 O(logN)。
空间复杂度：O(H)，递归栈的深度为 H。平均情况下深度为 O(logN)，最坏情况下深度为 O(N)。
*/
func searchBSTS1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBSTS1(root.Left, val)
	}

	if root.Val < val {
		return searchBSTS1(root.Right, val)
	}

	return nil
}

/*
方法二: 迭代
	复杂度分析来自官方

时间复杂度：O(H)，其中 H 是树高。平均时间复杂度为 O(logN)，最坏时间复杂度为 O(N)。
	使用主定理计算时间复杂度。
	T(N) = aT(N/b)+Θ(N^d)。 该公式表示花费 Θ(N^d) 时间分解一个问题，得到 a 个规模为 N/b 的子问题。
	在二叉搜索中，每次分解后只有一个子问题 a = 1，其规模为初始问题的一半 b = 2，
	每次分解花费恒定时间 d = 0，即 log_b(a)=d。 根据主定理，
	时间复杂度为 O(logN)。
空间复杂度：O(1)，恒定的额外空间。
*/
func searchBSTS2(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
	// for root != nil && root.Val != val {
	// 	if root.Val > val {
	// 		root = root.Left
	// 	} else {
	// 		root = root.Right
	// 	}
	// }
	// return nil
}
