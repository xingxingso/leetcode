/*
Package convert_sorted_array_to_binary_search_tree
https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/

108. 将有序数组转换为二叉搜索树

将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

*/
package convert_sorted_array_to_binary_search_tree

//TreeNode Definition for a binary tree node.
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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	root := &TreeNode{}

	// mid := (0 + len(nums) - 1) / 2
	mid := (0 + len(nums)) / 2
	root.Val = nums[mid]

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
