/*
Package average_of_levels_in_binary_tree
https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/

637. 二叉树的层平均值

给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10^-5 以内的答案可以被接受。

提示：
	树中节点数量在 [1, 10^4] 范围内
	-2^31 <= Node.val <= 2^31 - 1

*/
package average_of_levels_in_binary_tree

/*
TreeNode
Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// --- 自己

/*
方法一: 层序遍历

时间复杂度：
空间复杂度：
*/
func averageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{root}
	ans := make([]float64, 0)
	for len(queue) > 0 {
		sum := 0
		l := len(queue)
		for i := 0; i < l; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
		ans = append(ans, float64(sum)/float64(l))
	}

	return ans
}
