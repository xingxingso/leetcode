/*
https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

116. 填充每个节点的下一个右侧节点指针

给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
	struct Node {
	  int val;
	  Node *left;
	  Node *right;
	  Node *next;
	}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有 next 指针都被设置为 NULL。

进阶：
	你只能使用常量级额外空间。
	使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

提示：
	树中节点的数量少于 4096
	-1000 <= node.val <= 1000

*/
package populating_next_right_pointers_in_each_node

/**
Definition for a binary tree node.
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// --- 自己

/*
方法一: 层次遍历

时间复杂度：O(N)。每个节点会被访问一次且只会被访问一次，即从队列中弹出，并建立 next 指针。
空间复杂度：O(N)。这是一棵完美二叉树，它的最后一个层级包含 N/2 个节点。
	广度优先遍历的复杂度取决于一个层级上的最大元素数量。这种情况下空间复杂度为 O(N)。
*/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if i < l-1 {
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}

	return root
}

/*
方法二: 使用已建立的 next 指针
	官方提供的思路

时间复杂度：
空间复杂度：
*/
func connectS2(root *Node) *Node {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil || node.Left == nil {
			return
		}
		node.Left.Next = node.Right
		if node.Next != nil {
			node.Right.Next = node.Next.Left
		}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return root
}

// --- 官方

/*
方法二: 使用已建立的 next 指针

时间复杂度：O(N)，每个节点只访问一次。
空间复杂度：O(1)，不需要存储额外的节点。
*/
func connectO1(root *Node) *Node {
	if root == nil {
		return root
	}

	// 每次循环从该层的最左侧节点开始
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		// 通过 Next 遍历这一层节点，为下一层的节点更新 Next 指针
		for node := leftmost; node != nil; node = node.Next {
			// 左节点指向右节点
			node.Left.Next = node.Right

			// 右节点指向下一个左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	// 返回根节点
	return root
}

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247487126&idx=1&sn=4de13e66397bc35970963c5a1330ce18&chksm=9bd7f09eaca0798853c41fba05ad5fa958b31054eba18b69c785ae92f4bd8e4cc7a2179d7838&scene=21#wechat_redirect

/*
方法一:

时间复杂度：
空间复杂度：
*/
func connectP1(root *Node) *Node {
	var connectNode func(node1, node2 *Node)
	connectNode = func(node1, node2 *Node) {
		if node1 == nil || node2 == nil {
			return
		}

		node1.Next = node2
		connectNode(node1.Left, node1.Right)
		connectNode(node2.Left, node2.Right)
		// 连接跨越父节点的两个子节点
		connectNode(node1.Right, node2.Left)
	}
	connectNode(root.Left, root.Right)
	return root
}
