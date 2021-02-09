package public

/**
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func CreateTreeByArray(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	queue := make([]*TreeNode, 0)
	head := &TreeNode{Val: nums[0]}
	queue = append(queue, head)
	for i := 1; i < len(nums); {
		if nums[i] != 0 {
			queue[0].Left = &TreeNode{Val: nums[i]}
			queue = append(queue, queue[0].Left)
		}
		if i+1 < len(nums) && nums[i+1] != 0 {
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

func root1() *TreeNode {
	node5 := &TreeNode{}
	node5.Val = 5

	node4 := &TreeNode{}
	node4.Val = 4

	node3 := &TreeNode{}
	node3.Val = 3

	node2 := &TreeNode{}
	node2.Val = 2
	node2.Left = node4
	node2.Right = node5

	node1 := &TreeNode{}
	node1.Val = 1
	node1.Left = node2
	node1.Right = node3

	return node1
}

func want1() *TreeNode {
	node3 := &TreeNode{}
	node3.Val = 3

	node1 := &TreeNode{}
	node1.Val = 1

	node5 := &TreeNode{}
	node5.Val = 5

	node2 := &TreeNode{}
	node2.Val = 2
	node2.Left = node3
	node2.Right = node1

	node4 := &TreeNode{}
	node4.Val = 4
	node4.Left = node5
	node4.Right = node2

	return node4
}
