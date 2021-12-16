/*
Package maximum_depth_of_n_ary_tree
https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/

559. N 叉树的最大深度

给定一个 N 叉树，找到其最大深度。
最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。

提示：
	树的深度不会超过 1000 。
	树的节点数目位于 [0, 104] 之间。

*/
package maximum_depth_of_n_ary_tree

/**
 * Definition for a Node.
 */
type Node struct {
	Val      int
	Children []*Node
}

// --- 自己

/*
方法一: 深度优先搜索
	复杂度分析来自官方

时间复杂度：每个节点遍历一次，所以时间复杂度是 O(N)，其中 N 为节点数。
空间复杂度：最坏情况下, 树完全非平衡，
	例如 每个节点有且仅有一个孩子节点，递归调用会发生 N 次（等于树的深度），所以存储调用栈需要 O(N)。
	但是在最好情况下（树完全平衡），树的高度为 log(N)。
	所以在此情况下空间复杂度为 O(log(N))。
*/
func maxDepth(root *Node) int {
	ans := 0
	var dfs func(node *Node, dep int)
	dfs = func(node *Node, dep int) {
		if node == nil {
			return
		}
		dep++
		for _, child := range node.Children {
			dfs(child, dep)
		}
		ans = max(ans, dep)
	}
	dfs(root, 0)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
