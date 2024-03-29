package public

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func Bfs(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, 0)
			continue
		}

		res = append(res, node.Val)
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	// 移除右边的0
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == 0 {
			continue
		}

		return res[:i+1]
	}

	return res
}

// CreateTreeByArray
// 通过数组生成二叉树，nilNum 为 nil, 类似 leetcode 中数组转化为 二叉树
// [1,2,3,nilNum,4,nilNum,5,nilNum,6] 为二叉树
//                  1
//          2              3
//   nil         4      nil   5
// nil nil   nil   6
func CreateTreeByArray(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	nilNum := -9999
	queue := make([]*TreeNode, 0)
	head := &TreeNode{Val: nums[0]}
	queue = append(queue, head)
	for i := 1; i < len(nums); {
		if nums[i] != nilNum {
			queue[0].Left = &TreeNode{Val: nums[i]}
			queue = append(queue, queue[0].Left)
		}
		if i+1 < len(nums) && nums[i+1] != nilNum {
			queue[0].Right = &TreeNode{Val: nums[i+1]}
			queue = append(queue, queue[0].Right)
		}

		queue = queue[1:]
		i += 2
	}
	return head
}

func createTreeByArray1(nums []int) *TreeNode {
	var createTree func(i int, nums []int) *TreeNode
	createTree = func(i int, nums []int) *TreeNode {
		if nums[i] == 0 { // 特殊处理， 因为没有 null 可用， 先用 0 替代
			return nil
		}

		t := &TreeNode{Val: nums[i]}
		if i < len(nums) && 2*i+1 < len(nums) {
			t.Left = createTree(2*i+1, nums)
		}
		if i < len(nums) && 2*i+2 < len(nums) {
			t.Right = createTree(2*i+2, nums)
		}
		return t
	}

	root := createTree(0, nums)
	return root
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func inOrderPrint(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderPrint(root.Left)
	fmt.Printf("%d,", root.Val)
	inOrderPrint(root.Right)
	// fmt.Println()
}

func printTree(root *TreeNode) {
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		fmt.Printf("l:%d", level)
		size := len(queue)
		noLeft := true
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node == nil {
				fmt.Print("nil")
				queue = append(queue, nil)
				queue = append(queue, nil)
				continue
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
			if node.Left != nil || node.Right != nil {
				noLeft = false
			}
		}
		queue = queue[size:]
		level++
		fmt.Println()
		if noLeft {
			break
		}
	}
}

// 多维数组
func createNAryTree() *Node {
	return &Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{
						Val: 5,
					},
					&Node{
						Val: 6,
					},
				},
			},
			&Node{
				Val: 2,
			},
			&Node{
				Val: 4,
			},
		},
	}
}
