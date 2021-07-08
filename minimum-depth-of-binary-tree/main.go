/*
Package minimum_depth_of_binary_tree
https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

111. 二叉树的最小深度

给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：
	叶子节点是指没有子节点的节点。

提示：
	树中节点数的范围在 [0, 10^5] 内
	-1000 <= Node.val <= 1000

*/
package minimum_depth_of_binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 深度优先遍历

时间复杂度：O(n)
空间复杂度：O(H)
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 {
		return right + 1
	}
	if right == 0 {
		return left + 1
	}
	/*left, right := math.MaxInt32, math.MaxInt32
	if root.Left != nil {
		left = minDepth(root.Left)
	}
	if root.Right != nil {
		right = minDepth(root.Right)
	}*/

	return min(left, right) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// --- 官方

/*
方法一: 深度优先搜索

时间复杂度：O(N)，其中 N 是树的节点数。对每个节点访问一次。
空间复杂度：O(H)，其中 H 是树的高度。空间复杂度主要取决于递归时栈空间的开销，最坏情况下，树呈现链状，空间复杂度为 O(N)。
	平均情况下树的高度与节点数的对数正相关，空间复杂度为 O(logN)。
*/
func minDepthO1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepthO1(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepthO1(root.Right), minD)
	}
	return minD + 1
}

/*
方法二: 广度优先搜索

时间复杂度：O(N)，其中 N 是树的节点数。对每个节点访问一次。
空间复杂度：O(N)，其中 N 是树的节点数。空间复杂度主要取决于队列的开销，队列中的元素个数不会超过树的节点数。
*/
func minDepthO2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	count := []int{}
	queue = append(queue, root)
	count = append(count, 1)
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			count = append(count, depth+1)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}
