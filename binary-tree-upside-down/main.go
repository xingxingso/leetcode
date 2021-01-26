/*
https://leetcode-cn.com/problems/binary-tree-upside-down/

156. 上下翻转二叉树
给定一个二叉树，其中所有的右节点要么是具有兄弟节点（拥有相同父节点的左节点）的叶节点，要么为空，将此二叉树上下翻转并将它变成一棵树，
原来的右节点将转换成左叶节点。返回新的根。
*/
package binary_tree_upside_down

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 他人

/*
https://leetcode-cn.com/problems/binary-tree-upside-down/solution/shang-xia-fan-zhuan-er-cha-shu-by-nortondark/
方法一:
*/
func upsideDownBinaryTreeP1(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) { // root为空或是叶子节点返回
		return root
	}

	newRoot := upsideDownBinaryTreeP1(root.Left) // 只需递归处理左子树，不需要递归右子树，右子树都是叶子节点

	root.Left.Left = root.Right // 三角关系翻转
	root.Left.Right = root

	// 根或子树的root变为右叶子节点

	// 注意，题解说：所有右节点，都是叶子节点，且有兄弟节点。所以，root旋转后，都会变为右叶子节点，所以left和right设为null。
	// 估计很多人看漏了这个条件，感觉莫名其妙。我自己还模拟右子节点还有两个子节点的情况，想了半天。。
	root.Left = nil
	root.Right = nil
	return newRoot // 同链表翻转，返回整颗树最左的叶子节点
}

/*
https://leetcode-cn.com/problems/binary-tree-upside-down/solution/binary-tree-upside-down-top-downdie-dai-fa-by-jin4/
方法一:
*/
func upsideDownBinaryTreeP2(root *TreeNode) *TreeNode {
	var parent, parentRight *TreeNode = nil, nil

	for root != nil {
		rootLeft := root.Left
		root.Left = parentRight
		parentRight = root.Right
		root.Right = parent
		parent = root
		root = rootLeft
	}
	return parent
}
