/*
Package delete_nodes_and_return_forest
https://leetcode-cn.com/problems/delete-nodes-and-return-forest/

1110. 删点成林

给出二叉树的根节点 root，树上每个节点都有一个不同的值。

如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。

返回森林中的每棵树。你可以按任意顺序组织答案。

提示：
	树中的节点数最大为 1000。
	每个节点都有一个介于 1 到 1000 之间的值，且各不相同。
	to_delete.length <= 1000
	to_delete 包含一些从 1 到 1000、各不相同的值。
*/
package delete_nodes_and_return_forest

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 后序遍历
	// 删除对象是叶子节点的时候：直接删除
	// 删除对象非叶子节点的时候，非空子树单独出来
	// 尽量从叶子节点开始删除，所以采用后续遍历
	// 注意头节点的特殊处理

时间复杂度：
空间复杂度：
*/
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	ans := make([]*TreeNode, 0)
	delHash := make(map[int]bool)
	for _, v := range to_delete {
		delHash[v] = true
	}

	// 删除时直接赋值为 nil 不行， 所以要引入父节点
	// 直接传递 TreeNode, 这样同样是拷贝
	var dfs func(root, parent *TreeNode, pos int)
	dfs = func(root, parent *TreeNode, pos int) {
		if root == nil {
			return
		}
		dfs(root.Left, root, 1)
		dfs(root.Right, root, 2)
		if !delHash[root.Val] {
			return
		}
		if root.Left == nil && root.Right == nil {
			if parent != nil {
				if pos == 1 {
					parent.Left = nil
				} else if pos == 2 {
					parent.Right = nil
				}
			}
			return
		}
		if root.Left != nil {
			ans = append(ans, root.Left)
		}
		if root.Right != nil {
			ans = append(ans, root.Right)
		}
		if parent != nil {
			if pos == 1 {
				parent.Left = nil
			} else if pos == 2 {
				parent.Right = nil
			}
		}
		return
	}

	dfs(root, nil, 0)
	// 根节点这里无法删除
	if root != nil && !delHash[root.Val] {
		ans = append(ans, root)
	}

	return ans
}
