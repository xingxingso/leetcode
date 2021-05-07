/*
https://leetcode-cn.com/problems/count-complete-tree-nodes/

222. 完全二叉树的节点个数

给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。
若最底层为第 h 层，则该层包含 1~ 2^h 个节点。

提示：
	树中节点的数目范围是[0, 5 * 10^4]
	0 <= Node.val <= 5 * 10^4
	题目数据保证输入的树是 完全二叉树

进阶：
	遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？

*/
package count_complete_tree_nodes

import (
	"sort"
)

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485057&idx=1&sn=45a3b89a4efef236cb662d5505d7ce36&chksm=9bd7f889aca0719f4915de681f983355e187151030991ab1944494ffe4b73e484068b85eb01e&scene=21#wechat_redirect

/*
方法一:

时间复杂度：O(logN*logN)
空间复杂度：
*/
func countNodesP1(root *TreeNode) int {
	hl, hr := 0, 0
	l, r := root, root
	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil {
		r = r.Right
		hr++
	}
	if hl == hr {
		return 1<<hl - 1
	}
	return 1 + countNodesP1(root.Left) + countNodesP1(root.Right)
}

// --- 官方

/*
方法一: 二分查找 + 位运算

时间复杂度：O(log_2(n))，其中 n 是完全二叉树的节点数。
	首先需要 O(h) 的时间得到完全二叉树的最大层数，其中 h 是完全二叉树的最大层数。
	使用二分查找确定节点个数时，需要查找的次数为 O(log2^h)=O(h)，每次查找需要遍历从根节点开始的一条长度为 h 的路径，需要 O(h) 的时间，
	因此二分查找的总时间复杂度是 O(h^2)。因此总时间复杂度是 O(h^2)。由于完全二叉树满足 2^h <= n < 2^{h+1}，
	因此有 O(h)=O(logn)，O(h^2)=O(log_2(n))。
空间复杂度：O(1)。只需要维护有限的额外空间。
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0

	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		// fmt.Println(level, k, 1<<level)
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		// fmt.Printf("%.4b, %.4b\n", k, bits)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}
