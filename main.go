package main

import count_complete_tree_nodes "github.com/xingxingso/leetcode/count-complete-tree-nodes"

func main() {
	//count_complete_tree_nodes.CountNodesS1(CreateTreeByArray([]int{1, 2, 3, 4, 5, 6}))
}

func CreateTreeByArray(nums []int) *count_complete_tree_nodes.TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	queue := make([]*count_complete_tree_nodes.TreeNode, 0)
	head := &count_complete_tree_nodes.TreeNode{Val: nums[0]}
	queue = append(queue, head)
	for i := 1; i < len(nums); {
		if nums[i] != 0 {
			queue[0].Left = &count_complete_tree_nodes.TreeNode{Val: nums[i]}
			queue = append(queue, queue[0].Left)
		}
		if i+1 < len(nums) && nums[i+1] != 0 {
			queue[0].Right = &count_complete_tree_nodes.TreeNode{Val: nums[i+1]}
			queue = append(queue, queue[0].Right)
		}

		queue = queue[1:]
		i += 2
	}
	return head
}
