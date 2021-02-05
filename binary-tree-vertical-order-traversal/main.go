/*
https://leetcode-cn.com/problems/binary-tree-vertical-order-traversal/

314. 二叉树的垂直遍历

给你一个二叉树的根结点，返回其结点按 垂直方向（从上到下，逐列）遍历的结果。
如果两个结点在同一行和列，那么顺序则为 从左到右。

提示：
	树中结点的数目在范围 [0, 100] 内
	-100 <= Node.val <= 100

*/
package binary_tree_vertical_order_traversal

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
	父节点起始值为m, 则左节点为 m-1, 右节点为 m+1, 遍历顺序 m-1, m, m+1

时间复杂度：
空间复杂度：
*/
func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	hash := make(map[int][]int)
	minI, maxI := 0, 0

	var bfs func(node *TreeNode, p int)
	bfs = func(node *TreeNode, p int) {
		queue := []*treeVIndex{{
			vIndex: p,
			node:   node,
		}}

		for len(queue) > 0 {
			popNode := queue[0]
			hash[popNode.vIndex] = append(hash[popNode.vIndex], popNode.node.Val)

			if popNode.node.Left != nil {
				queue = append(queue, &treeVIndex{
					vIndex: popNode.vIndex - 1,
					node:   popNode.node.Left,
				})
				minI = min(minI, popNode.vIndex-1)
			}

			if popNode.node.Right != nil {
				queue = append(queue, &treeVIndex{
					vIndex: popNode.vIndex + 1,
					node:   popNode.node.Right,
				})
				maxI = max(maxI, popNode.vIndex+1)
			}
			queue = queue[1:]
		}
	}

	bfs(root, 0)

	for i := minI; i <= maxI; i++ {
		ans = append(ans, hash[i])
	}

	// fmt.Println(hash, minI, maxI)

	return ans
}

type treeVIndex struct {
	vIndex int
	node   *TreeNode
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
