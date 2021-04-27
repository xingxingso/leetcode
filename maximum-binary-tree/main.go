/*
https://leetcode-cn.com/problems/maximum-binary-tree/

654. 最大二叉树

给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：
	1.二叉树的根是数组 nums 中的最大元素。
	2.左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
	3.右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
返回有给定数组 nums 构建的 最大二叉树 。

提示：
	1 <= nums.length <= 1000
	0 <= nums[i] <= 1000
	nums 中的所有整数 互不相同

*/
package maximum_binary_tree

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
方法一: 递归

时间复杂度：
空间复杂度：
*/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	pos := 0
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			pos = i
		}
	}
	root := &TreeNode{Val: max}
	root.Left = constructMaximumBinaryTree(nums[:pos])
	root.Right = constructMaximumBinaryTree(nums[pos+1:])
	return root
}
