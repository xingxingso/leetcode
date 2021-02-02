/*
https://leetcode-cn.com/problems/boundary-of-binary-tree/

545. 二叉树的边界

二叉树的边界是按顺序包括 左边界 、叶节点 和 右边界 而 不包括重复的节点 。
左边界 的定义是从根到 最左侧 节点的路径。右边界 的定义是从根到 最右侧 节点的路径。
最左侧 节点的定义是：在左子树存在时总是优先访问，如果不存在左子树则访问右子树。重复以上操作，首先抵达的节点就是 最左侧 节点。
最右侧 节点的定义方式相同，只是将左替换成右。
注意，根节点可能是 最左侧 和/或 最右侧 的节点。
给定一棵二叉树的根节点 root ，以 逆时针 顺序从根开始返回其 边界 。

提示：
	树中节点的数目在范围 [0, 104] 内
	-1000 <= Node.val <= 1000

*/
package boundary_of_binary_tree

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 官方

/*
方法一: 简单解法

时间复杂度：O(N)，一次完整的叶子节点遍历，和两次深度的遍历。
空间复杂度：O(N)，res 和 stack 的空间开销。
*/
func boundaryOfBinaryTreeO1(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	if !isLeaf(root) {
		res = append(res, root.Val)
	}
	t := root.Left
	for t != nil {
		if !isLeaf(t) {
			res = append(res, t.Val)
		}
		if t.Left != nil {
			t = t.Left
		} else {
			t = t.Right
		}
	}
	addLeaves(&res, root)
	stack := make([]int, 0)
	t = root.Right
	for t != nil {
		if !isLeaf(t) {
			stack = append(stack, t.Val)
		}
		if t.Right != nil {
			t = t.Right
		} else {
			t = t.Left
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}
	return res
}

func isLeaf(t *TreeNode) bool {
	return t.Left == nil && t.Right == nil
}

func addLeaves(res *[]int, root *TreeNode) {
	if isLeaf(root) {
		*res = append(*res, root.Val)
	}

	if root.Left != nil {
		addLeaves(res, root.Left)
	}
	if root.Right != nil {
		addLeaves(res, root.Right)
	}
}

/*
方法二: 先序遍历

时间复杂度：O(N)，只需要一次树的遍历。
空间复杂度：O(N)，递归栈的深度，此外还有 left_boundary，right_boundary 和 leaves 数组的大小。
*/
func boundaryOfBinaryTreeO2(root *TreeNode) []int {
	leftBoundary := make([]int, 0)
	rightBoundary := make([]int, 0)
	leaves := make([]int, 0)
	preOrder(root, &leftBoundary, &rightBoundary, &leaves, 0)
	leftBoundary = append(leftBoundary, leaves...)
	leftBoundary = append(leftBoundary, rightBoundary...)
	return leftBoundary
}

func preOrder(cur *TreeNode, leftBoundary, rightBoundary, leaves *[]int, flag int) {
	if cur == nil {
		return
	}
	if isRightBoundary(flag) {
		// 先这么处理
		*rightBoundary = append([]int{cur.Val}, *rightBoundary...)
	} else if isLeftBoundary(flag) || isRoot(flag) {
		*leftBoundary = append(*leftBoundary, cur.Val)
	} else if isLeaf(cur) {
		*leaves = append(*leaves, cur.Val)
	}
	preOrder(cur.Left, leftBoundary, rightBoundary, leaves, leftChildFlag(cur, flag))
	preOrder(cur.Right, leftBoundary, rightBoundary, leaves, rightChildFlag(cur, flag))
}
func isRoot(flag int) bool {
	return flag == 0
}

func isLeftBoundary(flag int) bool {
	return flag == 1
}

func isRightBoundary(flag int) bool {
	return flag == 2
}

func leftChildFlag(cur *TreeNode, flag int) int {
	if isLeftBoundary(flag) || isRoot(flag) {
		return 1
	}
	if isRightBoundary(flag) && cur.Right == nil {
		return 2
	}
	return 3
}

func rightChildFlag(cur *TreeNode, flag int) int {
	if isRightBoundary(flag) || isRoot(flag) {
		return 2
	}
	if isLeftBoundary(flag) && cur.Left == nil {
		return 1
	}
	return 3
}
